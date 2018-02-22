package htest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJSON_Exist(t *testing.T) {
	client := NewClient().To(Mux)
	client.Get("/name").Send().With(t).OK().JSON().Exist("name").NotExist("stuid")
}

func TestJSON_Bind(t *testing.T) {
	user := new(User)
	client := NewClient().To(Mux)
	client.Get("/body/user").Send().With(t).OK().JSON().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}
