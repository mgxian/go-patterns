package flyweight

import (
	"strings"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestSharePic(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	pf := newPicFactory()
	pic1 := pf.getPic("mgxian", "/tmp/will.png")
	pic2 := pf.getPic("mgxian", "/tmp/will.png")

	assert.Equal(t, unsafe.Pointer(pic1), unsafe.Pointer(pic2))

	pic1.show(position{1, 2}, size{10, 10})
	assert.Equal(t, "pic name is mgxian path is /tmp/will.png position (1, 2) and size (10, 10)\n",
		outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	pic2.show(position{3, 1}, size{50, 50})
	assert.Equal(t, "pic name is mgxian path is /tmp/will.png position (3, 1) and size (50, 50)\n",
		outputWriter.(*strings.Builder).String())
}

func TestShareCartoon(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	cf := newCartoonFactory()
	cartoon1 := cf.getCartoon("mgxian", "/tmp/will.flv")
	cartoon2 := cf.getCartoon("mgxian", "/tmp/will.flv")

	assert.Equal(t, unsafe.Pointer(cartoon1), unsafe.Pointer(cartoon2))

	cartoon1.show(position{1, 2}, size{10, 10})
	assert.Equal(t, "cartoon name is mgxian path is /tmp/will.flv position (1, 2) and size (10, 10)\n",
		outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	cartoon2.show(position{3, 1}, size{50, 50})
	assert.Equal(t, "cartoon name is mgxian path is /tmp/will.flv position (3, 1) and size (50, 50)\n",
		outputWriter.(*strings.Builder).String())
}

func TestShareVideo(t *testing.T) {
	oriOutputWriter := outputWriter
	outputWriter = new(strings.Builder)
	defer func() { outputWriter = oriOutputWriter }()

	vf := newVideoFactory()
	video1 := vf.getVideo("mgxian", "/tmp/will.mp4")
	video2 := vf.getVideo("mgxian", "/tmp/will.mp4")

	assert.Equal(t, unsafe.Pointer(video1), unsafe.Pointer(video2))

	video1.show(position{1, 2}, size{10, 10})
	assert.Equal(t, "video name is mgxian path is /tmp/will.mp4 position (1, 2) and size (10, 10)\n",
		outputWriter.(*strings.Builder).String())

	outputWriter.(*strings.Builder).Reset()
	video2.show(position{3, 1}, size{50, 50})
	assert.Equal(t, "video name is mgxian path is /tmp/will.mp4 position (3, 1) and size (50, 50)\n",
		outputWriter.(*strings.Builder).String())
}
