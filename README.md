## htest is a http-test package

[![Coverage Status](https://coveralls.io/repos/github/Hexilee/htest/badge.svg)](https://coveralls.io/github/Hexilee/htest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hexilee/htest)](https://goreportcard.com/report/github.com/Hexilee/htest)
[![Build Status](https://travis-ci.org/Hexilee/htest.svg?branch=master)](https://travis-ci.org/Hexilee/htest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hexilee/htest/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/Hexilee/htest?status.svg)](https://godoc.org/github.com/Hexilee/htest)


Table of Contents
=================

* [Basic Usage](#basic-usage)
  * [Test MockServer](#test-mockserver)
     * [Test HandlerFunc](#test-handlerfunc)
     * [To ServeMux](#to-servemux)
     * [To Echo](#to-echo)
  * [Test RealServer](#test-realserver)
     * [Github API](#github-api)
* [Client](#client)
  * [Set MockServer](#set-mockserver)
     * [HandlerFunc](#handlerfunc)
     * [Handler](#handler)
  * [Construct Request](#construct-request)
     * [Http Methods](#http-methods)
* [Request](#request)
  * [Set Headers](#set-headers)
  * [Add Cookie](#add-cookie)
  * [Test](#test)
  * [Send](#send)
  * [As http.Request](#as-httprequest)
* [Response](#response)
  * [Assert StatusCode](#assert-statuscode)
     * [Code](#code)
     * [StatusXXX](#statusxxx)
  * [Assert Headers](#assert-headers)
     * [Headers](#headers)
     * [HeaderXXX](#headerxxx)
  * [Assert Body](#assert-body)
  * [Get Body](#get-body)
  * [Bind Body](#bind-body)
  * [Body Types](#body-types)
  * [As http.Response](#as-httpresponse)
* [Body](#body)
  * [JSON](#json)
     * [Assert JSON Key](#assert-json-key)
     * [Assert JSON Empty or Not](#assert-json-empty-or-not)
     * [Bind JSON](#bind-json)
  * [XML](#xml)
  * [MD5](#md5)
     * [Assert MD5 Hash](#assert-md5-hash)
     * [Get MD5 Hash value](#get-md5-hash-value)
  * [SHA1](#sha1)
* [Appendix](#appendix)
  * [consts](#consts)


### Basic Usage

-----------------

#### Test MockServer

> Test a Handler or a HandlerFunc

##### Test HandlerFunc

```go
// example/basic_mock_client.go
package myapp

import (
	"io"
	"net/http"
)

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}
```

```go
// example/basic_mock_client_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerFunc(t *testing.T) {
	htest.NewClient(t).
		ToFunc(NameHandler).
		Get("").
		Test().
		StatusOK().
		JSON().
		String("name", "hexi")
}
```

You can also test handler (*http.ServeMux, *echo.Echo .etc.)

##### To ServeMux

```go
// example/basic_mock_client.go
package myapp

import (
	"io"
	"net/http"
)

var (
	Mux *http.ServeMux
)

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/name", NameHandler)
}

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}
```

```go
// example/basic_mock_client_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandler(t *testing.T) {
	htest.NewClient(t).
		To(Mux).
		Get("/name").
		Test().
		StatusOK().
		JSON().
		String("name", "hexi")
}
```

##### To Echo

```go
// example/basic_mock_client.go
package myapp

import (
	"io"
	"github.com/labstack/echo"
)

var (
	server *echo.Echo
)

func init() {
	server = echo.New()
	server.GET("/name", NameHandlerEcho)
}

func NameHandlerEcho(c echo.Context) error {
	return c.String(http.StatusOK, `{"name": "hexi"}`)
}
```

```go
// example/basic_mock_client_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerEcho(t *testing.T) {
	htest.NewClient(t).
		To(server).
		Get("/name").
		Test().
		StatusOK().
		JSON().
		String("name", "hexi")
}
```

#### Test RealServer

> Send a http request and test the response

##### Github API

```go
// request_test.go
func TestRequest_Send(t *testing.T) {
	NewClient(t).
		Get("https://api.github.com/users/Hexilee").
		Send().
		StatusOK().
		JSON().
		String("login", "Hexilee")
}
```


### Client

-------

#### Set MockServer

> Set mock server to be tested (Do not need it when you test real server)

##### HandlerFunc

> Set a HandlerFunc as mock server

```go
// example/basic_mock_client_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerFunc(t *testing.T) {
	htest.NewClient(t).
		ToFunc(NameHandler).
		Get("").
		Test().
		StatusOK().
		JSON().
		String("name", "hexi")
}
```

##### Handler

> Set a Handler as mock server

```go
// example/basic_mock_client_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandler(t *testing.T) {
	htest.NewClient(t).
		To(Mux).
		Get("/name").
		Test().
		StatusOK().
		JSON().
		String("name", "hexi")
}
```

#### Construct Request

> Construct htest.Request using different http methods

##### Http Methods

> For example

- Get

```go
// client.go
func (c Client) Get(path string) *Request
```

> More

- Head
- Trace
- Options
- Connect
- Delete
- Post
- Put
- Patch

### Request

-------

#### Set Headers

> Set headers and return *Request for chaining-call

- SetHeader

```go
// server_test.go

Mux.Get("/request/header", HeaderHandler)

// request_test.go

func HeaderHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(HeaderContentType) == MIMEApplicationJSON {
		io.WriteString(w, `{"result": "JSON"}`)
		return
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func TestRequest_SetHeader(t *testing.T) {
	client := NewClient(t).To(Mux)
	// bad content type
	client.
		Get("/request/header").
		SetHeader(HeaderContentType, MIMEApplicationForm).
		Test().
		StatusBadRequest()

	// right
	client.
		Get("/request/header").
		SetHeader(HeaderContentType, MIMEApplicationJSON).
		Test().
		StatusOK().
		JSON().
		String("result", "JSON")
}
```

> HeaderContentType, MIMEApplicationForm are constants in const.go
> For more information, you can refer to [Appendix](#appendix)
 

- SetHeaders

```go
// request_test.go

func TestRequest_SetHeaders(t *testing.T) {
	client := NewClient(t).To(Mux)
	// bad content type
	client.Get("/request/header").
		SetHeaders(
			map[string]string{
				HeaderContentType: MIMEApplicationForm,
			},
		).
		Test().
		StatusBadRequest()

	// right
	client.Get("/request/header").
		SetHeaders(
			map[string]string{
				HeaderContentType: MIMEApplicationJSON,
			},
		).
		Test().
		StatusOK().
		JSON().
		String("result", "JSON")
}
```

#### Add Cookie

> Add cookie and return *Request for chaining-call

```go
// server_test.go

Mux.Get("/request/cookie", CookieHandler)

// request_test.go

var (
	testCookie = http.Cookie{Name: "test_cookie", Value: "cookie_value"}
)

func CookieHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(testCookie.Name)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	io.WriteString(w, fmt.Sprintf(`{"cookie": "%s"}`, cookie))
}


func TestRequest_AddCookie(t *testing.T) {
	client := NewClient(t).
		To(Mux)
	client.
		Get("/request/cookie").
		Test().
		StatusForbidden()
	client.
		Get("/request/cookie").
		AddCookie(&testCookie).
		Test().
		StatusOK().
		JSON().
		String("cookie", testCookie.String())
}
```

#### Test

> Calling *Request.Test will test the mock server and return a *Response. 

> You must have called Client.To or Client.ToFunc, otherwise causing a panic (htest.MockNilError)

```go
// request_test.go

func TestRequest_Test(t *testing.T) {
	defer func() {
		assert.Equal(t, MockNilError, recover())
	}()

	NewClient(t).
		Get("/request/header").
		SetHeader(HeaderContentType, MIMEApplicationForm).
		Test().
		StatusBadRequest()
}
```

#### Send

> Calling *Request.Send will send a real http request and return a *Response

```go
// request_test.go

func TestRequest_Send(t *testing.T) {
	NewClient(t).
		Get("https://api.github.com/users/Hexilee").
		Send().
		StatusOK().
		JSON().
		String("login", "Hexilee")
}

```

#### As http.Request

> As *http.Request is embedded in htest.Request, you can regard *htest.Request as *http.Request. Just like: 

```go
userAgent := NewClient(t).
		Get("https://api.github.com/users/Hexilee").
		UserAgent()
    
```


### Response

-------

#### Assert StatusCode

Assert Response.StatusCode

##### Code

> *Response.Code(statusCode int)

```go
// response_test.go

var (
	ResponseCodeServer    = chi.NewRouter()
)

func init() {
	ResponseCodeServer.Get("/response/statusCode/{code}", StatusHandler)
}

func StatusHandler(w http.ResponseWriter, req *http.Request) {
	codeStr := chi.URLParam(req, "code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(code)
}	

func TestResponse_Code(t *testing.T) {
	NewClient(t).
		To(ResponseCodeServer).
		Get(fmt.Sprintf("/response/statusCode/%d", http.StatusBadRequest)).
		Test().
		Code(http.StatusBadRequest)
}
```


##### StatusXXX

> For more ergonomic development, *htest.Response has many methods to assert all the StatusCode in net/http

```go
// response_test.go

func TestResponse_StatusContinue(t *testing.T) {
	NewClient(t).
		To(ResponseCodeServer).
		Get(fmt.Sprintf("/response/statusCode/%d", http.StatusContinue)).
		Test().
		StatusContinue()
}

```

#### Assert Headers

Assert Response.Headlers

##### Headers

> *Response.Headers(key, expect string)

```go
// response_test.go

var (
	ResponseHeadersServer = chi.NewRouter()
)

func init() {
	ResponseHeadersServer.Get("/response/headers", HeadersHandler)
}

func HeadersHandler(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	header := query.Get("header")
	value := query.Get("value")
	w.Header().Set(header, value)
}

func TestResponse_Headers(t *testing.T) {
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderContentType, MIMEApplicationJSON)
	NewClient(t).
		To(ResponseHeadersServer).
		Get(url).
		Test().
		Headers(HeaderContentType, MIMEApplicationJSON)
}
```

##### HeaderXXX

> For more ergonomic development, *htest.Response has many methods to assert all the Headers in const.go


```go
// response_test.go

func TestResponse_HeaderAccept(t *testing.T) {
	url := fmt.Sprintf("/response/headers?header=%s&value=%s", HeaderAccept, "htest")
	NewClient(t).
		To(ResponseHeadersServer).
		Get(url).
		Test().
		HeaderAccept("htest")
}
```

#### Assert Body

You can assert data in body straightly.

```go
// server_test.go
Mux.Get("/body/user", UserDataHandler)

func UserDataHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserData)
}

// response_test.go

const (
	UserData = `{
	"id": 1,
	"name": "hexi"
}`
)



func TestResponse_Expect(t *testing.T) {
	NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		Expect(UserData)
}
```


#### Get Body

You can get data in body straightly

- String

```go
// response_test.go

func TestResponse_String(t *testing.T) {
	assert.Equal(t, UserData, NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		String())
}
```

- Bytes

```go
// response_test.go

func TestResponse_Bytes(t *testing.T) {
	assert.Equal(t, []byte(UserData), NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		Bytes())
}

```
#### Bind Body

If type of data in body is JSON, you can unmarshal it straightly

```go
// response_test.go

type (
	User struct {
		Id   uint   
		Name string
	}
)

func TestResponse_Bind(t *testing.T) {
	user := new(User)
	NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}
```

#### Body Types

You can return data in 4 types

* [JSON](#json)
* [XML](#xml)
* [MD5](#md5)
* [SHA1](#sha1)


#### As http.Response

> As *http.Response is embedded in htest.Response, you can regard *htest.Response as *http.Response. Just like: 

```go

assert.Equal(t, "HTTP/1.1", NewClient(t).
                        		To(Mux).
                        		Get("/body/user").
                        		Test().
                        		Proto
)
    
```

### Body

htest provide 4 types of data to be returned

#### JSON

data as JSON

##### Assert JSON Key

- Exist(key string)
- NotExist(key string)
- String(key, expect string)
- Int(key string, expect int64)
- True(key string)
- False(key string)
- Uint(key string, expect uint64)
- Time(key string, expect time.Time)
- Float(key string, expect float64)

```go
// body_test.go

func TestJSON_Exist(t *testing.T) {
	NewClient(t).
		To(Mux).
		Get("/name").
		Test().
		StatusOK().
		JSON().
		Exist("name").
		NotExist("stuid")
}

```

```go
func TestJSON_String(t *testing.T) {
	user := new(User)
	NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		JSON().
		String("name", "hexi)
}
```

##### Assert JSON Empty or Not

```go
func TestJSON_NotEmpty(t *testing.T) {
	user := new(User)
	NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		JSON().
		NotEmpty()
}
```

##### Bind JSON

```go
// body_test.go

type (
	User struct {
		Id   uint   
		Name string
	}
)

func TestJSON_Bind(t *testing.T) {
	user := new(User)
	NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		JSON().
		Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}
```

#### XML

Same as JSON.

For more examples, you can find them in body_test.go

#### MD5

##### Assert MD5 Hash

```go
// body_test.go

func TestMD5_Expect(t *testing.T) {
	NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		MD5().
		Expect(UserDataMD5)
}
```

##### Get MD5 Hash value

```go
hash := NewClient(t).
		To(Mux).
		Get("/body/user").
		Test().
		StatusOK().
		MD5().
		Body()
```

#### SHA1

Same as MD5.

For more examples, you can find them in body_test.go


### Appendix

#### consts

There are many constants of header or header value in const.go

```go
// const.go

package htest

// HTTP methods
const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	GET     = "GET"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	PATCH   = "PATCH"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

// MIME types
const (
	MIMEApplicationJSON                  = "application/json"
	MIMEApplicationJSONCharsetUTF8       = MIMEApplicationJSON + "; " + charsetUTF8
	MIMEApplicationJavaScript            = "application/javascript"
	MIMEApplicationJavaScriptCharsetUTF8 = MIMEApplicationJavaScript + "; " + charsetUTF8
	MIMEApplicationXML                   = "application/xml"
	MIMEApplicationXMLCharsetUTF8        = MIMEApplicationXML + "; " + charsetUTF8
	MIMETextXML                          = "text/xml"
	MIMETextXMLCharsetUTF8               = MIMETextXML + "; " + charsetUTF8
	MIMEApplicationForm                  = "application/x-www-form-urlencoded"
	MIMEApplicationProtobuf              = "application/protobuf"
	MIMEApplicationMsgpack               = "application/msgpack"
	MIMETextHTML                         = "text/html"
	MIMETextHTMLCharsetUTF8              = MIMETextHTML + "; " + charsetUTF8
	MIMETextPlain                        = "text/plain"
	MIMETextPlainCharsetUTF8             = MIMETextPlain + "; " + charsetUTF8
	MIMEMultipartForm                    = "multipart/form-data"
	MIMEOctetStream                      = "application/octet-stream"
)

const (
	charsetUTF8 = "charset=UTF-8"
)

// Headers
const (
	HeaderAccept              = "Accept"
	HeaderAcceptEncoding      = "Accept-Encoding"
	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Proto"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-IP"
	HeaderXRequestID          = "X-Request-ID"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"

	// Access control
	HeaderAccessControlRequestMethod    = "Access-Control-Request-Method"
	HeaderAccessControlRequestHeaders   = "Access-Control-Request-Headers"
	HeaderAccessControlAllowOrigin      = "Access-Control-Allow-Origin"
	HeaderAccessControlAllowMethods     = "Access-Control-Allow-Methods"
	HeaderAccessControlAllowHeaders     = "Access-Control-Allow-Headers"
	HeaderAccessControlAllowCredentials = "Access-Control-Allow-Credentials"
	HeaderAccessControlExposeHeaders    = "Access-Control-Expose-Headers"
	HeaderAccessControlMaxAge           = "Access-Control-Max-Age"

	// Security
	HeaderStrictTransportSecurity = "Strict-Transport-Security"
	HeaderXContentTypeOptions     = "X-Content-Type-Options"
	HeaderXXSSProtection          = "X-XSS-Protection"
	HeaderXFrameOptions           = "X-Frame-Options"
	HeaderContentSecurityPolicy   = "Content-Security-Policy"
	HeaderXCSRFToken              = "X-CSRF-Token"
)


```




