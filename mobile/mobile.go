package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	//"github.com/maripiyoko/my-first-mobile-game/hello"
	"github.com/maripiyoko/my-first-mobile-game/tiles"
)

func init() {
	game, err := tiles.NewGame()
	if err != nil {
		panic(err)
	}

	mobile.SetGame(game)
}

func Dummy() {}
