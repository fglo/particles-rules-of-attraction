package particle

import (
	"image/color"
)

type Particle struct {
	Y  int
	X  int
	Vx float64
	Vy float64
}

func New(x, y int) *Particle {
	p := new(Particle)
	p.X = x
	p.Y = y

	return p
}

var RED = color.RGBA{
	R: 183,
	G: 53,
	B: 41,
	A: 100,
}

var GREEN = color.RGBA{
	R: 129,
	G: 193,
	B: 90,
	A: 100,
}

var BLUE = color.RGBA{
	R: 49,
	G: 95,
	B: 229,
	A: 100,
}

var YELLOW = color.RGBA{
	R: 248,
	G: 220,
	B: 96,
	A: 100,
}

var WHITE = color.RGBA{
	R: 255,
	G: 255,
	B: 255,
	A: 100,
}

var TEAL = color.RGBA{
	R: 0,
	G: 128,
	B: 128,
	A: 100,
}
