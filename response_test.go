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
	assert.Equal(t, UserData, client.Get("/body/user").Send().StatusOK().String())
}

func TestResponse_Bytes(t *testing.T) {
	client := NewClient(t).To(Mux)
	assert.Equal(t, []byte(UserData), client.Get("/body/user").Send().StatusOK().Bytes())
}

func TestResponse_Expect(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Send().StatusOK().Expect(UserData)
}

func TestResponse_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Send().StatusOK().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func TestResponse_Code(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadRequest)).Send().Code(http.StatusBadRequest)
}

func TestResponse_Headers(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentType, MIMEApplicationJSON)
	client.Get(url).Send().Headers(HeaderContentType, MIMEApplicationJSON)
}

func TestResponse_StatusContinue(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusContinue)).Send().StatusContinue()
}

func TestResponse_StatusSwitchingProtocols(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusSwitchingProtocols)).Send().StatusSwitchingProtocols()
}

func TestResponse_StatusProcessing(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusProcessing)).Send().StatusProcessing()
}

func TestResponse_StatusOK(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusOK)).Send().StatusOK()
}

func TestResponse_StatusCreated(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusCreated)).Send().StatusCreated()
}

func TestResponse_StatusAccepted(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusAccepted)).Send().StatusAccepted()
}

func TestResponse_StatusNonAuthoritativeInfo(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNonAuthoritativeInfo)).Send().StatusNonAuthoritativeInfo()
}

func TestResponse_StatusNoContent(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNoContent)).Send().StatusNoContent()
}

func TestResponse_StatusResetContent(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusResetContent)).Send().StatusResetContent()
}

func TestResponse_StatusPartialContent(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPartialContent)).Send().StatusPartialContent()
}

func TestResponse_StatusMultiStatus(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMultiStatus)).Send().StatusMultiStatus()
}

func TestResponse_StatusAlreadyReported(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusAlreadyReported)).Send().StatusAlreadyReported()
}

func TestResponse_StatusIMUsed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusIMUsed)).Send().StatusIMUsed()
}

func TestResponse_StatusMultipleChoices(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMultipleChoices)).Send().StatusMultipleChoices()
}

func TestResponse_StatusMovedPermanently(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMovedPermanently)).Send().StatusMovedPermanently()
}

func TestResponse_StatusFound(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusFound)).Send().StatusFound()
}

func TestResponse_StatusSeeOther(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusSeeOther)).Send().StatusSeeOther()
}

func TestResponse_StatusNotModified(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotModified)).Send().StatusNotModified()
}

func TestResponse_StatusUseProxy(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUseProxy)).Send().StatusUseProxy()
}

func TestResponse_StatusTemporaryRedirect(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusTemporaryRedirect)).Send().StatusTemporaryRedirect()
}

func TestResponse_StatusPermanentRedirect(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPermanentRedirect)).Send().StatusPermanentRedirect()
}

func TestResponse_StatusBadRequest(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadRequest)).Send().StatusBadRequest()
}

func TestResponse_StatusUnauthorized(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnauthorized)).Send().StatusUnauthorized()
}

func TestResponse_StatusPaymentRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPaymentRequired)).Send().StatusPaymentRequired()
}

func TestResponse_StatusForbidden(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusForbidden)).Send().StatusForbidden()
}

func TestResponse_StatusNotFound(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotFound)).Send().StatusNotFound()
}

func TestResponse_StatusMethodNotAllowed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusMethodNotAllowed)).Send().StatusMethodNotAllowed()
}

func TestResponse_StatusNotAcceptable(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotAcceptable)).Send().StatusNotAcceptable()
}

func TestResponse_StatusProxyAuthRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusProxyAuthRequired)).Send().StatusProxyAuthRequired()
}

func TestResponse_StatusRequestTimeout(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestTimeout)).Send().StatusRequestTimeout()
}

func TestResponse_StatusConflict(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusConflict)).Send().StatusConflict()
}

func TestResponse_StatusGone(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusGone)).Send().StatusGone()
}

func TestResponse_StatusLengthRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusLengthRequired)).Send().StatusLengthRequired()
}

func TestResponse_StatusPreconditionFailed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPreconditionFailed)).Send().StatusPreconditionFailed()
}

func TestResponse_StatusRequestEntityTooLarge(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestEntityTooLarge)).Send().StatusRequestEntityTooLarge()
}

func TestResponse_StatusRequestURITooLong(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestURITooLong)).Send().StatusRequestURITooLong()
}

func TestResponse_StatusUnsupportedMediaType(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnsupportedMediaType)).Send().StatusUnsupportedMediaType()
}

