package htest

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestClient_ToFunc(t *testing.T) {
	client := NewClient().ToFunc(NameHandler)
	client.Get("").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_To(t *testing.T) {
	// if Client immutable
	client := NewClient()
	client.To(Mux)
	assert.Nil(t, client.handler)
}

func TestClient_Get(t *testing.T) {
	client := NewClient().To(Mux)
	client.Get("/client/get").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Trace(t *testing.T) {
	client := NewClient().To(Mux)
	client.Trace("/client/trace").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Connect(t *testing.T) {
	client := NewClient().To(Mux)
	client.Connect("/client/connect").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Delete(t *testing.T) {
	client := NewClient().To(Mux)
	client.Delete("/client/delete").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Options(t *testing.T) {
	client := NewClient().To(Mux)
	client.Options("/client/options").Send().With(t).OK().JSON().String("name", "hexi")
}

func TestClient_Head(t *testing.T) {
	client := NewClient().To(Mux)
	client.Head("/client/head").Send().With(t).OK()
}

func TestClient_Post(t *testing.T) {
	user := &User{Id: 0}
	data, _ := json.Marshal(user)
	dataReader := bytes.NewBuffer(data)
	client := NewClient().To(Mux)
	client.Post("/client/post", dataReader).Send().With(t).OK().Bind(user)

	assert.Equal(t, uint(1), user.Id)
	assert.Equal(t, "hexi", user.Name)
}

func TestClient_Put(t *testing.T) {
	user := &User{Id: 0}
	data, _ := json.Marshal(user)
	dataReader := bytes.NewBuffer(data)
	client := NewClient().To(Mux)
	client.Put("/client/put", dataReader).Send().With(t).OK().Bind(user)

	assert.Equal(t, uint(1), user.Id)
	assert.Equal(t, "hexi", user.Name)
}

func TestClient_Patch(t *testing.T) {
	user := &User{Id: 0}
	data, _ := json.Marshal(user)
	dataReader := bytes.NewBuffer(data)
	client := NewClient().To(Mux)
	client.Patch("/client/patch", dataReader).Send().With(t).OK().Bind(user)

	assert.Equal(t, uint(1), user.Id)
	assert.Equal(t, "hexi", user.Name)
}

func ClientDataHandler(w http.ResponseWriter, req *http.Request) {
	user := new(User)
	resp, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(resp, user)
	if user.Id == 0 {
		io.WriteString(w, UserData)
	}
}
