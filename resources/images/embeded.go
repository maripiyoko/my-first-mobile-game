package images

import (
	_ "embed"
)

var (
	//go:embed tiles.png
	Tiles_png []byte

	//go:embed walls.png
	Walls_png []byte
)
