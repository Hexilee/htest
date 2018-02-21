package htest

import (
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

func (b *JSON) String(key, expect string) {
	assert.Equal(b.T, gjson.GetBytes(b.body, key).String(), expect)
}
