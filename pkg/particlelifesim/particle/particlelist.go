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
	return &ParticleList{
		Name:      name,
		Color:     color,
		Particles: make([]*Particle, 0),
	}
}
