package htest

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
)

type (
	JSON struct {
		body []byte
		*testing.T
	}
)

func NewJSON(body []byte, t *testing.T) *JSON {
	return &JSON{
		body: body,
		T:    t,
	}
}

func (j *JSON) GetJSON(key string) (result gjson.Result, exist bool) {
	result = gjson.GetBytes(j.body, key)
	exist = result.Exists()
	return
}

func (j *JSON) Exist(key string) *JSON {
	_, exist := j.GetJSON(key)
	assert.True(j.T, exist)
	return j
}

func (j *JSON) NotExist(key string) *JSON {
	_, exist := j.GetJSON(key)
	assert.False(j.T, exist)
	return j
}

func (j *JSON) String(key, expect string) *JSON {
	result, exist := j.GetJSON(key)
	if exist {
		assert.Equal(j.T, result.String(), expect)
	}
	return j
}

func (j *JSON) Bind(obj interface{}) error {
	return json.Unmarshal(j.body, obj)
}
