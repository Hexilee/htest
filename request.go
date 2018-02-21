package htest

import (
	"net/http"
	"net/http/httptest"
)

type (
	Request struct {
		*http.Request
		Handler http.Handler
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

func (r *Request) Send() *Response {
	recorder := httptest.NewRecorder()
	r.Handler.ServeHTTP(recorder, r.Request)
	return NewResponse(recorder.Result())
}
