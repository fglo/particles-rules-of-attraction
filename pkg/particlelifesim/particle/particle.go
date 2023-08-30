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

func (p *Particle) Clone() *Particle {
	return &Particle{
		X: p.X,
		Y: p.Y,
	}
}
