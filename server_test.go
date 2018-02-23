package htest

import (
	"github.com/go-chi/chi"
	"io"
	"net/http"
)

var (
	Mux *chi.Mux
)

func init() {
	Mux = chi.NewRouter()
	Mux.Get("/name", NameHandler)
	Mux.Get("/client/get", NameHandler)
	Mux.Trace("/client/trace", NameHandler)
	Mux.Delete("/client/delete", NameHandler)
	Mux.Connect("/client/connect", NameHandler)
	Mux.Options("/client/options", NameHandler)
	Mux.Head("/client/head", NameHandler)
	Mux.Post("/client/post", ClientDataHandler)
	Mux.Put("/client/put", ClientDataHandler)
	Mux.Patch("/client/patch", ClientDataHandler)
	Mux.Post("/client/patch", ClientDataHandler)
	Mux.Get("/request/header", HeaderHandler)
	Mux.Get("/request/cookie", CookieHandler)
	Mux.Get("/body/user", UserDataHandler)
	Mux.Get("/xml_body/user", UserDataXMLHandler)
}

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}
