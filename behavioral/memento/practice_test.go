package memento

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRPGGame(t *testing.T) {
	rg := &rpgGame{}
	rg.scene = "changan"
	rg.level = 19
	rg.money = 2000

	assert.Equal(t, "scene: changan, level: 19, money: 2000", rg.String())

	gsm := &gameStateManager{}
	gsm.saveGameState(rg.saveGameState())

	rg.scene = "wudang"
	rg.level = 20
	rg.money = 2100

	assert.Equal(t, "scene: wudang, level: 20, money: 2100", rg.String())

	rg.restoreGameState(gsm.getGameState())
	assert.Equal(t, "scene: changan, level: 19, money: 2000", rg.String())
}
