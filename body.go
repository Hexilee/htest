package htest

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"github.com/basgys/goxml2json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"testing"
)

type (
	JSON struct {
		body []byte
		*testing.T
	}

	XML struct {
		*JSON
		body []byte
	}
)

func NewJSON(body []byte, t *testing.T) *JSON {
	return &JSON{
		body: body,
		T:    t,
	}
}

func NewXML(body []byte, t *testing.T) *XML {
	jsonBuf, _ := xml2json.Convert(bytes.NewBuffer(body))
	jsonBody, _ := ioutil.ReadAll(jsonBuf)
	return &XML{
		body: body,
		JSON: NewJSON(jsonBody, t),
	}
}

func (j *JSON) GetKey(key string) (result gjson.Result, exist bool) {
	result = gjson.GetBytes(j.body, key)
	exist = result.Exists()
	return
}

func (j *JSON) Exist(key string) *JSON {
	_, exist := j.GetKey(key)
	assert.True(j.T, exist)
	return j
}

func (j *JSON) NotExist(key string) *JSON {
	_, exist := j.GetKey(key)
	assert.False(j.T, exist)
	return j
}

func (j *JSON) String(key, expect string) *JSON {
	result, _ := j.GetKey(key)
	assert.Equal(j.T, result.String(), expect)
	return j
}

func (j *JSON) Empty() *JSON {
	body := bytes.Trim(j.Body(), "\"\n")
	assert.Equal(j.T, "", string(body))
	return j
}

func (j *JSON) NotEmpty() *JSON {
	body := bytes.Trim(j.Body(), "\"\n")
	assert.NotEqual(j.T, "", string(body))
	return j
}

func (j *JSON) Body() []byte {
	return j.body
}

func (j *JSON) Bind(obj interface{}) error {
	return json.Unmarshal(j.body, obj)
}

func (x *XML) Exist(key string) *XML {
	x.JSON.Exist(key)
	return x
}

func (x *XML) NotExist(key string) *XML {
	x.JSON.NotExist(key)
	return x
}
func (x *XML) String(key, expect string) *XML {
	x.JSON.String(key, expect)
	return x
}

func (x *XML) Empty() *XML {
	assert.Equal(x.T, "", string(x.Body()))
	return x
}

func (x *XML) NotEmpty() *XML {
	assert.NotEqual(x.T, "", string(x.Body()))
	return x
}

func (x *XML) Body() []byte {
	return x.body
}

func (x *XML) Bind(obj interface{}) error {
	return xml.Unmarshal(x.body, obj)
}
