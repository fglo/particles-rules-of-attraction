package particle

type Particle struct {
	Y  float32
	X  float32
	Vx float32
	Vy float32
}

func New(x, y float32) *Particle {
	return &Particle{
		X: x,
		Y: y,
	}
}
