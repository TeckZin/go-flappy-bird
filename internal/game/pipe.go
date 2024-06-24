package game

import (
	"math/rand"
)

type Pipe struct {
	topPipeY    int32
	bottomPipeY int32
	posX        int32
	estimateX   float32
}

func (p *Pipe) initPipe(height int32, width int32) {

	randomInt := rand.Intn(int(height-15)) + 5

	p.topPipeY = int32(randomInt)

	p.estimateX = float32(width) - 1
	p.posX = int32(width) - 1

	p.bottomPipeY = int32(randomInt) + 5

}

func (p *Pipe) update() bool {

	p.estimateX += -0.5

	p.posX = int32(p.estimateX)

	return p.posX <= 0
}
