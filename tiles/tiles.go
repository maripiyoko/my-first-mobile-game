package tiles

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/maripiyoko/my-first-mobile-game/resources/images"
)

const (
	screenWidth  = 240
	screenHeight = 240
)

const (
	tileSize = 16
)

var (
	tilesImage *ebiten.Image
)

func init() {
	// Decode an image from the image file's byte slice.
	img, _, err := image.Decode(bytes.NewReader(images.Walls_png))
	if err != nil {
		log.Fatal(err)
	}
	tilesImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	layers [][]int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := tilesImage.Bounds().Dx()
	tileXCount := w / tileSize

	// Draw each tile with each DrawImage call.
	// As the source images of all DrawImage calls are always same,
	// this rendering is done very efficiently.
	// For more detail, see https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Image.DrawImage
	const xCount = screenWidth / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((i%xCount)*tileSize), float64((i/xCount)*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			screen.DrawImage(tilesImage.SubImage(image.Rect(sx, sy, sx+tileSize, sy+tileSize)).(*ebiten.Image), op)
		}
	}
	//ebitenutil.DebugPrint(screen, fmt.Sprintf("w: %d", w))
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{
		layers: [][]int{
			{
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 151, 16, 16, 16, 16, 16, 16, 16, 16, 16, 151, 16, 178, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,

				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 178, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 194, 16, 16, 16, 194, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,

				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
				16, 151, 16, 16, 16, 16, 16, 16, 16, 16, 16, 178, 16, 16, 16,
				16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
			},
			{
				4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 17, 29, 29, 29, 29, 30, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 31, 16, 16, 16, 16, 1, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 31, 16, 16, 16, 16, 1, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 31, 16, 16, 16, 16, 1, 4, 4, 4, 4,

				4, 4, 4, 4, 4, 2, 3, 25, 3, 3, 15, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 142, 155, 106, 92, 142, 155, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,

				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
				4, 4, 4, 4, 4, 4, 4, 106, 92, 4, 4, 4, 4, 4, 4,
			},
		},
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Tiles (Ebitengine Demo)")
	return game, nil
}
