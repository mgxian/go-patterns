package factory_method

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateGifImageReader(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var imageReaderFactory = newGifImageReaderFactory()
	var imageReader = imageReaderFactory.createImageReader()
	imageReader.readFile()

	assert.Equal(t, "read gif image\n", outputWriter.(*strings.Builder).String())
}

func TestCreateJpgImageReader(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var imageReaderFactory = newJpgImageReaderFactory()
	var imageReader = imageReaderFactory.createImageReader()
	imageReader.readFile()

	assert.Equal(t, "read jpg image\n", outputWriter.(*strings.Builder).String())
}

func TestCreatePngImageReader(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	var imageReaderFactory = newPngImageReaderFactory()
	var imageReader = imageReaderFactory.createImageReader()
	imageReader.readFile()

	assert.Equal(t, "read png image\n", outputWriter.(*strings.Builder).String())
}
