package proxy

import (
	"fmt"
	"io"
	"os"
	"sync"
)

var outputWriter io.Writer = os.Stdout

func getPhotoUrls(url string) []string {
	return []string{
		"http://will.com/1.png",
		"http://will.com/2.png",
		"http://will.com/3.gif",
		"http://will.com/4.gif",
		"http://will.com/5.jpg",
		"http://will.com/6.jpg",
	}
}

type photoIcon struct {
	suffix string
	path   string
}

func newPhotoIcon(s, p string) *photoIcon {
	return &photoIcon{
		suffix: s,
		path:   p,
	}
}

type photo interface {
	show()
}

type photoViewer struct {
	photoURL string
}

func (p *photoViewer) show() {
	fmt.Fprintln(outputWriter, "download photo "+p.photoURL)
}

type photoViewerProxy struct {
	photoURL string
	pv       *photoViewer
	pi       *photoIcon
}

func newPhotoViewerProxy(photoURL string) *photoViewerProxy {
	suffix := photoURL[len(photoURL)-3 : len(photoURL)]
	pi := &photoIcon{
		suffix: suffix,
		path:   suffix + ".icon",
	}
	return &photoViewerProxy{
		photoURL: photoURL,
		pv: &photoViewer{
			photoURL: photoURL,
		},
		pi: pi,
	}
}

func (p *photoViewerProxy) show() {
	fmt.Fprintln(outputWriter, "show icon "+p.pi.path)
	p.downloadPhoto()
	fmt.Fprintln(outputWriter, "show photo "+p.photoURL)
}

func (p *photoViewerProxy) downloadPhoto() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		p.pv.show()
	}()
	wg.Wait()
}

type photoDownloader struct {
	photos []photo
}

func newPhotoDownloader() *photoDownloader {
	return &photoDownloader{
		photos: make([]photo, 0),
	}
}

func (p *photoDownloader) getPhotos(url string) {
	photoUrls := getPhotoUrls(url)
	for _, photoURL := range photoUrls {
		pvp := newPhotoViewerProxy(photoURL)
		p.photos = append(p.photos, pvp)
	}
}

func (p *photoDownloader) showAll() {
	for _, p := range p.photos {
		p.show()
	}
}
