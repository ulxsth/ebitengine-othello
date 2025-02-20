package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	SCREEN_WIDTH  = 330
	SCREEN_HEIGHT = 330
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// オセロ盤のベースの部分（緑の四角）
	borderWidth := 5
	rectWidth := SCREEN_WIDTH - borderWidth*2
	rectHeight := SCREEN_HEIGHT - borderWidth*2
	vector.DrawFilledRect(
		screen,
		float32(borderWidth), 
		float32(borderWidth),
		float32(rectWidth),
		float32(rectHeight),
		color.RGBA{0x00, 0x80, 0x00, 0xff},
		false,
	)

	// グリッドを引く
	gridWidth := 1
	rectLength := (SCREEN_WIDTH-borderWidth*2) / 8
	for i := 1; i < 8; i++ {
		// 縦
		vector.StrokeLine(
			screen,
			float32(gridWidth + rectLength*i),
			0,
			float32(gridWidth + rectLength*i),
			float32(SCREEN_HEIGHT),
			float32(gridWidth),
			color.Black,
			false,
		)

		// 横
		vector.StrokeLine(
			screen,
			0,
			float32(gridWidth + rectLength*i),
			float32(SCREEN_WIDTH),
			float32(gridWidth + rectLength*i),
			float32(gridWidth),
			color.Black,
			false,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
