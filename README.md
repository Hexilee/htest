## htest is a http-test package

[![Coverage Status](https://coveralls.io/repos/github/Hexilee/htest/badge.svg)](https://coveralls.io/github/Hexilee/htest)
[![Go Report Card](https://goreportcard.com/badge/github.com/Hexilee/htest)](https://goreportcard.com/report/github.com/Hexilee/htest)
[![Build Status](https://travis-ci.org/Hexilee/htest.svg?branch=master)](https://travis-ci.org/Hexilee/htest)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Hexilee/htest/blob/master/LICENSE)
[![Documentation](https://godoc.org/github.com/Hexilee/htest?status.svg)](https://godoc.org/github.com/Hexilee/htest)

### Usage

#### Test HandlerFunc

```go
// handler.go
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
// handler_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerFunc(t *testing.T) {
	client := htest.NewClient(t).ToFunc(NameHandler)
	body := client.Get("").Test().StatusOK().JSON()
	body.String("name", "hexi")
}
```

You can also test handler (*http.ServeMux, *echo.Echo .etc.)

#### To ServeMux

```go
// handler.go
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
// handler_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandler(t *testing.T) {
	client := htest.NewClient(t).To(Mux)
	body := client.Get("/name").Test().StatusOK().JSON()
	body.String("name", "hexi")
}
```

#### To Echo

```go
// handler.go
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
// handler_test.go
package myapp

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerEcho(t *testing.T) {
	client := htest.NewClient(t).To(server)
	client.Get("/name").Test().StatusOK().JSON().String("name", "hexi")
}
```