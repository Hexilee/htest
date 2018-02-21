package htest

import "testing"

type (
	User struct {
		Id uint
		Name string
	}
)

const (
	UserData = `{
	"id": 1,
	"name": "hexi"
}`
)

func TestJSON_Bind(t *testing.T) {

}