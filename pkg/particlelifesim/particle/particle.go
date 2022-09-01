package particle

import (
	"image/color"

	"github.com/google/uuid"
)

type Particle struct {
	Id    uuid.UUID
	Name  string
	Y     int
	X     int
	Vx    float64
	Vy    float64
	Color color.Color
}

func New(name string, x, y int, color color.Color) *Particle {
	p := new(Particle)
	p.Id = uuid.New()
	p.Name = name
	p.X = x
	p.Y = y
	p.Color = color

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
