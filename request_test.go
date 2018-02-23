package htest

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

var (
	testCookie = http.Cookie{Name: "test_cookie", Value: "cookie_value"}
)

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

func TestRequest_Send(t *testing.T) {
	NewClient(t).
		Get("https://api.github.com/users/Hexilee").
		Send().
		StatusOK().
		JSON().
		String("login", "Hexilee")
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

func HeaderHandler(w http.ResponseWriter, req *http.Request) {
	if req.Header.Get(HeaderContentType) == MIMEApplicationJSON {
		io.WriteString(w, `{"result": "JSON"}`)
		return
	}
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func CookieHandler(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie(testCookie.Name)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		return
	}
	io.WriteString(w, fmt.Sprintf(`{"cookie": "%s"}`, cookie))
}
