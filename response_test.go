package htest

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

type (
	User struct {
		Id   uint
		Name string
	}
)

const (
	UserData = `{
	"id": 1,
	"name": "hexi"
}`
)

func TestResponse_String(t *testing.T) {
	client := NewClient().To(Mux)
	assert.Equal(t, UserData, client.Get("/body/user").Send().With(t).OK().String())
}

func TestResponse_Bytes(t *testing.T) {
	client := NewClient().To(Mux)
	assert.Equal(t, []byte(UserData), client.Get("/body/user").Send().With(t).OK().Bytes())
}

func TestResponse_Bind(t *testing.T) {
	user := new(User)
	client := NewClient().To(Mux)
	client.Get("/body/user").Send().With(t).OK().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}