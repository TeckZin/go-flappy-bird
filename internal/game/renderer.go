package game

import "fmt"

func (g *Game) render() {
	var i, j int32

	out := make([][]byte, 0)

	for i = 0; i < g.height; i++ {
		row := make([]byte, 0)
		for j = 0; j < g.width; j++ {
			if i == 0 || i == g.height-1 {
				row = append(row, byte('-'))

			} else {

				row = append(row, byte(' '))
			}

		}
		out = append(out, row)

	}

	out[g.bird.posY][g.bird.posX] = byte('B')

	for _, pipe := range g.pipes {
		out[pipe.topPipeY][pipe.posX] = byte('T')
		out[pipe.bottomPipeY][pipe.posX] = byte('D')

		for i := pipe.topPipeY - 1; i > 0; i-- {
			out[i][pipe.posX] = byte('|')
		}

		for i := pipe.bottomPipeY + 1; i < g.height-1; i++ {
			out[i][pipe.posX] = byte('|')
		}
	}

	g.pixels = out

}

func (g *Game) display() {

	fmt.Print("\033[2J")
	var out string
	for _, x := range g.pixels {
		out = ""

		for _, y := range x {
			out += string(y)
		}
		fmt.Println(out)

	}

}
