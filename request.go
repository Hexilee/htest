package htest

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type (
	Request struct {
		*http.Request
		Handler http.Handler
		*testing.T
	}
)

func (r *Request) SetHeader(key, value string) *Request {
	r.Header.Set(key, value)
	return r
}

func (r *Request) SetHeaders(headers map[string]string) *Request {
	var key, value string
	for key, value = range headers {
		r.Header.Set(key, value)
	}
	return r
}

func (r *Request) Test() *Response {
	recorder := httptest.NewRecorder()
	r.Handler.ServeHTTP(recorder, r.Request)
	return NewResponse(recorder.Result(), r.T)
}

func (r *Request) Send() *Response {
	resp, err := (&http.Client{}).Do(r.Request)
	assert.Nil(r.T, err)
	return NewResponse(resp, r.T)
}
