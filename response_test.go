package htest

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	client := NewClient(t).To(Mux)
	assert.Equal(t, UserData, client.Get("/body/user").Send().StatusOK().String())
}

func TestResponse_Bytes(t *testing.T) {
	client := NewClient(t).To(Mux)
	assert.Equal(t, []byte(UserData), client.Get("/body/user").Send().StatusOK().Bytes())
}

func TestResponse_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Send().StatusOK().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}
