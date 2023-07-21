package particle

type Particle struct {
	Y  int
	X  int
	Vx float32
	Vy float32
}

func New(x, y int) *Particle {
	return &Particle{
		X: x,
		Y: y,
	}
}
