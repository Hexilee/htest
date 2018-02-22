package htest

import (
	"io"
	"net/http"
	"github.com/go-chi/chi"
)

var (
	Mux *chi.Mux
)

func init() {
	Mux = chi.NewRouter()
	Mux.Get("/name", NameHandler)
	Mux.Head("/name", NameHandler)
	Mux.Get("/request/header", HeaderHandler)
	Mux.Get("/body/user", UserDataHandler)
}

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}
