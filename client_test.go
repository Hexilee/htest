package htest

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClient_ToFunc(t *testing.T) {
	client := NewClient().ToFunc(NameHandler)
	client.Get("").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Get(t *testing.T) {
	client := NewClient().To(Mux)
	client.Get("/name").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_To(t *testing.T) {
	// if Client immutable
	client := NewClient()
	client.To(Mux)
	assert.Nil(t, client.handler)
}
