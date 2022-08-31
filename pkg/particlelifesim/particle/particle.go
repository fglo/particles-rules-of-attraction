package particle

import "image/color"

type Particle struct {
	Name  string
	Y     int
	X     int
	Vx    float64
	Vy    float64
	Color color.Color
}

var RED = color.RGBA{
	R: 255,
	G: 0,
	B: 0,
	A: 100,
}

var GREEN = color.RGBA{
	R: 0,
	G: 255,
	B: 0,
	A: 100,
}

var BLUE = color.RGBA{
	R: 0,
	G: 0,
	B: 255,
	A: 100,
}

var YELLOW = color.RGBA{
	R: 255,
	G: 255,
	B: 0,
	A: 100,
}
