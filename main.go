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

	BORDER_WIDTH = 5
	RECT_WIDTH = SCREEN_WIDTH - BORDER_WIDTH*2
	RECT_HEIGHT = SCREEN_HEIGHT - BORDER_WIDTH*2
	GRID_WIDTH = 1
	RECT_LENGTH = (SCREEN_WIDTH - BORDER_WIDTH*2) / 8

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
	// 駒設置
	cursorX, cursorY := ebiten.CursorPosition()
	if(cursorX < 0 || cursorY < 0 || cursorX > SCREEN_WIDTH || cursorY > SCREEN_HEIGHT) {
		return nil
	}
	
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// オセロ盤のベースの部分（緑の四角）
	vector.DrawFilledRect(
		screen,
		float32(BORDER_WIDTH),
		float32(BORDER_WIDTH),
		float32(RECT_WIDTH),
		float32(RECT_HEIGHT),
		color.RGBA{0x00, 0x80, 0x00, 0xff},
		false,
	)

	// グリッドを引く
	for i := 1; i < 8; i++ {
		// 縦
		vector.StrokeLine(
			screen,
			float32(GRID_WIDTH*i+RECT_LENGTH*i),
			0,
			float32(GRID_WIDTH*i+RECT_LENGTH*i),
			float32(SCREEN_HEIGHT),
			float32(GRID_WIDTH),
			color.Black,
			false,
		)

		// 横
		vector.StrokeLine(
			screen,
			0,
			float32(GRID_WIDTH*i+RECT_LENGTH*i),
			float32(SCREEN_WIDTH),
			float32(GRID_WIDTH*i+RECT_LENGTH*i),
			float32(GRID_WIDTH),
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
				float32(RECT_LENGTH/2 + RECT_LENGTH*x + GRID_WIDTH*x),
				float32(RECT_LENGTH/2 + RECT_LENGTH*y + GRID_WIDTH*y),				
				float32(RECT_LENGTH / 2) - 2,
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
