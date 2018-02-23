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
	NewClient(t).
		ToFunc(NameHandler).
		Get("").
		Test().
		StatusOK().JSON().String("name", "hexi")
}

func TestClient_To(t *testing.T) {
	// if Client immutable
	client := NewClient(t)
	client.To(Mux)
	assert.Nil(t, client.handler)
}

func TestClient_Get(t *testing.T) {
	NewClient(t).
		To(Mux).
		Get("/client/get").
		Test().
		StatusOK().JSON().String("name", "hexi")
}

func TestClient_Trace(t *testing.T) {
	NewClient(t).
		To(Mux).
		Trace("/client/trace").
		Test().
		StatusOK().JSON().String("name", "hexi")
}

func TestClient_Connect(t *testing.T) {
	NewClient(t).
		To(Mux).
		Connect("/client/connect").
		Test().
		StatusOK().JSON().String("name", "hexi")
}

func TestClient_Delete(t *testing.T) {
	NewClient(t).
		To(Mux).
		Delete("/client/delete").
		Test().
		StatusOK().JSON().String("name", "hexi")
}

func TestClient_Options(t *testing.T) {
	NewClient(t).
		To(Mux).
		Options("/client/options").
		Test().
		StatusOK().JSON().String("name", "hexi")
}

func TestClient_Head(t *testing.T) {
	NewClient(t).
		To(Mux).
		Head("/client/head").
		Test().
		StatusOK()
}

func TestClient_Post(t *testing.T) {
	user := &User{Id: 0}
	data, _ := json.Marshal(user)
	dataReader := bytes.NewBuffer(data)
	NewClient(t).
		To(Mux).
		Post("/client/post", dataReader).
		Test().
		StatusOK().
		Bind(user)

	assert.Equal(t, uint(1), user.Id)
	assert.Equal(t, "hexi", user.Name)
}

func TestClient_Put(t *testing.T) {
	user := &User{Id: 0}
	data, _ := json.Marshal(user)
	dataReader := bytes.NewBuffer(data)
	NewClient(t).
		To(Mux).
		Put("/client/put", dataReader).
		Test().
		StatusOK().
		Bind(user)

	assert.Equal(t, uint(1), user.Id)
	assert.Equal(t, "hexi", user.Name)
}

func TestClient_Patch(t *testing.T) {
	user := &User{Id: 0}
	data, _ := json.Marshal(user)
	dataReader := bytes.NewBuffer(data)
	NewClient(t).
		To(Mux).
		Patch("/client/patch", dataReader).
		Test().
		StatusOK().
		Bind(user)

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
