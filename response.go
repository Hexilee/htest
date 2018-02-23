package htest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"crypto/md5"
	"io"
	"crypto/sha1"
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

func (r *Response) XML() *XML {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return NewXML(body, r.T)
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

func (r *Response) Expect(expect string) {
	assert.Equal(r.T, expect, r.String())
}

func (r *Response) MD5() *MD5 {
	buf := md5.New()
	io.Copy(buf, r.Response.Body)
	result := buf.Sum(nil)
	return NewMD5(result, r.T)
}

func (r *Response) SHA1() *SHA1 {
	buf := sha1.New()
	io.Copy(buf, r.Response.Body)
	result := buf.Sum(nil)
	return NewSHA1(result, r.T)
}

func (r *Response) Bind(obj interface{}) error {
	body, err := ioutil.ReadAll(r.Response.Body)
	assert.Nil(r.T, err)
	return json.Unmarshal(body, obj)
}

func (r *Response) Headers(key, expect string) *Response {
	assert.Equal(r.T, expect, r.Header.Get(key))
	return r
}

func (r *Response) HeaderAccept(expect string) *Response {
	return r.Headers(HeaderAccept, expect)
}

func (r *Response) HeaderAcceptEncoding(expect string) *Response {
	return r.Headers(HeaderAcceptEncoding, expect)
}

func (r *Response) HeaderAllow(expect string) *Response {
	return r.Headers(HeaderAllow, expect)
}

func (r *Response) HeaderAuthorization(expect string) *Response {
	return r.Headers(HeaderAuthorization, expect)
}

func (r *Response) HeaderContentDisposition(expect string) *Response {
	return r.Headers(HeaderContentDisposition, expect)
}

func (r *Response) HeaderContentEncoding(expect string) *Response {
	return r.Headers(HeaderContentEncoding, expect)
}

func (r *Response) HeaderContentLength(expect string) *Response {
	return r.Headers(HeaderContentLength, expect)
}

func (r *Response) HeaderContentType(expect string) *Response {
	return r.Headers(HeaderContentType, expect)
}

func (r *Response) HeaderCookie(expect string) *Response {
	return r.Headers(HeaderCookie, expect)
}

func (r *Response) HeaderSetCookie(expect string) *Response {
	return r.Headers(HeaderSetCookie, expect)
}

func (r *Response) HeaderIfModifiedSince(expect string) *Response {
	return r.Headers(HeaderIfModifiedSince, expect)
}

func (r *Response) HeaderLastModified(expect string) *Response {
	return r.Headers(HeaderLastModified, expect)
}

func (r *Response) HeaderLocation(expect string) *Response {
	return r.Headers(HeaderLocation, expect)
}

func (r *Response) HeaderUpgrade(expect string) *Response {
	return r.Headers(HeaderUpgrade, expect)
}

func (r *Response) HeaderVary(expect string) *Response {
	return r.Headers(HeaderVary, expect)
}

func (r *Response) HeaderWWWAuthenticate(expect string) *Response {
	return r.Headers(HeaderWWWAuthenticate, expect)
}

func (r *Response) HeaderXForwardedFor(expect string) *Response {
	return r.Headers(HeaderXForwardedFor, expect)
}

func (r *Response) HeaderXForwardedProto(expect string) *Response {
	return r.Headers(HeaderXForwardedProto, expect)
}

func (r *Response) HeaderXForwardedProtocol(expect string) *Response {
	return r.Headers(HeaderXForwardedProtocol, expect)
}

func (r *Response) HeaderXForwardedSsl(expect string) *Response {
	return r.Headers(HeaderXForwardedSsl, expect)
}

func (r *Response) HeaderXUrlScheme(expect string) *Response {
	return r.Headers(HeaderXUrlScheme, expect)
}

func (r *Response) HeaderXHTTPMethodOverride(expect string) *Response {
	return r.Headers(HeaderXHTTPMethodOverride, expect)
}

func (r *Response) HeaderXRealIP(expect string) *Response {
	return r.Headers(HeaderXRealIP, expect)
}

func (r *Response) HeaderXRequestID(expect string) *Response {
	return r.Headers(HeaderXRequestID, expect)
}

func (r *Response) HeaderServer(expect string) *Response {
	return r.Headers(HeaderServer, expect)
}

func (r *Response) HeaderOrigin(expect string) *Response {
	return r.Headers(HeaderOrigin, expect)
}

func (r *Response) HeaderAccessControlRequestMethod(expect string) *Response {
	return r.Headers(HeaderAccessControlRequestMethod, expect)
}

func (r *Response) HeaderAccessControlRequestHeaders(expect string) *Response {
	return r.Headers(HeaderAccessControlRequestHeaders, expect)
}

func (r *Response) HeaderAccessControlAllowOrigin(expect string) *Response {
	return r.Headers(HeaderAccessControlAllowOrigin, expect)
}

func (r *Response) HeaderAccessControlAllowMethods(expect string) *Response {
	return r.Headers(HeaderAccessControlAllowMethods, expect)
}

func (r *Response) HeaderAccessControlAllowHeaders(expect string) *Response {
	return r.Headers(HeaderAccessControlAllowHeaders, expect)
}

func (r *Response) HeaderAccessControlAllowCredentials(expect string) *Response {
	return r.Headers(HeaderAccessControlAllowCredentials, expect)
}

func (r *Response) HeaderAccessControlExposeHeaders(expect string) *Response {
	return r.Headers(HeaderAccessControlExposeHeaders, expect)
}

func (r *Response) HeaderAccessControlMaxAge(expect string) *Response {
	return r.Headers(HeaderAccessControlMaxAge, expect)
}

func (r *Response) HeaderStrictTransportSecurity(expect string) *Response {
	return r.Headers(HeaderStrictTransportSecurity, expect)
}

func (r *Response) HeaderXContentTypeOptions(expect string) *Response {
	return r.Headers(HeaderXContentTypeOptions, expect)
}

func (r *Response) HeaderXXSSProtection(expect string) *Response {
	return r.Headers(HeaderXXSSProtection, expect)
}

func (r *Response) HeaderXFrameOptions(expect string) *Response {
	return r.Headers(HeaderXFrameOptions, expect)
}

func (r *Response) HeaderContentSecurityPolicy(expect string) *Response {
	return r.Headers(HeaderContentSecurityPolicy, expect)
}

func (r *Response) HeaderXCSRFToken(expect string) *Response {
	return r.Headers(HeaderXCSRFToken, expect)
}
