package htest

import (
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
	"testing"
	"encoding/json"
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

func (j *JSON) String(key, expect string) *JSON {
	assert.True(j.T, gjson.GetBytes(j.body, key).Exists())
	assert.Equal(j.T, gjson.GetBytes(j.body, key).String(), expect)
	return j
}

func (j *JSON) Bind(obj interface{}) error {
	return json.Unmarshal(j.body, obj)
}
