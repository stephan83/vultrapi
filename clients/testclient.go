package clients

import (
	"bytes"
	"errors"
	"net/http"
	"net/url"
	"regexp"
)

type testClient struct {
	resp http.Response
}

func NewTestClient(statusCode int, body []byte) *testClient {
	return &testClient{
		http.Response{
			StatusCode: statusCode,
			Body:       testBody{bytes.NewReader(body)},
		},
	}
}

func (c *testClient) Get(url string) (resp *http.Response, err error) {
	return &c.resp, nil
}

func (c *testClient) PostForm(url string,
	data url.Values) (resp *http.Response, err error) {
	return &c.resp, nil
}

type testBody struct {
	*bytes.Reader
}

func (_ testBody) Close() error {
	return nil
}

type testMultiClient struct {
	router map[string]Client
}

func NewTestMultiClient(router map[string]Client) testMultiClient {
	return testMultiClient{router}
}

func (c testMultiClient) Get(url string) (resp *http.Response, err error) {
	for r, c := range c.router {
		if regexp.MustCompile(r).MatchString(url) {
			return c.Get(url)
		}
	}

	return nil, errors.New("Not found.")
}

func (c testMultiClient) PostForm(url string,
	data url.Values) (resp *http.Response, err error) {
	for r, c := range c.router {
		if regexp.MustCompile(r).MatchString(url) {
			return c.PostForm(url, data)
		}
	}

	return nil, errors.New("Not found.")
}
