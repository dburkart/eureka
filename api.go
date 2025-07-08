package eureka

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func (c *AhaClient) ListProducts() *Request[ProductListResponse] {
	var request Request[ProductListResponse]

	r, err := c.newRequest(http.MethodGet, "products", nil)
	if err != nil {
		panic(err)
	}

	request.request = r
	return &request
}

//
// Idea endpoints
//

type ListIdeasOptions struct {
	ProductId *string `json:"product_id,omitempty"`
}

func (c *AhaClient) ListIdeas(options *ListIdeasOptions) *Request[IdeaListResponse] {
	var request Request[IdeaListResponse]

	r, err := c.newRequest(http.MethodGet, "ideas", nil)
	if err != nil {
		panic(err)
	}

	if options.ProductId != nil {
		r, err = c.newRequest(http.MethodGet, fmt.Sprintf("products/%s/ideas", *options.ProductId), nil)
	}

	request.request = r

	return &request
}
