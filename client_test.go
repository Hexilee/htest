package htest

import (
	"testing"
)

func TestClient_ToFunc(t *testing.T) {
	client := NewClient().ToFunc(NameHandler)
	body := client.Get("").Send().OK().JSON()
	body.String("name", "hexi")
}

func TestClient_Get(t *testing.T) {
	client := NewClient().To(Mux)
	body := client.Get("/name").Send().OK().JSON()
	body.String("name", "hexi")
}
