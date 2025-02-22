package usecase

import (
	"ebitengine-othello/src/config"
	"ebitengine-othello/src/domain"
)

func Place(g *domain.GameStatus, y int, x int) {
	// クリックした位置に駒がある場合は何もしない
	if g.Board[y][x] != config.CELL_EMPTY {
		return
	}

	// ひっくり返しが発生しないなら、設置しない
	if canPlace(g.Board, y, x, g.Side) {
		return
	}

	currentSide := g.Side
	var enemySide int
	if currentSide == config.CELL_BLACK {
		enemySide = config.CELL_WHITE
	} else {
		enemySide = config.CELL_BLACK
	}

	// 各方向ごとにひっくり返しを実施する
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			flipPositions := []map[string]int{}
			px, py := x+dx, y+dy
			for {
				// 盤外なら終了
				if px < 0 || px >= 8 || py < 0 || py >= 8 {
					break
				}

				// 空白なら終了
				if g.Board[py][px] == config.CELL_EMPTY {
					break
				}

				// 敵の駒なら記録して次へ
				if g.Board[py][px] == enemySide {
					flipPositions = append(flipPositions, map[string]int{"x": px, "y": py})
					px += dx
					py += dy
					continue
				}

				// 自分の駒なら、これまでの駒をひっくり返す
				if g.Board[py][px] == currentSide {
					if len(flipPositions) > 0 {
						for _, pos := range flipPositions {
							println("Turn over at:", pos["x"], pos["y"])
							g.Board[pos["y"]][pos["x"]] = currentSide
						}
					}
					break
				}
			}
		}
	}

	g.Board[y][x] = currentSide
	g.Side = enemySide
}

func canPlace(board [8][8]int, y int, x int, currentSide int) bool {
	// 指定セルが空かチェック
	if board[y][x] != config.CELL_EMPTY {
		return false
	}

	var enemySide int
	if currentSide == config.CELL_BLACK {
		enemySide = config.CELL_WHITE
	} else {
		enemySide = config.CELL_BLACK
	}

	// ひっくり返す可能性がある方向があれば false を返す
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}

			canTurnOver := false
			px, py := x+dx, y+dy
			i := 1
			for {
				// 盤外なら終了
				if px < 0 || px >= 8 || py < 0 || py >= 8 {
					break
				}

				// 空または1マス目で自分の駒なら終了
				if board[py][px] == config.CELL_EMPTY || (board[py][px] == currentSide && i == 1) {
					break
				}

				// 相手の駒なら次へ
				if board[py][px] == enemySide {
					px += dx
					py += dy
					i++
					continue
				}

				// 自分の駒であれば、ひっくり返しが発生する
				if board[py][px] == currentSide {
					canTurnOver = true
					break
				}
			}
			if canTurnOver {
				return false
			}
		}
	}

	return true
}
