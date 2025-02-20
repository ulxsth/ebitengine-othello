package main

import (
	"errors"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"ebitengine-othello/src/config"
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
var turn = config.CELL_BLACK

type Game struct{}

func (g *Game) Update() error {
	// ウィンドウ外にカーソルがある場合は何もしない
	cursorX, cursorY := ebiten.CursorPosition()
	if(cursorX < 0 || cursorY < 0 || cursorX > config.SCREEN_WIDTH || cursorY > config.SCREEN_HEIGHT) {
		return nil
	}

	// 駒設置
	targetRow := cursorY / (config.RECT_LENGTH+config.GRID_WIDTH)
	targetCol := cursorX / (config.RECT_LENGTH+config.GRID_WIDTH)
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return nil
	}
	
	if board[targetRow][targetCol] != config.CELL_EMPTY {
		return nil	
	}

	board[targetRow][targetCol] = turn
	if turn == config.CELL_BLACK {
		turn = config.CELL_WHITE
	} else {
		turn = config.CELL_BLACK
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// オセロ盤のベースの部分（緑の四角）
	vector.DrawFilledRect(
		screen,
		float32(config.BORDER_WIDTH),
		float32(config.BORDER_WIDTH),
		float32(config.RECT_WIDTH),
		float32(config.RECT_HEIGHT),
		color.RGBA{0x00, 0x80, 0x00, 0xff},
		false,
	)

	// グリッドを引く
	for i := 1; i < 8; i++ {
		// 縦
		vector.StrokeLine(
			screen,
			float32(config.GRID_WIDTH*i+config.RECT_LENGTH*i),
			0,
			float32(config.GRID_WIDTH*i+config.RECT_LENGTH*i),
			float32(config.SCREEN_HEIGHT),
			float32(config.GRID_WIDTH),
			color.Black,
			false,
		)

		// 横
		vector.StrokeLine(
			screen,
			0,
			float32(config.GRID_WIDTH*i+config.RECT_LENGTH*i),
			float32(config.SCREEN_WIDTH),
			float32(config.GRID_WIDTH*i+config.RECT_LENGTH*i),
			float32(config.GRID_WIDTH),
			color.Black,
			false,
		)
	}

	// オセロの石を描画
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := board[y][x]
			if piece == config.CELL_EMPTY {
				continue
			}

			var pieceColor color.Color
			if piece == config.CELL_BLACK {
				pieceColor = color.Black
			} else if piece == config.CELL_WHITE {
				pieceColor = color.White
			} else {
				err := errors.New("invalid piece")
				log.Fatal(err)
			}

			vector.DrawFilledCircle(
				screen,
				float32(config.RECT_LENGTH/2 + config.RECT_LENGTH*x + config.GRID_WIDTH*x),
				float32(config.RECT_LENGTH/2 + config.RECT_LENGTH*y + config.GRID_WIDTH*y),				
				float32(config.RECT_LENGTH / 2) - 2,
				pieceColor,
				true,
			)
		}
	}

	// ホバーした位置に薄く駒を表示
	cursorX, cursorY := ebiten.CursorPosition()
	targetRow := cursorY / (config.RECT_LENGTH+config.GRID_WIDTH)
	targetCol := cursorX / (config.RECT_LENGTH+config.GRID_WIDTH)
	if targetRow < 0 || targetRow >= 8 || targetCol < 0 || targetCol >= 8 {
		return
	}

	if board[targetRow][targetCol] == config.CELL_EMPTY {
		var pieceColor color.Color
		if turn == config.CELL_BLACK {
			pieceColor = color.RGBA{0x00, 0x00, 0x00, 0x77}
		} else if turn == config.CELL_WHITE {
			pieceColor = color.RGBA{0xaa, 0xaa, 0xaa, 0x77}
		} else {
			err := errors.New("invalid piece")
			log.Fatal(err)
		}

		vector.DrawFilledCircle(
			screen,
			float32(config.RECT_LENGTH/2 + config.RECT_LENGTH*targetCol + config.GRID_WIDTH*targetCol),
			float32(config.RECT_LENGTH/2 + config.RECT_LENGTH*targetRow + config.GRID_WIDTH*targetRow),				
			float32(config.RECT_LENGTH / 2) - 2,
			pieceColor,
			true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_WIDTH, config.SCREEN_HEIGHT
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
