## htest is a http-test package

[![Coverage Status](https://coveralls.io/repos/github/Hexilee/htest/badge.svg)](https://coveralls.io/github/Hexilee/htest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hexilee/htest)](https://goreportcard.com/report/github.com/Hexilee/htest)
[![Build Status](https://travis-ci.org/Hexilee/htest.svg?branch=master)](https://travis-ci.org/Hexilee/htest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hexilee/htest/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/Hexilee/htest?status.svg)](https://godoc.org/github.com/Hexilee/htest)

## Contents

- [Basic Usage](#BasicUsage)
    - [As MockClient](#AsMockClient)
        - [Test HandlerFunc](#TestHandlerFunc)
        - [To ServeMux](#ToServeMux)
        - [To Echo](#ToEcho)
    - [As RealClient](#AsRealClient)
        - [Github API](#GithubAPI)

<h3 id="BasicUsage">Basic Usage</h3>

-----------------

<h4 id="AsMockClient">As MockClient</h4>


<h5 id="TestHandlerFunc">Test HandlerFunc</h5>

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

<h5 id="ToServeMux">To ServeMux</h5>

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

<h5 id="ToEcho">To Echo</h5>

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

<h4 id="AsRealClient">As RealClient</h4>

<h5 id="GithubAPI">Github API</h5>

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

