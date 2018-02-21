package htest

import (
	"testing"
)

func TestClient_ToFunc(t *testing.T) {
	client := NewClient().ToFunc(NameHandler)
	client.Get("").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Get(t *testing.T) {
	client := NewClient().To(Mux)
	client.Get("/name").Send().With(t).OK().JSON().String("name", "hexi")
}