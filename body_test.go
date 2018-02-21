package htest

import "testing"

func TestJSON_Exist(t *testing.T) {
	client := NewClient().To(Mux)
	client.Get("/name").Send().OK().JSON().Exist("name").NotExist("stuid")
}

func TestJSON_Bind(t *testing.T) {
	user := new(User)
	client := NewClient().To(Mux)
	client.Get("/body/user").Send().OK().JSON().Bind(user)
}