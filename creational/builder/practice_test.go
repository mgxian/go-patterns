package builder

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateFullModeVideoPlayer(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var vpb = new(fullModeBuilder)
	vc := newVidePlayerController(vpb)
	vc.construct()

	assert.Equal(t, "create main window\ncreate menu\ncreate playlist\ncreate control bar\n", outputWriter.(*strings.Builder).String())
}

func TestCreateTidyModeVideoPlayer(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var vpb = new(tidyModeBuilder)
	vc := newVidePlayerController(vpb)
	vc.construct()

	assert.Equal(t, "create main window\ncreate control bar\n", outputWriter.(*strings.Builder).String())
}

func TestCreateMemoryModeVideoPlayer(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var vpb = new(memoryModeBuilder)
	vc := newVidePlayerController(vpb)
	vc.construct()

	assert.Equal(t, "create main window\ncreate control bar\ncreate favorites\n", outputWriter.(*strings.Builder).String())
}
