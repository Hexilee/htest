package htest

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

type (
	User struct {
		Id   uint   `xml:"id"`
		Name string `xml:"name"`
	}
)

const (
	UserData = `{
	"id": 1,
	"name": "hexi"
}`
	UserDataXML = `
<?xml version="1.0" encoding="UTF-8"?>
<user>
	<id>1</id>
	<name>hexi</name>
</user>
`
)

var (
	ResponseCodeServer    = chi.NewRouter()
	ResponseHeadersServer = chi.NewRouter()
	UserDataMD5           string
	UserDataSHA1          string
)

func init() {
	ResponseCodeServer.Get("/response/statusCode/{code}", StatusHandler)
	ResponseHeadersServer.Get("/response/headers", HeadersHandler)

	UserMD5 := md5.New()
	UserMD5.Write([]byte(UserData))
	UserDataMD5 = string(UserMD5.Sum(nil))

	UserSHA1 := sha1.New()
	UserSHA1.Write([]byte(UserData))
	UserDataSHA1 = string(UserSHA1.Sum(nil))

}

func TestResponse_String(t *testing.T) {
	client := NewClient(t).To(Mux)
	assert.Equal(t, UserData, client.Get("/body/user").Test().StatusOK().String())
}

func TestResponse_Bytes(t *testing.T) {
	client := NewClient(t).To(Mux)
	assert.Equal(t, []byte(UserData), client.Get("/body/user").Test().StatusOK().Bytes())
}

func TestResponse_Expect(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Test().StatusOK().Expect(UserData)
}

func TestResponse_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Test().StatusOK().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func TestResponse_Code(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadRequest)).Test().Code(http.StatusBadRequest)
}

func TestResponse_Headers(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentType, MIMEApplicationJSON)
	client.Get(url).Test().Headers(HeaderContentType, MIMEApplicationJSON)
}

func TestResponse_StatusContinue(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusContinue)).Test().StatusContinue()
}

func TestResponse_StatusSwitchingProtocols(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusSwitchingProtocols)).Test().StatusSwitchingProtocols()
}

func TestResponse_StatusProcessing(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusProcessing)).Test().StatusProcessing()
}

func TestResponse_StatusOK(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusOK)).Test().StatusOK()
}

func TestResponse_StatusCreated(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusCreated)).Test().StatusCreated()
}

func TestResponse_StatusAccepted(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusAccepted)).Test().StatusAccepted()
}

func TestResponse_StatusNonAuthoritativeInfo(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNonAuthoritativeInfo)).Test().StatusNonAuthoritativeInfo()
}

func TestResponse_StatusNoContent(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNoContent)).Test().StatusNoContent()
}

func TestResponse_StatusResetContent(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusResetContent)).Test().StatusResetContent()
}

func TestResponse_StatusPartialContent(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPartialContent)).Test().StatusPartialContent()
}

func TestResponse_StatusMultiStatus(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMultiStatus)).Test().StatusMultiStatus()
}

func TestResponse_StatusAlreadyReported(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusAlreadyReported)).Test().StatusAlreadyReported()
}

func TestResponse_StatusIMUsed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusIMUsed)).Test().StatusIMUsed()
}

func TestResponse_StatusMultipleChoices(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMultipleChoices)).Test().StatusMultipleChoices()
}

func TestResponse_StatusMovedPermanently(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMovedPermanently)).Test().StatusMovedPermanently()
}

func TestResponse_StatusFound(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusFound)).Test().StatusFound()
}

func TestResponse_StatusSeeOther(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusSeeOther)).Test().StatusSeeOther()
}

func TestResponse_StatusNotModified(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotModified)).Test().StatusNotModified()
}

func TestResponse_StatusUseProxy(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUseProxy)).Test().StatusUseProxy()
}

func TestResponse_StatusTemporaryRedirect(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusTemporaryRedirect)).Test().StatusTemporaryRedirect()
}

func TestResponse_StatusPermanentRedirect(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPermanentRedirect)).Test().StatusPermanentRedirect()
}

func TestResponse_StatusBadRequest(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadRequest)).Test().StatusBadRequest()
}

func TestResponse_StatusUnauthorized(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnauthorized)).Test().StatusUnauthorized()
}

func TestResponse_StatusPaymentRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPaymentRequired)).Test().StatusPaymentRequired()
}

func TestResponse_StatusForbidden(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusForbidden)).Test().StatusForbidden()
}

func TestResponse_StatusNotFound(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotFound)).Test().StatusNotFound()
}

func TestResponse_StatusMethodNotAllowed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMethodNotAllowed)).Test().StatusMethodNotAllowed()
}

func TestResponse_StatusNotAcceptable(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotAcceptable)).Test().StatusNotAcceptable()
}

func TestResponse_StatusProxyAuthRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusProxyAuthRequired)).Test().StatusProxyAuthRequired()
}

func TestResponse_StatusRequestTimeout(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestTimeout)).Test().StatusRequestTimeout()
}

