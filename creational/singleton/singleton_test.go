package singleton

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInstaceWithSingleCall(t *testing.T) {
	singleton := GetInstance()
	assert.NotNil(t, singleton)
}

func TestGetInstaceWithMultipleCall(t *testing.T) {
	singleton := GetInstance()
	newSingleton := GetInstance()
	assert.Equal(t, singleton, newSingleton)
}
