package memento

import (
	"fmt"
)

type gameState struct {
	scene string
	level int
	money int
}

func newGameState(rg *rpgGame) *gameState {
	return &gameState{
		scene: rg.scene,
		level: rg.level,
		money: rg.money,
	}
}

type rpgGame struct {
	scene string
	level int
	money int
}

func (rg *rpgGame) saveGameState() *gameState {
	return newGameState(rg)
}

func (rg *rpgGame) restoreGameState(gs *gameState) {
	rg.scene = gs.scene
	rg.level = gs.level
	rg.money = gs.money
}

func (rg *rpgGame) String() string {
	return fmt.Sprintf("scene: %s, level: %d, money: %d", rg.scene, rg.level, rg.money)
}

type gameStateManager struct {
	gs *gameState
}

func (gsm *gameStateManager) saveGameState(gs *gameState) {
	gsm.gs = gs
}

func (gsm *gameStateManager) getGameState() *gameState {
	return gsm.gs
}
