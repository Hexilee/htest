package example

import (
	"testing"
	"github.com/Hexilee/htest"
)

func TestNameHandlerFunc(t *testing.T) {
	client := htest.NewClient(t).ToFunc(NameHandler)
	body := client.Get("").Send().StatusOK().JSON()
	body.String("name", "hexi")
}

func TestNameHandler(t *testing.T) {
	client := htest.NewClient(t).To(Mux)
	body := client.Get("/name").Send().StatusOK().JSON()
	body.String("name", "hexi")
}

func TestNameHandlerEcho(t *testing.T) {
	client := htest.NewClient(t).To(server)
	client.Get("/name").Send().StatusOK().JSON().String("name", "hexi")
	client.Get("/stuid").Send().StatusOK().JSON().String("stuid", "3160100001")
}