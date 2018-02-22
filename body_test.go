package htest

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

const (
	WrongXMLData = `
<?xml version="1.0" encoding="UTF-8"?>
<user>
	<id>1</id>
	<name>hexi</name>
`
)

func TestJSON_Exist(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/name").Send().StatusOK().JSON().Exist("name").NotExist("stuid")
}

func TestJSON_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Send().StatusOK().JSON().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func TestXML_Exist(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/xml_body/user").Send().StatusOK().XML().Exist("user.name").NotExist("user.stuid")
}

func TestWrongXML_Exist(t *testing.T) {
	NewXML([]byte(WrongXMLData), t).Empty()
}

func TestXML_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/xml_body/user").Send().StatusOK().XML().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func UserDataHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserData)
}

func UserDataXMLHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserDataXML)
}
