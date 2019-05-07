package flyweight

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type mediaResource interface {
	show(aPosition position, aSize size)
}

type picFactory struct {
	pics map[string]*pic
}

func newPicFactory() *picFactory {
	return &picFactory{
		pics: make(map[string]*pic, 0),
	}
}

func (pf *picFactory) getPic(name string, path string) *pic {
	picKey := name + " " + path
	if pic, ok := pf.pics[picKey]; ok {
		return pic
	}

	pic := newPic(name, path)
	pf.pics[picKey] = pic
	return pic
}

type pic struct {
	name string
	path string
}

func newPic(name string, path string) *pic {
	return &pic{
		name: name,
		path: path,
	}
}

func (p *pic) show(aPosition position, aSize size) {
	fmt.Fprintf(outputWriter, "pic name is %s path is %s position (%d, %d) and size (%d, %d)\n",
		p.name, p.path, aPosition.row, aPosition.column, aSize.width, aSize.height)
}

type cartoonFactory struct {
	cartoons map[string]*cartoon
}

func newCartoonFactory() *cartoonFactory {
	return &cartoonFactory{
		cartoons: make(map[string]*cartoon, 0),
	}
}

func (cf *cartoonFactory) getCartoon(name string, path string) *cartoon {
	cartoonKey := name + " " + path
	if cartoon, ok := cf.cartoons[cartoonKey]; ok {
		return cartoon
	}

	cartoon := newCartoon(name, path)
	cf.cartoons[cartoonKey] = cartoon
	return cartoon
}

type cartoon struct {
	name string
	path string
}

func newCartoon(name string, path string) *cartoon {
	return &cartoon{
		name: name,
		path: path,
	}
}

func (c *cartoon) show(aPosition position, aSize size) {
	fmt.Fprintf(outputWriter, "cartoon name is %s path is %s position (%d, %d) and size (%d, %d)\n",
		c.name, c.path, aPosition.row, aPosition.column, aSize.width, aSize.height)
}

type videoFactory struct {
	videos map[string]*video
}

func newVideoFactory() *videoFactory {
	return &videoFactory{
		videos: make(map[string]*video, 0),
	}
}

func (vf *videoFactory) getVideo(name string, path string) *video {
	videoKey := name + " " + path
	if video, ok := vf.videos[videoKey]; ok {
		return video
	}

	video := newVideo(name, path)
	vf.videos[videoKey] = video
	return video
}

type video struct {
	name string
	path string
}

func newVideo(name string, path string) *video {
	return &video{
		name: name,
		path: path,
	}
}

func (v *video) show(aPosition position, aSize size) {
	fmt.Fprintf(outputWriter, "video name is %s path is %s position (%d, %d) and size (%d, %d)\n",
		v.name, v.path, aPosition.row, aPosition.column, aSize.width, aSize.height)
}

type position struct {
	row    int
	column int
}

type size struct {
	width  int
	height int
}