func TestResponse_StatusRequestedRangeNotSatisfiable(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestedRangeNotSatisfiable)).Send().StatusRequestedRangeNotSatisfiable()
}

func TestResponse_StatusExpectationFailed(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusExpectationFailed)).Send().StatusExpectationFailed()
}

func TestResponse_StatusTeapot(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusTeapot)).Send().StatusTeapot()
}

func TestResponse_StatusUnprocessableEntity(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnprocessableEntity)).Send().StatusUnprocessableEntity()
}

func TestResponse_StatusLocked(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusLocked)).Send().StatusLocked()
}

func TestResponse_StatusFailedDependency(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusFailedDependency)).Send().StatusFailedDependency()
}

func TestResponse_StatusUpgradeRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUpgradeRequired)).Send().StatusUpgradeRequired()
}

func TestResponse_StatusPreconditionRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusPreconditionRequired)).Send().StatusPreconditionRequired()
}

func TestResponse_StatusTooManyRequests(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusTooManyRequests)).Send().StatusTooManyRequests()
}

func TestResponse_StatusRequestHeaderFieldsTooLarge(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusRequestHeaderFieldsTooLarge)).Send().StatusRequestHeaderFieldsTooLarge()
}

func TestResponse_StatusUnavailableForLegalReasons(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusUnavailableForLegalReasons)).Send().StatusUnavailableForLegalReasons()
}

func TestResponse_StatusInternalServerError(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusInternalServerError)).Send().StatusInternalServerError()
}

func TestResponse_StatusNotImplemented(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotImplemented)).Send().StatusNotImplemented()
}

func TestResponse_StatusBadGateway(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadGateway)).Send().StatusBadGateway()
}

func TestResponse_StatusServiceUnavailable(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusServiceUnavailable)).Send().StatusServiceUnavailable()
}

func TestResponse_StatusGatewayTimeout(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusGatewayTimeout)).Send().StatusGatewayTimeout()
}

func TestResponse_StatusHTTPVersionNotSupported(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusHTTPVersionNotSupported)).Send().StatusHTTPVersionNotSupported()
}

func TestResponse_StatusVariantAlsoNegotiates(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusVariantAlsoNegotiates)).Send().StatusVariantAlsoNegotiates()
}

func TestResponse_StatusInsufficientStorage(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusInsufficientStorage)).Send().StatusInsufficientStorage()
}

func TestResponse_StatusLoopDetected(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusLoopDetected)).Send().StatusLoopDetected()
}

func TestResponse_StatusNotExtended(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNotExtended)).Send().StatusNotExtended()
}

func TestResponse_StatusNetworkAuthenticationRequired(t *testing.T) {
	client := NewClient(t).To(ResponseCodeServer)
	client.Get(fmt.Sprintf("/response/statusCode/%d", http.StatusNetworkAuthenticationRequired)).Send().StatusNetworkAuthenticationRequired()
}

func TestResponse_HeaderAccept(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccept, "htest")
	client.Get(url).Send().HeaderAccept("htest")
}

func TestResponse_HeaderAcceptEncoding(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAcceptEncoding, "htest")
	client.Get(url).Send().HeaderAcceptEncoding("htest")
}

func TestResponse_HeaderAllow(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAllow, "htest")
	client.Get(url).Send().HeaderAllow("htest")
}

func TestResponse_HeaderAuthorization(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAuthorization, "htest")
	client.Get(url).Send().HeaderAuthorization("htest")
}

func TestResponse_HeaderContentDisposition(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentDisposition, "htest")
	client.Get(url).Send().HeaderContentDisposition("htest")
}

func TestResponse_HeaderContentEncoding(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentEncoding, "htest")
	client.Get(url).Send().HeaderContentEncoding("htest")
}

func TestResponse_HeaderContentLength(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentLength, "htest")
	client.Get(url).Send().HeaderContentLength("htest")
}

func TestResponse_HeaderContentType(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentType, "htest")
	client.Get(url).Send().HeaderContentType("htest")
}

func TestResponse_HeaderCookie(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderCookie, "htest")
	client.Get(url).Send().HeaderCookie("htest")
}

func TestResponse_HeaderSetCookie(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderSetCookie, "htest")
	client.Get(url).Send().HeaderSetCookie("htest")
}

func TestResponse_HeaderIfModifiedSince(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderIfModifiedSince, "htest")
	client.Get(url).Send().HeaderIfModifiedSince("htest")
}

func TestResponse_HeaderLastModified(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderLastModified, "htest")
	client.Get(url).Send().HeaderLastModified("htest")
}

