package proxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProxy(t *testing.T) {
	proxy := newProxy()
	result := proxy.request()
	assert.Equal(t, "preRequest realRequest postRequest", result)
}
