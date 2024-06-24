package game

import (
	"time"
)

type Bird struct {
	posX      int32
	posY      int32
	estimateY float32
	deltaY    float32

	time time.Duration
}

func (b *Bird) initBird() {
	b.posX = 1
	b.posY = 2
	b.estimateY = float32(2)
	b.deltaY = 0.1
	b.time = time.Second

}

func (b *Bird) update() {
	b.estimateY += b.deltaY
	if b.deltaY < 0.30 {
		b.deltaY += 0.01
	}

	b.posY = int32(b.estimateY)

}

func (b *Bird) flap() {
	b.deltaY = -0.30

}
