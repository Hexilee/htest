package htest

import (
	"encoding/json"
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

func NewResponse(response *http.Response, t *testing.T) *Response {
	return &Response{
		Response: response,
		T:        t,
	}
}

func (r *Response) Code(statusCode int) *Response {
	assert.Equal(r.T, statusCode, r.StatusCode)
	return r
}

func (r *Response) StatusContinue() *Response {
	return r.Code(http.StatusContinue)
}
func (r *Response) StatusSwitchingProtocols() *Response {
	return r.Code(http.StatusSwitchingProtocols)
}
func (r *Response) StatusProcessing() *Response {
	return r.Code(http.StatusProcessing)
}
func (r *Response) StatusOK() *Response {
	return r.Code(http.StatusOK)
}
func (r *Response) StatusCreated() *Response {
	return r.Code(http.StatusCreated)
}
func (r *Response) StatusAccepted() *Response {
	return r.Code(http.StatusAccepted)
}
func (r *Response) StatusNonAuthoritativeInfo() *Response {
	return r.Code(http.StatusNonAuthoritativeInfo)
}
func (r *Response) StatusNoContent() *Response {
	return r.Code(http.StatusNoContent)
}
func (r *Response) StatusResetContent() *Response {
	return r.Code(http.StatusResetContent)
}
func (r *Response) StatusPartialContent() *Response {
	return r.Code(http.StatusPartialContent)
}
func (r *Response) StatusMultiStatus() *Response {
	return r.Code(http.StatusMultiStatus)
}
func (r *Response) StatusAlreadyReported() *Response {
	return r.Code(http.StatusAlreadyReported)
}
func (r *Response) StatusIMUsed() *Response {
	return r.Code(http.StatusIMUsed)
}
func (r *Response) StatusMultipleChoices() *Response {
	return r.Code(http.StatusMultipleChoices)
}
func (r *Response) StatusMovedPermanently() *Response {
	return r.Code(http.StatusMovedPermanently)
}
func (r *Response) StatusFound() *Response {
	return r.Code(http.StatusFound)
}
func (r *Response) StatusSeeOther() *Response {
	return r.Code(http.StatusSeeOther)
}
func (r *Response) StatusNotModified() *Response {
	return r.Code(http.StatusNotModified)
}
func (r *Response) StatusUseProxy() *Response {
	return r.Code(http.StatusUseProxy)
}
func (r *Response) StatusTemporaryRedirect() *Response {
	return r.Code(http.StatusTemporaryRedirect)
}
func (r *Response) StatusPermanentRedirect() *Response {
	return r.Code(http.StatusPermanentRedirect)
}
func (r *Response) StatusBadRequest() *Response {
	return r.Code(http.StatusBadRequest)
}
func (r *Response) StatusUnauthorized() *Response {
	return r.Code(http.StatusUnauthorized)
}
func (r *Response) StatusPaymentRequired() *Response {
	return r.Code(http.StatusPaymentRequired)
}
func (r *Response) StatusForbidden() *Response {
	return r.Code(http.StatusForbidden)
}
func (r *Response) StatusNotFound() *Response {
	return r.Code(http.StatusNotFound)
}
func (r *Response) StatusMethodNotAllowed() *Response {
	return r.Code(http.StatusMethodNotAllowed)
}
func (r *Response) StatusNotAcceptable() *Response {
	return r.Code(http.StatusNotAcceptable)
}
func (r *Response) StatusProxyAuthRequired() *Response {
	return r.Code(http.StatusProxyAuthRequired)
}
func (r *Response) StatusRequestTimeout() *Response {
	return r.Code(http.StatusRequestTimeout)
}
func (r *Response) StatusConflict() *Response {
	return r.Code(http.StatusConflict)
}
func (r *Response) StatusGone() *Response {
	return r.Code(http.StatusGone)
}
func (r *Response) StatusLengthRequired() *Response {
	return r.Code(http.StatusLengthRequired)
}
func (r *Response) StatusPreconditionFailed() *Response {
	return r.Code(http.StatusPreconditionFailed)
}
func (r *Response) StatusRequestEntityTooLarge() *Response {
	return r.Code(http.StatusRequestEntityTooLarge)
}
func (r *Response) StatusRequestURITooLong() *Response {
	return r.Code(http.StatusRequestURITooLong)
}
func (r *Response) StatusUnsupportedMediaType() *Response {
	return r.Code(http.StatusUnsupportedMediaType)
}
func (r *Response) StatusRequestedRangeNotSatisfiable() *Response {
	return r.Code(http.StatusRequestedRangeNotSatisfiable)
}
func (r *Response) StatusExpectationFailed() *Response {
	return r.Code(http.StatusExpectationFailed)
}
func (r *Response) StatusTeapot() *Response {
	return r.Code(http.StatusTeapot)
}
func (r *Response) StatusUnprocessableEntity() *Response {
	return r.Code(http.StatusUnprocessableEntity)
}
func (r *Response) StatusLocked() *Response {
	return r.Code(http.StatusLocked)
}
func (r *Response) StatusFailedDependency() *Response {
	return r.Code(http.StatusFailedDependency)
}
func (r *Response) StatusUpgradeRequired() *Response {
	return r.Code(http.StatusUpgradeRequired)
}
func (r *Response) StatusPreconditionRequired() *Response {
	return r.Code(http.StatusPreconditionRequired)
}
func (r *Response) StatusTooManyRequests() *Response {
	return r.Code(http.StatusTooManyRequests)
}
func (r *Response) StatusRequestHeaderFieldsTooLarge() *Response {
	return r.Code(http.StatusRequestHeaderFieldsTooLarge)
}
func (r *Response) StatusUnavailableForLegalReasons() *Response {
	return r.Code(http.StatusUnavailableForLegalReasons)
}
func (r *Response) StatusInternalServerError() *Response {
	return r.Code(http.StatusInternalServerError)
}
func (r *Response) StatusNotImplemented() *Response {
	return r.Code(http.StatusNotImplemented)
}
func (r *Response) StatusBadGateway() *Response {
	return r.Code(http.StatusBadGateway)
}
func (r *Response) StatusServiceUnavailable() *Response {
	return r.Code(http.StatusServiceUnavailable)
}
func (r *Response) StatusGatewayTimeout() *Response {
	return r.Code(http.StatusGatewayTimeout)
}
func (r *Response) StatusHTTPVersionNotSupported() *Response {
	return r.Code(http.StatusHTTPVersionNotSupported)
}
func (r *Response) StatusVariantAlsoNegotiates() *Response {
	return r.Code(http.StatusVariantAlsoNegotiates)
}
func (r *Response) StatusInsufficientStorage() *Response {
	return r.Code(http.StatusInsufficientStorage)
}
func (r *Response) StatusLoopDetected() *Response {
	return r.Code(http.StatusLoopDetected)
}
func (r *Response) StatusNotExtended() *Response {
	return r.Code(http.StatusNotExtended)
}
func (r *Response) StatusNetworkAuthenticationRequired() *Response {
	return r.Code(http.StatusNetworkAuthenticationRequired)
}

func (r *Response) JSON() *JSON {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return NewJSON(body, r.T)
}

func (r *Response) Bytes() []byte {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return body
}

func (r *Response) String() string {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return string(body)
}

func (r *Response) Bind(obj interface{}) error {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return json.Unmarshal(body, obj)
}

func (r *Response) Headers(key, expect string) {
	assert.Equal(r.T, expect, r.Header.Get(key))
}
