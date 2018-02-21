package htest

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
	Mux.HandleFunc("/header", headerHandler)
}

func NameHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, `{"name": "hexi"}`)
}

func headerHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(HeaderContentType) == MIMEApplicationJSON {
		io.WriteString(w, `{"result": "JSON"}`)
		return
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
