package clients

import (
	"fmt"
	"net/http"
	"net/url"
)

type vultrClient struct {
	http.Client
	baseURL string
}

func NewVultrClient(baseURL string) *vultrClient {
	return &vultrClient{http.Client{}, baseURL}
}

func (c *vultrClient) Get(url string) (resp *http.Response, err error) {
	return c.Client.Get(fmt.Sprintf("%s%s", c.baseURL, url))
}

func (c *vultrClient) PostForm(url string,
	data url.Values) (resp *http.Response, err error) {
	return c.Client.PostForm(fmt.Sprintf("%s%s", c.baseURL, url), data)
}
