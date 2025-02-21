package domain

import (
	"ebitengine-othello/src/config"
)

type GameStatus struct {
	Board [8][8]int
	Side  int
}

func (g *GameStatus) Place(col int, row int) {
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

	return
}