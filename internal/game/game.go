package game

import (
	"time"
)

type Game struct {
	height          int32
	width           int32
	bird            *Bird
	generateNewPipe bool
	pipes           []*Pipe
	pixels          [][]byte
	gameEnd         bool
}

func NewGame(height int32, width int32) *Game {

	bird := &Bird{}
	bird.initBird()
	game := &Game{}
	game.height = height
	game.width = width
	game.gameEnd = false

	game.pipes = make([]*Pipe, 0)
	game.bird = bird

	game.pixels = make([][]byte, 0)

	return game

}

func (g *Game) Start(tick int32) {
	duration := time.Second / time.Duration(tick)

	ticker := time.NewTicker(duration)

	defer ticker.Stop()

	go g.listenToInput()

	counter := 1

	for range ticker.C {

		if counter == 80 {
			counter = 0
			g.generateNewPipe = true

		}

		g.update()
		g.render()
		g.display()
		counter++

		if g.gameEnd {
			break
		}
	}

}

func (g *Game) update() {
	g.bird.update()

	for idx, pipe := range g.pipes {
		flag := pipe.update()
		if flag {
			g.pipes = append(g.pipes[:idx], g.pipes[idx+1:]...)
		}
	}

	if g.generateNewPipe {
		g.generateNewPipe = false
		pipe := &Pipe{}
		pipe.initPipe(g.height, g.width)
		g.pipes = append(g.pipes, pipe)

	}

}
