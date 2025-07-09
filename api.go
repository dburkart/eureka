package eureka

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

type Request[T any] struct {
	request        *http.Request
	responseFormat T
}

func (r *Request[T]) Do() (*T, error) {
	var data T
	response, err := http.DefaultClient.Do(r.request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *Request[T]) DoAll(c chan *T, e chan error) {
	go func() {
		defer close(c)
		defer close(e)

		for {
			var data T
			var pagination PaginatedResponse
			response, err := http.DefaultClient.Do(r.request)
			if err != nil {
				e <- err
			}
			defer response.Body.Close()

			b, err := io.ReadAll(response.Body)
			if err != nil {
				e <- err
			}

			err = json.Unmarshal(b, &data)
			if err != nil {
				e <- err
			}

			// Send it over
			c <- &data

			err = json.Unmarshal(b, &pagination)
			if err != nil {
				e <- err
			}

			if pagination.Pagination.CurrentPage < pagination.Pagination.TotalPages {
				v := r.request.URL.Query()
				v.Set("page", strconv.Itoa(pagination.Pagination.CurrentPage+1))

				r.request.URL.RawQuery = v.Encode()
			} else {
				break
			}
		}
	}()
}

func (r *Request[T]) ForEach(callback func(*T)) error {
	var c = make(chan *T)
	var e = make(chan error)

	r.DoAll(c, e)

	for {
		select {
		case resp := <-c:
			if resp == nil {
				return nil
			}

			callback(resp)
		case err := <-e:
			if err == nil {
				return nil
			}
			return err
		}
	}
}

type AhaClientOptions struct {
	ApiKey      string          `json:"api_key"`
	CompanyName string          `json:"company_name"`
	Context     context.Context `json:"-"`
}

type AhaClient struct {
	options *AhaClientOptions
}

func NewAhaClient(options *AhaClientOptions) *AhaClient {
	if options.Context == nil {
		options.Context = context.Background()
	}

	return &AhaClient{
		options: options,
	}
}

func (c *AhaClient) must(str string, err error) string {
	if err != nil {
		panic(err)
	}
	return str
}

func (c *AhaClient) pathForEndpoint(endpoint string) (string, error) {
	return url.JoinPath(fmt.Sprintf("https://%s.aha.io/api/v1", c.options.CompanyName), endpoint)
}

func (c *AhaClient) newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(c.options.Context, method, c.must(c.pathForEndpoint(endpoint)), body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.options.ApiKey))
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	return request, nil
}

//
// Product endpoints
//

type ProductAPIRequest struct {
	client *AhaClient
}

func (c *AhaClient) Products() *ProductAPIRequest {
	return &ProductAPIRequest{c}
}

type ProductResponse struct {
	Product Product `json:"product"`
}

func (p *ProductAPIRequest) First(product *Product) error {
	var request Request[ProductResponse]

	r, err := p.client.newRequest(http.MethodGet, fmt.Sprintf("products/%s", product.ID), nil)
	if err != nil {
		panic(err)
	}

	request.request = r
	x, err := request.Do()

	if err == nil && x != nil {
		*product = x.Product
	}
	return err
}

type ProductListResponse struct {
	Products []Product `json:"products"`
}

func (p *ProductAPIRequest) List(products *[]Product) error {
	var request Request[ProductListResponse]

	r, err := p.client.newRequest(http.MethodGet, "products", nil)
	if err != nil {
		panic(err)
	}

	request.request = r
	return request.ForEach(func(resp *ProductListResponse) {
		*products = append(*products, resp.Products...)
	})
}

//
// Idea endpoints
//

type IdeaAPIRequest struct {
	client  *AhaClient
	product *Product
}

func (c *AhaClient) Ideas() *IdeaAPIRequest {
	return &IdeaAPIRequest{client: c}
}

func (i *IdeaAPIRequest) For(product *Product) *IdeaAPIRequest {
	if product != nil && product.ID != "" {
		i.product = product
	}
	return i
}

type IdeaResponse struct {
	Idea Idea `json:"idea"`
}

func (i *IdeaAPIRequest) First(idea *Idea) error {
	var request Request[IdeaResponse]

	r, err := i.client.newRequest(http.MethodGet, fmt.Sprintf("ideas/%s", idea.ID), nil)
	if err != nil {
		return err
	}

	request.request = r
	x, err := request.Do()

	if err == nil && x != nil {
		*idea = x.Idea
	}
	return err
}

