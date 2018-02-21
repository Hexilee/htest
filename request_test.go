package htest

import "testing"

func TestRequest_SetHeader(t *testing.T) {
	client := NewClient().To(Mux)
	// bad content type
	client.Get("/request/header").SetHeader(HeaderContentType, MIMEApplicationForm).Send().BadRequest()

	// right
	body := client.Get("/request/header").SetHeader(HeaderContentType, MIMEApplicationJSON).Send().OK().JSON()
	body.String("result", "JSON")
}

func TestRequest_SetHeaders(t *testing.T) {
	client := NewClient().To(Mux)
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
