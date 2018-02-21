package example

import (
	"io"
	"net/http"
	"github.com/labstack/echo"
)

var (
	Mux    *http.ServeMux
	server *echo.Echo
)

func init() {
	Mux = http.NewServeMux()
	Mux.HandleFunc("/name", NameHandler)
	server = echo.New()
	server.GET("/name", NameHandlerEcho)
	server.GET("/stuid", StuidHandlerEcho)
}

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}

func NameHandlerEcho(c echo.Context) error {
	return c.String(http.StatusOK, `{"name": "hexi"}`)
}

func StuidHandlerEcho(c echo.Context) error {
	return c.String(http.StatusOK, `{"stuid": "3160100001"}`)
}
