package builder

import (
	"fmt"
	"io"
	"os"
)

var outputWriter io.Writer = os.Stdout

type videPlayerController struct {
	idb videPlayerBuilder
}

func newVidePlayerController(idb videPlayerBuilder) *videPlayerController {
	return &videPlayerController{
		idb: idb,
	}
}

func (vpc *videPlayerController) construct() *videoPlayer {
	vpc.idb.buildMainWindow()
	vpc.idb.buildMenu()
	vpc.idb.buildPlaylist()
	vpc.idb.buildControlBar()
	vpc.idb.buildFavorites()
	return vpc.idb.getPlayer()
}

type videPlayerBuilder interface {
	buildMenu()
	buildPlaylist()
	buildMainWindow()
	buildControlBar()
	buildFavorites()
	getPlayer() *videoPlayer
}

type baseModeBuilder struct {
	vp *videoPlayer
}

func (bb *baseModeBuilder) buildMainWindow() {
	bb.vp.createMainWindow(false)
}

func (bb *baseModeBuilder) buildControlBar() {
	bb.vp.createControlBar(false)
}

func (bb *baseModeBuilder) buildMenu() {
	bb.vp.createMenu(false)
}

func (bb *baseModeBuilder) buildPlaylist() {
	bb.vp.createPlaylist(false)
}

func (bb *baseModeBuilder) buildFavorites() {
	bb.vp.createFavorites(false)
}

func (bb *baseModeBuilder) getPlayer() *videoPlayer {
	return bb.vp
}

type fullModeBuilder struct {
	baseModeBuilder
	menu, playlist, mainWindow, controlBar string
}

func (fb *fullModeBuilder) buildMainWindow() {
	fb.vp.createMainWindow(true)
}

func (fb *fullModeBuilder) buildControlBar() {
	fb.vp.createControlBar(true)
}

func (fb *fullModeBuilder) buildMenu() {
	fb.vp.createMenu(true)
}

func (fb *fullModeBuilder) buildPlaylist() {
	fb.vp.createPlaylist(true)
}

type tidyModeBuilder struct {
	baseModeBuilder
	mainWindow, controlBar string
}

func (tb *tidyModeBuilder) buildMainWindow() {
	tb.vp.createMainWindow(true)
}

func (tb *tidyModeBuilder) buildControlBar() {
	tb.vp.createControlBar(true)
}

type memoryModeBuilder struct {
	baseModeBuilder
	mainWindow, controlBar, favorites string
}

func (mb *memoryModeBuilder) buildMainWindow() {
	mb.vp.createMainWindow(true)
}

func (mb *memoryModeBuilder) buildControlBar() {
	mb.vp.createControlBar(true)
}

func (mb *memoryModeBuilder) buildFavorites() {
	mb.vp.createFavorites(true)
}

type videoPlayer struct {
}

func newVideoPlayer() *videoPlayer {
	return &videoPlayer{}
}

func (vp *videoPlayer) createMenu(create bool) {
	if create {
		fmt.Fprintln(outputWriter, "create menu")
	}
}

func (vp *videoPlayer) createPlaylist(create bool) {
	if create {
		fmt.Fprintln(outputWriter, "create playlist")
	}
}

func (vp *videoPlayer) createMainWindow(create bool) {
	if create {
		fmt.Fprintln(outputWriter, "create main window")
	}
}

func (vp *videoPlayer) createControlBar(create bool) {
	if create {
		fmt.Fprintln(outputWriter, "create control bar")
	}
}

func (vp *videoPlayer) createFavorites(create bool) {
	if create {
		fmt.Fprintln(outputWriter, "create favorites")
	}
}
