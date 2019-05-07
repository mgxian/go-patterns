package proxy

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPhotoDownloader(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	pd := newPhotoDownloader()
	pd.getPhotos("http://mgxian.io/")
	pd.showAll()

	expected := "show icon png.icon\ndownload photo http://will.com/1.png\nshow photo http://will.com/1.png\n" +
		"show icon png.icon\ndownload photo http://will.com/2.png\nshow photo http://will.com/2.png\n" +
		"show icon gif.icon\ndownload photo http://will.com/3.gif\nshow photo http://will.com/3.gif\n" +
		"show icon gif.icon\ndownload photo http://will.com/4.gif\nshow photo http://will.com/4.gif\n" +
		"show icon jpg.icon\ndownload photo http://will.com/5.jpg\nshow photo http://will.com/5.jpg\n" +
		"show icon jpg.icon\ndownload photo http://will.com/6.jpg\nshow photo http://will.com/6.jpg\n"

	assert.Equal(t, expected, outputWriter.(*strings.Builder).String())
}
