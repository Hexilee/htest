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
     * [Body Types](#body-types)
  * [As http.Response](#as-httpresponse)
* [Body](#body)
  * [JSON](#json)
  * [XML](#xml)
  * [MD5](#md5)
  * [SHA1](#sha1)
* [Appendix](#appendix)


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

> Calling *Request.Test will test the mock server
> You must have called Client.To or Client.ToFunc, otherwise causing a panic (htest.MockNilError)

#### Send
#### As http.Request

### Response

-------

#### Assert StatusCode
##### Code
##### StatusXXX
#### Assert Headers
##### Headers
##### HeaderXXX
#### Assert Body
#### Get Body
##### Body Types
#### As http.Response
### Body
#### JSON
#### XML
#### MD5
#### SHA1

### Appendix