func TestResponse_HeaderLocation(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderLocation, "htest")
	client.Get(url).Send().HeaderLocation("htest")
}

func TestResponse_HeaderUpgrade(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderUpgrade, "htest")
	client.Get(url).Send().HeaderUpgrade("htest")
}

func TestResponse_HeaderVary(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderVary, "htest")
	client.Get(url).Send().HeaderVary("htest")
}

func TestResponse_HeaderWWWAuthenticate(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderWWWAuthenticate, "htest")
	client.Get(url).Send().HeaderWWWAuthenticate("htest")
}

func TestResponse_HeaderXForwardedFor(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedFor, "htest")
	client.Get(url).Send().HeaderXForwardedFor("htest")
}

func TestResponse_HeaderXForwardedProto(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedProto, "htest")
	client.Get(url).Send().HeaderXForwardedProto("htest")
}

func TestResponse_HeaderXForwardedProtocol(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedProtocol, "htest")
	client.Get(url).Send().HeaderXForwardedProtocol("htest")
}

func TestResponse_HeaderXForwardedSsl(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXForwardedSsl, "htest")
	client.Get(url).Send().HeaderXForwardedSsl("htest")
}

func TestResponse_HeaderXUrlScheme(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXUrlScheme, "htest")
	client.Get(url).Send().HeaderXUrlScheme("htest")
}

func TestResponse_HeaderXHTTPMethodOverride(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXHTTPMethodOverride, "htest")
	client.Get(url).Send().HeaderXHTTPMethodOverride("htest")
}

func TestResponse_HeaderXRealIP(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXRealIP, "htest")
	client.Get(url).Send().HeaderXRealIP("htest")
}

func TestResponse_HeaderXRequestID(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXRequestID, "htest")
	client.Get(url).Send().HeaderXRequestID("htest")
}

func TestResponse_HeaderServer(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderServer, "htest")
	client.Get(url).Send().HeaderServer("htest")
}

func TestResponse_HeaderOrigin(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderOrigin, "htest")
	client.Get(url).Send().HeaderOrigin("htest")
}

func TestResponse_HeaderAccessControlRequestMethod(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlRequestMethod, "htest")
	client.Get(url).Send().HeaderAccessControlRequestMethod("htest")
}

func TestResponse_HeaderAccessControlRequestHeaders(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlRequestHeaders, "htest")
	client.Get(url).Send().HeaderAccessControlRequestHeaders("htest")
}

func TestResponse_HeaderAccessControlAllowOrigin(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowOrigin, "htest")
	client.Get(url).Send().HeaderAccessControlAllowOrigin("htest")
}

func TestResponse_HeaderAccessControlAllowMethods(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowMethods, "htest")
	client.Get(url).Send().HeaderAccessControlAllowMethods("htest")
}

func TestResponse_HeaderAccessControlAllowHeaders(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowHeaders, "htest")
	client.Get(url).Send().HeaderAccessControlAllowHeaders("htest")
}

func TestResponse_HeaderAccessControlAllowCredentials(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlAllowCredentials, "htest")
	client.Get(url).Send().HeaderAccessControlAllowCredentials("htest")
}

func TestResponse_HeaderAccessControlExposeHeaders(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlExposeHeaders, "htest")
	client.Get(url).Send().HeaderAccessControlExposeHeaders("htest")
}

func TestResponse_HeaderAccessControlMaxAge(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccessControlMaxAge, "htest")
	client.Get(url).Send().HeaderAccessControlMaxAge("htest")
}

func TestResponse_HeaderStrictTransportSecurity(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderStrictTransportSecurity, "htest")
	client.Get(url).Send().HeaderStrictTransportSecurity("htest")
}

func TestResponse_HeaderXContentTypeOptions(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXContentTypeOptions, "htest")
	client.Get(url).Send().HeaderXContentTypeOptions("htest")
}

func TestResponse_HeaderXXSSProtection(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXXSSProtection, "htest")
	client.Get(url).Send().HeaderXXSSProtection("htest")
}

func TestResponse_HeaderXFrameOptions(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXFrameOptions, "htest")
	client.Get(url).Send().HeaderXFrameOptions("htest")
}

func TestResponse_HeaderContentSecurityPolicy(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentSecurityPolicy, "htest")
	client.Get(url).Send().HeaderContentSecurityPolicy("htest")
}

func TestResponse_HeaderXCSRFToken(t *testing.T) {
	client := NewClient(t).To(ResponseHeadersServer)
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderXCSRFToken, "htest")
	client.Get(url).Send().HeaderXCSRFToken("htest")
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
