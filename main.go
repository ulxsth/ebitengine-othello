package main

import (
	"errors"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	SCREEN_WIDTH  = 330
	SCREEN_HEIGHT = 330

	CELL_EMPTY = 0
	CELL_BLACK = 1
	CELL_WHITE = 2
)

var board = [8][8]int{
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 1, 2, 0, 0, 0},
	{0, 0, 0, 2, 1, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0},
}

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
	rectLength := (SCREEN_WIDTH - borderWidth*2) / 8
	for i := 1; i < 8; i++ {
		// 縦
		vector.StrokeLine(
			screen,
			float32(gridWidth*i+rectLength*i),
			0,
			float32(gridWidth*i+rectLength*i),
			float32(SCREEN_HEIGHT),
			float32(gridWidth),
			color.Black,
			false,
		)

		// 横
		vector.StrokeLine(
			screen,
			0,
			float32(gridWidth*i+rectLength*i),
			float32(SCREEN_WIDTH),
			float32(gridWidth*i+rectLength*i),
			float32(gridWidth),
			color.Black,
			false,
		)
	}

	// オセロの石を描画
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := board[y][x]
			if piece == CELL_EMPTY {
				continue
			}

			var pieceColor color.Color
			if piece == CELL_BLACK {
				pieceColor = color.Black
			} else if piece == CELL_WHITE {
				pieceColor = color.White
			} else {
				err := errors.New("invalid piece")
				log.Fatal(err)
			}

			vector.DrawFilledCircle(
				screen,
				float32(SCREEN_WIDTH*x + gridWidth*x + borderWidth),
				float32(SCREEN_HEIGHT*y + gridWidth*y + borderWidth),				
				float32(rectLength / 2),
				pieceColor,
				true,
			)
		}
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