type IdeaListResponse struct {
	Ideas []Idea `json:"ideas"`
}

func (i *IdeaAPIRequest) List(ideas *[]Idea) error {
	var request Request[IdeaListResponse]

	r, err := i.client.newRequest(http.MethodGet, "ideas", nil)
	if err != nil {
		panic(err)
	}

	if i.product != nil && i.product.ID != "" {
		r, err = i.client.newRequest(http.MethodGet, fmt.Sprintf("products/%s/ideas", i.product.ID), nil)
	}

	request.request = r
	return request.ForEach(func(resp *IdeaListResponse) {
		*ideas = append(*ideas, resp.Ideas...)
	})
}

//
// Release endpoints
//

type ReleaseAPIRequest struct {
	client  *AhaClient
	product *Product
}

func (c *AhaClient) Releases() *ReleaseAPIRequest {
	return &ReleaseAPIRequest{client: c}
}

type ReleaseResponse struct {
	Release Release `json:"release"`
}

func (r *ReleaseAPIRequest) First(release *Release) error {
	var request Request[ReleaseResponse]

	req, err := r.client.newRequest(http.MethodGet, fmt.Sprintf("releases/%s", release.ID), nil)
	if err != nil {
		return err
	}

	request.request = req
	x, err := request.Do()

	if err == nil && x != nil {
		*release = x.Release
	}
	return err
}

func (r *ReleaseAPIRequest) For(product *Product) *ReleaseAPIRequest {
	if r.product != nil && product.ID != "" {
		r.product = product
	}
	return r
}

type ReleaseListResponse struct {
	Releases []Release `json:"releases"`
}

func (r *ReleaseAPIRequest) List(releases *[]Release) error {
	var request Request[ReleaseListResponse]

	req, err := r.client.newRequest(http.MethodGet, "releases", nil)
	if err != nil {
		return err
	}

	if r.product != nil && r.product.ID != "" {
		req, err = r.client.newRequest(http.MethodGet, fmt.Sprintf("products/%s/releases", r.product.ID), nil)
	}

	request.request = req
	return request.ForEach(func(resp *ReleaseListResponse) {
		*releases = append(*releases, resp.Releases...)
	})
}

//
// Comments
//

type CommentAPIRequest struct {
	Error       error
	client      *AhaClient
	urlFragment string
}

func (c *AhaClient) Comments() *CommentAPIRequest {
	return &CommentAPIRequest{client: c}
}

func (c *CommentAPIRequest) For(resource any) *CommentAPIRequest {
	switch t := resource.(type) {
	case *Idea:
		c.urlFragment = fmt.Sprintf("/ideas/%s/comments", t.ID)
	case *Feature:
		c.urlFragment = fmt.Sprintf("/features/%s/comments", t.ID)
	case *Epic:
		c.urlFragment = fmt.Sprintf("/epics/%s/comments", t.ID)
	default:
		c.Error = fmt.Errorf("unknown resource type: %T", resource)
	}
	return c
}

type CommentRequest struct {
	Comment Comment `json:"comment"`
}
type CommentResponse CommentRequest

func (c *CommentAPIRequest) Create(comment *Comment) error {
	if c.Error != nil {
		return c.Error
	}

	if c.urlFragment == "" {
		return errors.New("cannot create comment without first calling For()")
	}

	b, err := json.Marshal(&CommentRequest{Comment: *comment})
	if err != nil {
		return err
	}

	r, err := c.client.newRequest(http.MethodPost, c.urlFragment, bytes.NewBuffer(b))
	if err != nil {
		return err
	}

	var request Request[CommentResponse]
	request.request = r
	x, err := request.Do()

	if err == nil && x != nil {
		*comment = x.Comment
	}
	return err
}

type CommentListResponse struct {
	Comments []Comment `json:"comments"`
}

func (c *CommentAPIRequest) List(comments *[]Comment) error {
	if c.Error != nil {
		return c.Error
	}

	if c.urlFragment == "" {
		return errors.New("cannot list comments without first calling For()")
	}

	r, err := c.client.newRequest(http.MethodGet, c.urlFragment, nil)
	if err != nil {
		return err
	}

	var request Request[CommentListResponse]
	request.request = r
	return request.ForEach(func(resp *CommentListResponse) {
		*comments = append(*comments, resp.Comments...)
	})
}
