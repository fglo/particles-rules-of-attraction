package board

import (
	"image/color"
	"math"
	"math/rand"
	"sync"

	"golang.org/x/exp/slices"

	ebiten "github.com/hajimehoshi/ebiten/v2"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/rule"
)

type Board struct {
	particlesByName map[string]*particle.ParticleList
	particleNames   []string

	width  int
	height int

	Rules map[string]rule.Rule
}

func New(w, h int) *Board {
	b := new(Board)

	b.width = w
	b.height = h
	b.particlesByName = make(map[string]*particle.ParticleList)
	b.particleNames = make([]string, 0)

	return b
}

func (b *Board) randomX() int {
	return rand.Intn(b.width-50) + 25
}

func (b *Board) randomY() int {
	return rand.Intn(b.height-50) + 25
}

func (b *Board) createParticles(name string, numberOfParticles int, color color.Color) {
	if !slices.Contains(b.particleNames, name) {
		b.particleNames = append(b.particleNames, name)
	}

	b.particlesByName[name] = particle.NewList(name, color)
	for i := 0; i < numberOfParticles; i++ {
		p := particle.New(b.randomX(), b.randomY())
		b.particlesByName[name].Particles = append(b.particlesByName[name].Particles, p)
	}
}

func (b *Board) Setup() {
	numberOfParticles := 1000

	b.createParticles("red", numberOfParticles, particle.RED)
	b.createParticles("green", numberOfParticles, particle.GREEN)
	b.createParticles("blue", numberOfParticles, particle.BLUE)
	b.createParticles("yellow", numberOfParticles, particle.YELLOW)
	b.createParticles("white", numberOfParticles, particle.WHITE)
	b.createParticles("teal", numberOfParticles, particle.TEAL)

	b.Rules = rule.GenerateRandomRules(b.particleNames)
}

func (b *Board) Update() error {
	return nil
}

func (b *Board) Size() (w, h int) {
	return b.width, b.height
}

func (b *Board) Draw(boardImage *ebiten.Image) {
	b.drawParticles(boardImage)
}

func (b *Board) drawParticles(boardImage *ebiten.Image) {
	boardImage.Clear()

	b.applyRules()

	for _, pl := range b.particlesByName {
		for _, p := range pl.Particles {
			boardImage.Set(p.X, p.Y, pl.Color)
		}
	}
}

var rulesWg sync.WaitGroup

func (b *Board) applyRule(p1Name string) {
	defer rulesWg.Done()

	for i1, p1 := range b.particlesByName[p1Name].Particles {
		fx, fy := 0.0, 0.0
		for p2Name, pl := range b.particlesByName {
			g := b.Rules[p1Name][p2Name]
			for i2, p2 := range pl.Particles {
				if p1Name == p2Name && i1 == i2 {
					continue
				}
				dx := float64(p1.X - p2.X)
				dy := float64(p1.Y - p2.Y)

				if dx != 0 || dy != 0 {
					d := dx*dx + dy*dy
					if d < 6400 {
						F := g / math.Sqrt(d)
						fx += F * dx
						fy += F * dy
					}
				}
			}
		}

		factor := 0.08

		p1.Vx = (p1.Vx + fx) * factor
		if p1.Vx >= 1 || p1.Vx <= -1 {
			p1.X += int(p1.Vx)
			if p1.X <= 0 {
				p1.Vx *= -1
				p1.X = 0
			}
			if p1.X >= b.width {
				p1.Vx *= -1
				p1.X = b.width
			}
		}

		p1.Vy = (p1.Vy + fy) * factor
		if p1.Vy >= 1 || p1.Vy <= -1 {
			p1.Y += int(p1.Vy)
			if p1.Y <= 0 {
				p1.Vy *= -1
				p1.Y = 0
			}
			if p1.Y >= b.height {
				p1.Vy *= -1
				p1.Y = b.height
			}
		}
	}
}

func (b *Board) applyRules() {
	rulesWg.Add(len(b.particleNames))

	for _, name := range b.particleNames {
		go b.applyRule(name)
	}

	rulesWg.Wait()
}
