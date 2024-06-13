package maze

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/maripiyoko/my-first-mobile-game/resources/images"
)

const (
	tileSize  = 16
	tileCount = 19
)

const (
	screenWidth  = tileSize * tileCount
	screenHeight = tileSize * tileCount

	frameCount = 5
)

const (
	dotImageIndex    = 1
	wallImageIndex   = 2
	playerImageIndex = 3
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

type Player struct {
	x, y  int
	score int
}

type PositionXY struct {
	x, y int
}

type Game struct {
	layers      [][]int
	maxDotCount int
	player      Player

	touchIds         []ebiten.TouchID
	touchIdsPosition map[ebiten.TouchID]PositionXY
	clickedPosition  PositionXY
	debugMsg         string
	touchPositionX   int
	touchPositionY   int
	count            int
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		//fmt.Printf("cursor x=%d,y=%d", x, y)
		g.clickedPosition = PositionXY{x: x, y: y}
	}

	g.count++
	g.touchIds = ebiten.AppendTouchIDs(g.touchIds[:0])
	g.touchIdsPosition = map[ebiten.TouchID]PositionXY{}
	for _, id := range g.touchIds {
		x, y := ebiten.TouchPosition(id)
		g.touchPositionX = x
		g.touchPositionY = y
		g.touchIdsPosition[id] = PositionXY{x: x, y: y}
		CheckTappedDirection(g)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
		PlayerGoDown(g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyK) {
		PlayerGoUp(g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyL) {
		PlayerGoRight(g)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyH) {
		PlayerGoLeft(g)
	}
	return nil
}

func EatDot(g *Game, index int) {
	//fmt.Printf("EatDot at %d\n", index)
	g.player.score++
	g.layers[0][index] = 0
}

func PlayerGoDown(g *Game) {
	destx := g.player.x
	var desty int
	if g.player.y < tileCount-1 {
		desty = g.player.y + 1
	} else {
		desty = 0
	}
	index := destx + desty*tileCount
	if index < tileCount*tileCount {
		dest := g.layers[0][index]
		//fmt.Printf("%dx%d=%d", destx, desty, dest)
		if dest < wallImageIndex {
			g.player.y = desty
		}
		if dest == dotImageIndex {
			EatDot(g, index)
		}
	}
}

func PlayerGoUp(g *Game) {
	destx := g.player.x
	var desty int
	if g.player.y > 0 {
		desty = g.player.y - 1
	} else {
		desty = tileCount - 1
	}
	floor := g.layers[0]
	index := destx + desty*tileCount
	if index < tileCount*tileCount {
		dest := floor[index]
		//fmt.Printf("%dx%d=%d", destx, desty, dest)
		if dest < wallImageIndex {
			g.player.y = desty
		}
		if dest == dotImageIndex {
			EatDot(g, index)
		}
	}
}

func PlayerGoLeft(g *Game) {
	desty := g.player.y
	var destx int
	if g.player.x > 0 {
		destx = g.player.x - 1
	} else {
		destx = tileCount - 1
	}
	floor := g.layers[0]
	index := destx + desty*tileCount
	if index < tileCount*tileCount {
		dest := floor[index]
		//fmt.Printf("%dx%d=%d", destx, desty, dest)
		if dest < wallImageIndex {
			g.player.x = destx
		}
		if dest == dotImageIndex {
			EatDot(g, index)
		}
	}
}

func PlayerGoRight(g *Game) {
	desty := g.player.y
	var destx int
	if g.player.x < tileCount-1 {
		destx = g.player.x + 1
	} else {
		destx = 0
	}
	floor := g.layers[0]
	index := destx + desty*tileCount
	if index < tileCount*tileCount {
		dest := floor[index]
		//fmt.Printf("%dx%d=%d", destx, desty, dest)
		if dest < wallImageIndex {
			g.player.x = destx
		}
		if dest == dotImageIndex {
			EatDot(g, index)
		}
	}
}

// touchPositionX, touchPositionY と Player.x,y からどちらの方向が選択されたか判定する
func CheckTappedDirection(g *Game) {
	if g.count%frameCount > 0 {
		return
	}

	if g.touchPositionX < 150 {
		if g.touchPositionY > 150 {
			// DOWN
			PlayerGoDown(g)
			g.debugMsg = "DOWN"
		} else {
			// LEFT
			PlayerGoLeft(g)
			g.debugMsg = "LEFT"
		}
	} else {
		if g.touchPositionY > 150 {
			// RIGHT
			PlayerGoRight(g)
			g.debugMsg = "RIGHT"
		} else {
			// UP
			PlayerGoUp(g)
			g.debugMsg = "UP"

		}
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawFloor(g, screen)
	DrawPlayer(g, screen)
	// DrawSystemInfo(g, screen)
	DrawTouchPositionDebugPrint(g, screen)
	//DrawClickedPositonDebugPrint(g, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func NewGame() (*Game, error) {
	game := &Game{
		player: Player{x: 9, y: 7},
		layers: [][]int{
			{
				2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2,
				2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 2, 11, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 8, 2, 1, 2,
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
				2, 1, 2, 9, 2, 1, 2, 1, 1, 1, 1, 1, 2, 1, 2, 10, 2, 1, 2,
				2, 1, 2, 2, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 2, 2, 1, 2,
				2, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 1, 1, 1, 1, 1, 1, 1, 2,
				2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2,
			},
			{},
		},
	}

	var count int
	for _, v := range game.layers[0] {
		if v == 1 {
			count++
		}
	}
	game.maxDotCount = count

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Maze (Dot eater)")
	return game, nil
}

func DrawFloor(g *Game, screen *ebiten.Image) {
	w := floorsImage.Bounds().Dx()
	tileXCount := w / tileSize
	for _, l := range g.layers {
		for i, t := range l {
			op := &ebiten.DrawImageOptions{}
			xIndex := i % tileCount
			yIndex := i / tileCount
			//fmt.Printf("floor xIndex %d, yIndex %d\n", xIndex, yIndex)
			op.GeoM.Translate(float64(xIndex*tileSize), float64(yIndex*tileSize))

			sx := (t % tileXCount) * tileSize
			sy := (t / tileXCount) * tileSize
			//fmt.Printf("floor sx=%d, sy=%d\n", sx, sy)

			rect := image.Rect(sx, sy, sx+tileSize, sy+tileSize)
			subImage := floorsImage.SubImage(rect)
			screen.DrawImage(subImage.(*ebiten.Image), op)
		}
		//fmt.Println("")
	}
	//fmt.Println("------")
}

func DrawPlayer(g *Game, screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(g.player.x*tileSize), float64(g.player.y*tileSize))
	sx := playerImageIndex * tileSize
	sy := 0
	//fmt.Printf("player sx=%d, sy=%d\n", sx, sy)
	rect := image.Rect(sx, 0, sx+tileSize, sy+tileSize)
	subImage := floorsImage.SubImage(rect)
	screen.DrawImage(subImage.(*ebiten.Image), op)
}

func DrawTouchPositionDebugPrint(g *Game, screen *ebiten.Image) {
	var msg = "touch :"
	/*for i, id := range g.touchIds {
		pos := g.touchIdsPosition[id]
		msg += fmt.Sprintf("[%d] ID=%d (%d,%d) ", i, id, pos.x, pos.y)
	}*/
	msg += fmt.Sprintf(" pos x=%d,y=%d ", g.touchPositionX, g.touchPositionY)
	msg += fmt.Sprintf(" player x=%d,y=%d ", g.player.x, g.player.y)
	msg += g.debugMsg
	ebitenutil.DebugPrint(screen, msg)
}

func DrawClickedPositonDebugPrint(g *Game, screen *ebiten.Image) {
	msg := fmt.Sprintf("Clicked : (%d,%d)", g.clickedPosition.x, g.clickedPosition.y)
	ebitenutil.DebugPrint(screen, msg)
}

func DrawSystemInfo(g *Game, screen *ebiten.Image) {
	if g.player.score == g.maxDotCount {
		ebitenutil.DebugPrint(screen, "Win!!")
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.1f score %d/%d", ebiten.ActualTPS(), g.player.score, g.maxDotCount))
	}
}