func TestResponse_StatusConflict(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusConflict)).Test().StatusConflict()
}

func TestResponse_StatusGone(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusGone)).Test().StatusGone()
}

func TestResponse_StatusLengthRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusLengthRequired)).Test().StatusLengthRequired()
}

func TestResponse_StatusPreconditionFailed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPreconditionFailed)).Test().StatusPreconditionFailed()
}

func TestResponse_StatusRequestEntityTooLarge(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestEntityTooLarge)).Test().StatusRequestEntityTooLarge()
}

func TestResponse_StatusRequestURITooLong(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestURITooLong)).Test().StatusRequestURITooLong()
}

func TestResponse_StatusUnsupportedMediaType(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnsupportedMediaType)).Test().StatusUnsupportedMediaType()
}

func TestResponse_StatusRequestedRangeNotSatisfiable(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestedRangeNotSatisfiable)).Test().StatusRequestedRangeNotSatisfiable()
}

func TestResponse_StatusExpectationFailed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusExpectationFailed)).Test().StatusExpectationFailed()
}

func TestResponse_StatusTeapot(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusTeapot)).Test().StatusTeapot()
}

func TestResponse_StatusUnprocessableEntity(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnprocessableEntity)).Test().StatusUnprocessableEntity()
}

func TestResponse_StatusLocked(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusLocked)).Test().StatusLocked()
}

func TestResponse_StatusFailedDependency(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusFailedDependency)).Test().StatusFailedDependency()
}

func TestResponse_StatusUpgradeRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUpgradeRequired)).Test().StatusUpgradeRequired()
}

func TestResponse_StatusPreconditionRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPreconditionRequired)).Test().StatusPreconditionRequired()
}

func TestResponse_StatusTooManyRequests(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusTooManyRequests)).Test().StatusTooManyRequests()
}

func TestResponse_StatusRequestHeaderFieldsTooLarge(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestHeaderFieldsTooLarge)).Test().StatusRequestHeaderFieldsTooLarge()
}

func TestResponse_StatusUnavailableForLegalReasons(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnavailableForLegalReasons)).Test().StatusUnavailableForLegalReasons()
}

func TestResponse_StatusInternalServerError(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusInternalServerError)).Test().StatusInternalServerError()
}

func TestResponse_StatusNotImplemented(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotImplemented)).Test().StatusNotImplemented()
}

func TestResponse_StatusBadGateway(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadGateway)).Test().StatusBadGateway()
}

func TestResponse_StatusServiceUnavailable(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusServiceUnavailable)).Test().StatusServiceUnavailable()
}

func TestResponse_StatusGatewayTimeout(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusGatewayTimeout)).Test().StatusGatewayTimeout()
}

func TestResponse_StatusHTTPVersionNotSupported(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusHTTPVersionNotSupported)).Test().StatusHTTPVersionNotSupported()
}

func TestResponse_StatusVariantAlsoNegotiates(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusVariantAlsoNegotiates)).Test().StatusVariantAlsoNegotiates()
}

func TestResponse_StatusInsufficientStorage(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusInsufficientStorage)).Test().StatusInsufficientStorage()
}

func TestResponse_StatusLoopDetected(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusLoopDetected)).Test().StatusLoopDetected()
}

func TestResponse_StatusNotExtended(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotExtended)).Test().StatusNotExtended()
}

func TestResponse_StatusNetworkAuthenticationRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNetworkAuthenticationRequired)).Test().StatusNetworkAuthenticationRequired()
}

func TestResponse_HeaderAccept(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccept, "htest")
	client.Get(url).Test().HeaderAccept("htest")
}

func TestResponse_HeaderAcceptEncoding(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAcceptEncoding, "htest")
	client.Get(url).Test().HeaderAcceptEncoding("htest")
}

func TestResponse_HeaderAllow(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAllow, "htest")
	client.Get(url).Test().HeaderAllow("htest")
}

func TestResponse_HeaderAuthorization(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAuthorization, "htest")
	client.Get(url).Test().HeaderAuthorization("htest")
}

func TestResponse_HeaderContentDisposition(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentDisposition, "htest")
	client.Get(url).Test().HeaderContentDisposition("htest")
}

func TestResponse_HeaderContentEncoding(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentEncoding, "htest")
	client.Get(url).Test().HeaderContentEncoding("htest")
}

func TestResponse_HeaderContentLength(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentLength, "htest")
	client.Get(url).Test().HeaderContentLength("htest")
}

func TestResponse_HeaderContentType(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentType, "htest")
	client.Get(url).Test().HeaderContentType("htest")
}

func TestResponse_HeaderCookie(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderCookie, "htest")
	client.Get(url).Test().HeaderCookie("htest")
}

func TestResponse_HeaderSetCookie(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderSetCookie, "htest")
	client.Get(url).Test().HeaderSetCookie("htest")
}

