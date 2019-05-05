package factory_method

import (
	"fmt"
	"io"
	"os"
)

// collect stdout
var outputWriter io.Writer = os.Stdout

type imageReader interface {
	readFile()
}

type imageReaderFactory interface {
	createImageReader() imageReader
}

type gifImageReader struct {
}

func (g *gifImageReader) readFile() {
	fmt.Fprintln(outputWriter, "read gif image")
}

type jpgImageReader struct {
}

func (j *jpgImageReader) readFile() {
	fmt.Fprintln(outputWriter, "read jpg image")
}

type pngImageReader struct {
}

func (p *pngImageReader) readFile() {
	fmt.Fprintln(outputWriter, "read png image")
}

type gifImageReaderFactory struct {
}

func newGifImageReaderFactory() *gifImageReaderFactory {
	return &gifImageReaderFactory{}
}

func (g *gifImageReaderFactory) createImageReader() imageReader {
	return &gifImageReader{}
}

type jpgImageReaderFactory struct {
}

func newJpgImageReaderFactory() *jpgImageReaderFactory {
	return &jpgImageReaderFactory{}
}

func (j *jpgImageReaderFactory) createImageReader() imageReader {
	return &jpgImageReader{}
}

type pngImageReaderFactory struct {
}

func newPngImageReaderFactory() *pngImageReaderFactory {
	return &pngImageReaderFactory{}
}

func (p *pngImageReaderFactory) createImageReader() imageReader {
	return &pngImageReader{}
}
