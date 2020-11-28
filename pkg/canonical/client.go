package canonical

import (
	"fmt"
	"golang.org/x/net/publicsuffix"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// URLs
const DefaultBaseURL = "https://api.mercadolibre.com/"

// endpoints.
const (
	items = "items/"
)

type Client struct {
	client  *http.Client
	BaseURL *url.URL
	common  service // Reuse a single struct instead of allocating one for each service on the heap.
	Meli    *MeliService
}

// service holds a reference to the client
type service struct {
	client *Client
}

func NewClient(hc *http.Client) *Client {
	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	if hc == nil {
		hc = &http.Client{}
	}
	hc.Jar = jar

	baseURL, _ := url.Parse(DefaultBaseURL)

	c := &Client{client: hc, BaseURL: baseURL}
	c.common.client = c
	c.Meli = (*MeliService)(&c.common)
	return c
}

func (c *Client) NewRequest(method, path string, body io.Reader, token string) (*http.Request, error) {
	if !strings.HasSuffix(c.BaseURL.Path, "/") {
		return nil, fmt.Errorf("BaseURL must have a trailing slash, but %q does not.", c.BaseURL)
	}

	u, err := c.BaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	if token != "" {
		q := u.Query()
		q.Set("access_token", token) // Add token.

		u.RawQuery = q.Encode()
	}

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func (c *Client) Do(req *http.Request) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}
