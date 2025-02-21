package usecase

import (
	"ebitengine-othello/src/config"
	"ebitengine-othello/src/domain"
)

func Place(g *domain.GameStatus, col int, row int) {
	// クリックした位置に駒がある場合は何もしない
	if g.Board[col][row] != config.CELL_EMPTY {
		return
	}

	g.Board[col][row] = g.Side
	if g.Side == config.CELL_BLACK {
		g.Side = config.CELL_WHITE
	} else {
		g.Side = config.CELL_BLACK
	}
}