package model

import (
	"ebitengine-othello/src/config"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"log"
)

type Game struct {
	Board [8][8]int
	Turn int
}

func (g *Game) Update() error {
	// ウィンドウ外にカーソルがある場合は何もしない
	cursorX, cursorY := ebiten.CursorPosition()
	if cursorX < 0 || cursorY < 0 || cursorX > config.SCREEN_LENGTH || cursorY > config.SCREEN_LENGTH {
		return nil
	}

	// 駒設置
	targetRow := cursorY / (config.CELL_LENGTH + config.GRID_WIDTH)
	targetCol := cursorX / (config.CELL_LENGTH + config.GRID_WIDTH)
	if !ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		return nil
	}

	if g.Board[targetRow][targetCol] != config.CELL_EMPTY {
		return nil
	}

	g.Board[targetRow][targetCol] = g.Turn
	if g.Turn == config.CELL_BLACK {
		g.Turn = config.CELL_WHITE
	} else {
		g.Turn = config.CELL_BLACK
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// オセロ盤のベースの部分（緑の四角）
	vector.DrawFilledRect(
		screen,
		float32(config.OUTER_MARGIN),
		float32(config.OUTER_MARGIN),
		float32(config.BOARD_LENGTH),
		float32(config.BOARD_LENGTH),
		color.RGBA{0x00, 0x80, 0x00, 0xff},
		false,
	)

	// グリッドを引く
	for i := 1; i < 8; i++ {
		// 縦
		vector.StrokeLine(
			screen,
			float32(config.GRID_WIDTH*i+config.CELL_LENGTH*i),
			0,
			float32(config.GRID_WIDTH*i+config.CELL_LENGTH*i),
			float32(config.SCREEN_LENGTH),
			float32(config.GRID_WIDTH),
			color.Black,
			false,
		)

		// 横
		vector.StrokeLine(
			screen,
			0,
			float32(config.GRID_WIDTH*i+config.CELL_LENGTH*i),
			float32(config.SCREEN_LENGTH),
			float32(config.GRID_WIDTH*i+config.CELL_LENGTH*i),
			float32(config.GRID_WIDTH),
			color.Black,
			false,
		)
	}

	// オセロの石を描画
	for y := 0; y < 8; y++ {
		for x := 0; x < 8; x++ {
			piece := g.Board[y][x]
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
				float32(config.CELL_LENGTH/2+config.CELL_LENGTH*x+config.GRID_WIDTH*x),
				float32(config.CELL_LENGTH/2+config.CELL_LENGTH*y+config.GRID_WIDTH*y),
				float32(config.CELL_LENGTH/2)-2,
				pieceColor,
				true,
			)
		}
	}

	// ホバーした位置に薄く駒を表示
	cursorX, cursorY := ebiten.CursorPosition()
	targetRow := cursorY / (config.CELL_LENGTH + config.GRID_WIDTH)
	targetCol := cursorX / (config.CELL_LENGTH + config.GRID_WIDTH)
	if targetRow < 0 || targetRow >= 8 || targetCol < 0 || targetCol >= 8 {
		return
	}

	if g.Board[targetRow][targetCol] == config.CELL_EMPTY {
		var pieceColor color.Color
		if g.Turn == config.CELL_BLACK {
			pieceColor = color.RGBA{0x00, 0x00, 0x00, 0x77}
		} else if g.Turn == config.CELL_WHITE {
			pieceColor = color.RGBA{0xaa, 0xaa, 0xaa, 0x77}
		} else {
			err := errors.New("invalid piece")
			log.Fatal(err)
		}

		vector.DrawFilledCircle(
			screen,
			float32(config.CELL_LENGTH/2+config.CELL_LENGTH*targetCol+config.GRID_WIDTH*targetCol),
			float32(config.CELL_LENGTH/2+config.CELL_LENGTH*targetRow+config.GRID_WIDTH*targetRow),
			float32(config.CELL_LENGTH/2)-2,
			pieceColor,
			true,
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return config.SCREEN_LENGTH, config.SCREEN_LENGTH
}
