package main

import (
	"log"
	
	"github.com/hajimehoshi/ebiten/v2"
	"ebitengine-othello/src/config"
)

var turn = config.CELL_BLACK

func main() {
	ebiten.SetWindowSize(config.WINDOW_WIDTH, config.WINDOW_HEIGHT)
	ebiten.SetWindowTitle("Hello, World!")
	game := &Game{
		Board: [8][8]int{
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1, 2, 0, 0, 0},
			{0, 0, 0, 2, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0},
		},
	}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
