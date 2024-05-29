package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	"github.com/maripiyoko/my-first-mobile-game/hello"
)

func init() {
	game, err := hello.NewGame()
	if err != nil {
		panic(err)
	}

	mobile.SetGame(game)
}

func Dummy() {}
