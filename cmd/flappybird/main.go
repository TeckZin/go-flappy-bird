package main

import (
	"main/internal/game"
)

func main() {
	game := game.NewGame(40, 200)
	game.Start(30)

}