func TestResponse_HeaderIfModifiedSince(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderIfModifiedSince, "htest")
	client.Get(url).Test().HeaderIfModifiedSince("htest")
}

func TestResponse_HeaderLastModified(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderLastModified, "htest")
	client.Get(url).Test().HeaderLastModified("htest")
}

func TestResponse_HeaderLocation(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderLocation, "htest")
	client.Get(url).Test().HeaderLocation("htest")
}

func TestResponse_HeaderUpgrade(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderUpgrade, "htest")
	client.Get(url).Test().HeaderUpgrade("htest")
}

func TestResponse_HeaderVary(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderVary, "htest")
	client.Get(url).Test().HeaderVary("htest")
}

func TestResponse_HeaderWWWAuthenticate(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderWWWAuthenticate, "htest")
	client.Get(url).Test().HeaderWWWAuthenticate("htest")
}

func TestResponse_HeaderXForwardedFor(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedFor, "htest")
	client.Get(url).Test().HeaderXForwardedFor("htest")
}

func TestResponse_HeaderXForwardedProto(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedProto, "htest")
	client.Get(url).Test().HeaderXForwardedProto("htest")
}

func TestResponse_HeaderXForwardedProtocol(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedProtocol, "htest")
	client.Get(url).Test().HeaderXForwardedProtocol("htest")
}

func TestResponse_HeaderXForwardedSsl(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedSsl, "htest")
	client.Get(url).Test().HeaderXForwardedSsl("htest")
}

func TestResponse_HeaderXUrlScheme(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXUrlScheme, "htest")
	client.Get(url).Test().HeaderXUrlScheme("htest")
}

func TestResponse_HeaderXHTTPMethodOverride(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXHTTPMethodOverride, "htest")
	client.Get(url).Test().HeaderXHTTPMethodOverride("htest")
}

func TestResponse_HeaderXRealIP(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXRealIP, "htest")
	client.Get(url).Test().HeaderXRealIP("htest")
}

func TestResponse_HeaderXRequestID(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXRequestID, "htest")
	client.Get(url).Test().HeaderXRequestID("htest")
}

func TestResponse_HeaderServer(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderServer, "htest")
	client.Get(url).Test().HeaderServer("htest")
}

func TestResponse_HeaderOrigin(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderOrigin, "htest")
	client.Get(url).Test().HeaderOrigin("htest")
}

func TestResponse_HeaderAccessControlRequestMethod(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlRequestMethod, "htest")
	client.Get(url).Test().HeaderAccessControlRequestMethod("htest")
}

func TestResponse_HeaderAccessControlRequestHeaders(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlRequestHeaders, "htest")
	client.Get(url).Test().HeaderAccessControlRequestHeaders("htest")
}

func TestResponse_HeaderAccessControlAllowOrigin(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowOrigin, "htest")
	client.Get(url).Test().HeaderAccessControlAllowOrigin("htest")
}

func TestResponse_HeaderAccessControlAllowMethods(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowMethods, "htest")
	client.Get(url).Test().HeaderAccessControlAllowMethods("htest")
}

func TestResponse_HeaderAccessControlAllowHeaders(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowHeaders, "htest")
	client.Get(url).Test().HeaderAccessControlAllowHeaders("htest")
}

func TestResponse_HeaderAccessControlAllowCredentials(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowCredentials, "htest")
	client.Get(url).Test().HeaderAccessControlAllowCredentials("htest")
}

func TestResponse_HeaderAccessControlExposeHeaders(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlExposeHeaders, "htest")
	client.Get(url).Test().HeaderAccessControlExposeHeaders("htest")
}

func TestResponse_HeaderAccessControlMaxAge(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlMaxAge, "htest")
	client.Get(url).Test().HeaderAccessControlMaxAge("htest")
}

func TestResponse_HeaderStrictTransportSecurity(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderStrictTransportSecurity, "htest")
	client.Get(url).Test().HeaderStrictTransportSecurity("htest")
}

func TestResponse_HeaderXContentTypeOptions(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXContentTypeOptions, "htest")
	client.Get(url).Test().HeaderXContentTypeOptions("htest")
}

func TestResponse_HeaderXXSSProtection(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXXSSProtection, "htest")
	client.Get(url).Test().HeaderXXSSProtection("htest")
}

func TestResponse_HeaderXFrameOptions(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXFrameOptions, "htest")
	client.Get(url).Test().HeaderXFrameOptions("htest")
}

func TestResponse_HeaderContentSecurityPolicy(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentSecurityPolicy, "htest")
	client.Get(url).Test().HeaderContentSecurityPolicy("htest")
}

func TestResponse_HeaderXCSRFToken(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXCSRFToken, "htest")
	client.Get(url).Test().HeaderXCSRFToken("htest")
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	codeStr := chi.URLParam(req, "code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(code)
}

func HeadersHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	header := query.Get("header")
	value := query.Get("value")
	w.Header().Set(header, value)
}
