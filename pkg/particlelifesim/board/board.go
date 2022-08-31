package board

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/rule"
)

type Board struct {
	// image *ebiten.Image

	Particles []*particle.Particle
	Width     int
	Height    int
}

func New() *Board {
	b := new(Board)

	b.Width = 400
	b.Height = 400
	b.Particles = make([]*particle.Particle, 0)

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
	for i := 0; i < numberOfParticles; i++ {
		b.Particles = append(b.Particles, &particle.Particle{Name: name, X: b.randomX(), Y: b.randomY(), Color: color})
	}
}

func (b *Board) applyRules() error {
	for i, p1 := range b.Particles {
		fx, fy := 0.0, 0.0
		for j, p2 := range b.Particles {
			if i == j {
				continue
			}
			g := rule.RULES[p1.Name][p2.Name]
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
		p1.Vx = (p1.Vx + fx) * 0.5
		p1.Vy = (p1.Vy + fy) * 0.5

		p1.X += int(p1.Vx)
		p1.Y += int(p1.Vy)

		if p1.X <= 0 {
			p1.Vx *= -1
			p1.X = 0
		}
		if p1.X >= b.Width {
			p1.Vx *= -1
			p1.X = b.Width
		}
		// Y - axis
		if p1.Y <= 0 {
			p1.Vy *= -1
			p1.Y = 0
		}
		if p1.Y >= b.Height {
			p1.Vy *= -1
			p1.Y = b.Height
		}
	}
	return nil
}

func (b *Board) Update(screen *ebiten.Image) error {
	return nil
}

func (b *Board) Draw(screen *ebiten.Image) {
	screen.Clear()
	screen.Fill(color.RGBA{9, 32, 42, 100})
	b.applyRules()
	for _, p := range b.Particles {
		screen.Set(p.X, p.Y, p.Color)
	}
}

func (b *Board) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return b.Width, b.Height
}
