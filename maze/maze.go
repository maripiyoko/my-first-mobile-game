package maze

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
	tileSize  = 16
	tileCount = 19
)

const (
	screenWidth  = tileSize * tileCount
	screenHeight = tileSize * tileCount
)

var (
	floorsImage *ebiten.Image
)

func init() {
	img, _, err := image.Decode(bytes.NewReader(images.Floors_png))
	if err != nil {
		log.Fatal(err)
	}
	floorsImage = ebiten.NewImageFromImage(img)
}

type Game struct {
	layers [][]int
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	w := floorsImage.Bounds().Dx()
	tileXCount := w / tileSize
	const xCount = screenWidth / tileSize

	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			xIndex := i % xCount
			yIndex := i / xCount
			op.GeoM.Translate(float64(xIndex*tileSize), float64(yIndex*tileSize))

			/*if xIndex == 0 {
				fmt.Printf("\n%d ", t)
			} else {
				fmt.Printf("%d ", t)
			}*/

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize

			rect := image.Rect(sx, sy, sx+tileSize, sy+tileSize)
			subImage := floorsImage.SubImage(rect)
			screen.DrawImage(subImage.(*ebiten.Image), op)
		}
		//fmt.Println("")
	}
	//fmt.Println("------")
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.1f", ebiten.ActualTPS()))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{
		layers: [][]int{
			{
				2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2,
				2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 2, 0, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 0, 2, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2,
				2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2,
				1, 1, 1, 1, 1, 1, 1, 1, 2, 0, 2, 1, 1, 1, 1, 1, 1, 1, 1,
				2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 2, 2,
				2, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 2, 0, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 0, 2, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2,
				2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			},
			{},
		},
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Maze (Dot eater)")
	return game, nil
}
