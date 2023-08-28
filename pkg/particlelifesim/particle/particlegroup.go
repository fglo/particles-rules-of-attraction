package particle

import (
	"image/color"
	"math"
	"math/rand"
)

type ParticleGroup struct {
	Name             string
	Color            color.Color
	Particles        []*Particle
	initialPositions []Particle
}

func NewGroup(name string, color color.Color, initialPositions []Particle) *ParticleGroup {
	return &ParticleGroup{
		Name:             name,
		Color:            color,
		Particles:        make([]*Particle, 0),
		initialPositions: initialPositions,
	}
}

func (pg *ParticleGroup) ResetPosition() {
	pg.Particles = placeParticles(len(pg.Particles), pg.initialPositions)
}

func placeParticles(n int, p []Particle) (ptcs []*Particle) {
	nClusters := len(p)
	if nClusters <= 0 {
		// Place particles randonly if no clusters are passed
		for i := 0; i < n; i++ {
			p := New(rand.Int(), rand.Int())
			ptcs = append(ptcs, p)
		}
	} else {
		// Place particles proportionally in each clusters
		for i := 0; i < n; i++ {
			tIndex := int(math.Mod(float64(i), float64(nClusters)))
			tParticle := p[tIndex]
			p := New(tParticle.X, tParticle.Y)
			ptcs = append(ptcs, p)
		}
	}
	return ptcs
}
