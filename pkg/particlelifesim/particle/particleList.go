package particle

import (
	"image/color"
)

type ParticleList struct {
	Name      string
	Color     color.Color
	Particles []*Particle
}

func NewList(name string, color color.Color) *ParticleList {
	pl := new(ParticleList)
	pl.Name = name
	pl.Color = color
	pl.Particles = make([]*Particle, 0)

	return pl
}
