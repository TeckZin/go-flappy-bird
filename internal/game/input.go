package game

import (
	"github.com/eiannone/keyboard"
)

func (g *Game) listenToInput() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	defer keyboard.Close()

	for {
		_, key, err := keyboard.GetKey()

		if err != nil {
			panic(err)
		}

		switch key {
		case keyboard.KeyArrowUp:
			g.bird.flap()
		case keyboard.KeySpace:
			g.bird.flap()

		case keyboard.KeyEsc:
			return

		}

	}

}
