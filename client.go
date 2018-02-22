package htest

import (
	"net/http"
)

type (
	Client struct {
		handler http.Handler
	}
)

func NewClient() *Client {
	return &Client{}
}

func (c Client) To(handler http.Handler) *Client {
	c.handler = handler
	return &c
}

func (c Client) ToFunc(handlerFunc http.HandlerFunc) *Client {
	c.handler = handlerFunc
	return &c
}

func (c Client) NewRequest(req *http.Request) *Request {
	return &Request{
		Request: req,
		Handler: c.handler,
	}
}

func (c Client) Get(path string) *Request {
	req, _ := http.NewRequest(GET, path, nil)
	return c.NewRequest(req)
}
