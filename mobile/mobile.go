package mobile

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"

	//"github.com/maripiyoko/my-first-mobile-game/hello"
	//"github.com/maripiyoko/my-first-mobile-game/tiles"
	"github.com/maripiyoko/my-first-mobile-game/maze"
)

func init() {
	game, err := maze.NewGame()
	if err != nil {
		panic(err)
	}

	mobile.SetGame(game)
}

func Dummy() {}
