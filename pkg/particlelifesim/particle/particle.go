package particle

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
