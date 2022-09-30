package board

import (
	"image/color"
	"math"
	"math/rand"
	"sync"

	"github.com/hajimehoshi/ebiten"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/rule"
)

type Board struct {
	ParticlesByName map[string]*particle.ParticleList
	ParticleNames   []string

	Width  int
	Height int

	Rules map[string]rule.Rule
}

func New() *Board {
	b := new(Board)

	b.Width = 400
	b.Height = 400
	b.ParticlesByName = make(map[string]*particle.ParticleList)
	b.ParticleNames = make([]string, 0)

	ebiten.SetWindowSize(b.Width*2, b.Height*2)
	ebiten.SetWindowTitle("TRDQFGBJKNM")

	return b
}

func (b *Board) randomX() int {
	return rand.Intn(b.Width-50) + 25
}

func (b *Board) randomY() int {
	return rand.Intn(b.Height-50) + 25
}

func (b *Board) CreateParticles(name string, numberOfParticles int, color color.Color) {
	b.ParticleNames = append(b.ParticleNames, name)
	b.ParticlesByName[name] = particle.NewList(name, color)
	for i := 0; i < numberOfParticles; i++ {
		p := particle.New(b.randomX(), b.randomY())
		b.ParticlesByName[name].Particles = append(b.ParticlesByName[name].Particles, p)
	}
}

var rulesWg sync.WaitGroup

func (b *Board) applyRule(p1Name string) error {
	defer rulesWg.Done()

	for i1, p1 := range b.ParticlesByName[p1Name].Particles {
		fx, fy := 0.0, 0.0
		for p2Name, pl := range b.ParticlesByName {
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
			if p1.X >= b.Width {
				p1.Vx *= -1
				p1.X = b.Width
			}
		}

		p1.Vy = (p1.Vy + fy) * factor
		if p1.Vy >= 1 || p1.Vy <= -1 {
			p1.Y += int(p1.Vy)
			if p1.Y <= 0 {
				p1.Vy *= -1
				p1.Y = 0
			}
			if p1.Y >= b.Height {
				p1.Vy *= -1
				p1.Y = b.Height
			}
		}
	}
	return nil
}

func (b *Board) applyRules() {
	rulesWg.Add(len(b.ParticleNames))

	for _, name := range b.ParticleNames {
		go b.applyRule(name)
	}

	rulesWg.Wait()
}

func (b *Board) Init() {
	if b.Rules == nil {
		b.Rules = rule.GenerateRandomRules(b.ParticleNames)
	}
}

func (b *Board) Update(screen *ebiten.Image) error {
	return nil
}

func (b *Board) Draw(screen *ebiten.Image) {
	_ = screen.Clear()
	_ = screen.Fill(color.RGBA{9, 32, 42, 100})

	b.applyRules()

	for _, pl := range b.ParticlesByName {
		for _, p := range pl.Particles {
			screen.Set(p.X, p.Y, pl.Color)
		}
	}
}

func (b *Board) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return b.Width, b.Height
}
