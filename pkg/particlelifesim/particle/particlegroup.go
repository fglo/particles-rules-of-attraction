package particle

import (
	"image/color"
	"math/rand"
)

type ParticleGroup struct {
	Name             string
	Color            color.Color
	Particles        []*Particle
	initialPositions []Particle
}

func NewGroup(name string, color color.Color, numberOfParticles int) *ParticleGroup {
	group := &ParticleGroup{
		Name:             name,
		Color:            color,
		Particles:        createParticles(numberOfParticles),
		initialPositions: make([]Particle, 0),
	}

	for _, p := range group.Particles {
		group.initialPositions = append(group.initialPositions, *p)
	}

	return group
}

func (pg *ParticleGroup) ResetPosition() {
	for i, p := range pg.initialPositions {
		pg.Particles[i] = p.Clone()
	}
}

func createParticles(numberOfParticles int) []*Particle {
	particles := make([]*Particle, 0)

	for i := 0; i < numberOfParticles; i++ {
		p := New(rand.Float32(), rand.Float32())
		particles = append(particles, p)
	}

	return particles
}
