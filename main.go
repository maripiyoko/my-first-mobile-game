package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"

	//"github.com/maripiyoko/my-first-mobile-game/hello"
	//"github.com/maripiyoko/my-first-mobile-game/tiles"
	"github.com/maripiyoko/my-first-mobile-game/maze"
)

func main() {
	game, err := maze.NewGame()
	if err != nil {
		panic(err)
	}

	//ebiten.SetWindowSize(640, 480)
	//ebiten.SetWindowTitle("Hello, World!")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
