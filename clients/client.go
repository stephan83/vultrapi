package clients

import (
	"net/http"
	"net/url"
)

type Client interface {
	Get(url string) (resp *http.Response, err error)
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}
