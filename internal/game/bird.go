package game

import (
	"time"
)

type Bird struct {
	posX      int32
	posY      int32
	estimateX float32
	estimateY float32
	deltaY    float32

	time time.Duration
}

func (b *Bird) initBird() {
	b.posX = 1
	b.posY = 1
	b.deltaY = 0.1
	b.time = time.Second

}

func (b *Bird) update(floorY int32) {
	if b.posY >= floorY {
		b.estimateY = float32(floorY - 1)
	} else {
		b.estimateY += b.deltaY
		if b.deltaY < 0.30 {
			b.deltaY += 0.01
		}
	}

	b.posX = int32(b.estimateX)
	b.posY = int32(b.estimateY)

}

func (b *Bird) flap() {
	b.deltaY = -0.30

}
