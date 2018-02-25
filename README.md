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
##### HandlerFunc
##### Handler
#### Construct Request
##### Http Methods

### Request

-------

#### Set Headers
#### Add Cookie
#### Test
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


