package htest

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

type (
	Response struct {
		*http.Response
		*testing.T
	}
)

func NewResponse(response *http.Response) *Response {
	return &Response{
		Response: response,
	}
}

func (r *Response) OK() *Response {
	assert.Equal(r.T, http.StatusOK, r.StatusCode)
	return r
}

func (r *Response) BadRequest() *Response {
	assert.Equal(r.T, http.StatusBadRequest, r.StatusCode)
	return r
}

func (r *Response) JSON() *JSON {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return NewJSON(body, r.T)
}
