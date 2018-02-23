package htest

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
	"time"
)

const (
	WrongXMLData = `
<?xml version="1.0" encoding="UTF-8"?>
<user>
	<id>1</id>
	<name>hexi</name>
`

	JSONAssertData = `
{
	"number": 1,
	"time": "2018-02-22T00:00:00Z",
	"ok": true,
	"no": false,
}
`
	XMLAssertData = `
<data>
	<Number>1</Number>
	<Time>2018-02-22T00:00:00Z</Time>
	<OK>true</OK>
	<NO>false</NO>
</data>
`

	AssertDataTimeStr = "2018-02-22T00:00:00Z"
)

var (
	AssertDataTime, _ = time.Parse(time.RFC3339, AssertDataTimeStr)
)

type (
	AssertStruct struct {
		Number int       `json:"number" xml:"Number"`
		Time   time.Time `json:"time" xml:"Time"`
		OK     bool      `json:"ok" xml:"OK"`
		NO     bool      `json:"no" xml:"NO"`
	}
)

func TestJSON_Exist(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/name").Test().StatusOK().JSON().Exist("name").NotExist("stuid")
}

func TestJSON_String(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).String("time", AssertDataTimeStr)
}

func TestJSON_Int(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Int("number", int64(1))
}

func TestJSON_True(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).True("ok")
}

func TestJSON_False(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).False("no")
}

func TestJSON_Uint(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Uint("number", uint64(1))
}

func TestJSON_Time(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Time("time", AssertDataTime)
}

func TestJSON_Float(t *testing.T) {
	NewJSON([]byte(JSONAssertData), t).Float("number", float64(1))
}

func TestJSON_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Test().StatusOK().JSON().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func TestJSON_NotEmpty(t *testing.T) {
	NewXML([]byte(UserDataXML), t).JSON.NotEmpty()
}

func TestXML_Exist(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/xml_body/user").Test().StatusOK().XML().Exist("user.name").NotExist("user.stuid")
}

func TestXML_String(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).String("data.Time", AssertDataTimeStr)
}

func TestXML_Int(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).Int("data.Number", int64(1))
}

func TestXML_True(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).True("data.OK")
}

func TestXML_False(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).False("data.NO")
}

func TestXML_Uint(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).Uint("data.Number", uint64(1))
}

func TestXML_Time(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).Time("data.Time", AssertDataTime)
}

func TestXML_Float(t *testing.T) {
	NewXML([]byte(XMLAssertData), t).Float("data.Number", float64(1))
}

func TestXML_Empty(t *testing.T) {
	NewXML([]byte(""), t).Empty()
}

func TestXML_NotEmpty(t *testing.T) {
	NewXML([]byte(WrongXMLData), t).NotEmpty()
}

// TO assure methods of XML never return JSON
func TestXML_Bind_After_Assert(t *testing.T) {
	data := new(AssertStruct)
	NewXML([]byte(XMLAssertData), t).
		Exist("data.Time").
		String("data.Time", AssertDataTimeStr).
		Int("data.Number", int64(1)).
		True("data.OK").
		False("data.NO").
		Uint("data.Number", uint64(1)).
		Time("data.Time", AssertDataTime).
		Float("data.Number", float64(1)).
		NotEmpty().
		Bind(data)
	assert.Equal(t, 1, data.Number)
	assert.Equal(t, AssertDataTime, data.Time)
	assert.True(t, data.OK)
	assert.False(t, data.NO)
}

func TestWrongXML_JSON_Empty(t *testing.T) {
	NewXML([]byte(WrongXMLData), t).JSON.Empty()
}

func TestXML_Bind(t *testing.T) {
	user := new(User)
	client := NewClient(t).To(Mux)
	client.Get("/xml_body/user").Test().StatusOK().XML().Bind(user)
	assert.Equal(t, user.Id, uint(1))
	assert.Equal(t, user.Name, "hexi")
}

func TestMD5_Expect(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Test().StatusOK().MD5().Expect(UserDataMD5)
}

func TestSHA1_Expect(t *testing.T) {
	client := NewClient(t).To(Mux)
	client.Get("/body/user").Test().StatusOK().SHA1().Expect(UserDataSHA1)
}

func UserDataHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserData)
}

func UserDataXMLHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, UserDataXML)
}
