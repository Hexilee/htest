package htest

import (
	"io"
	"net/http"
	"testing"
)

func TestRequest_SetHeader(t *testing.T) {
	client := NewClient(t).To(Mux)
	// bad content type
	client.Get("/request/header").SetHeader(HeaderContentType, MIMEApplicationForm).Send().BadRequest()

	// right
	body := client.Get("/request/header").SetHeader(HeaderContentType, MIMEApplicationJSON).Send().OK().JSON()
	body.String("result", "JSON")
}

func TestRequest_SetHeaders(t *testing.T) {
	client := NewClient(t).To(Mux)
	// bad content type
	client.Get("/request/header").SetHeaders(
		map[string]string{
			HeaderContentType: MIMEApplicationForm,
		},
	).Send().BadRequest()

	// right
	body := client.Get("/request/header").SetHeaders(
		map[string]string{
			HeaderContentType: MIMEApplicationJSON,
		},
	).Send().OK().JSON()
	body.String("result", "JSON")
}

func HeaderHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(HeaderContentType) == MIMEApplicationJSON {
		io.WriteString(w, `{"result": "JSON"}`)
		return
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
