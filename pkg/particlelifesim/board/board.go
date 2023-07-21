package board

import (
	image "image/color"
	"math/rand"
	"sync"

	"golang.org/x/exp/slices"

	ebiten "github.com/hajimehoshi/ebiten/v2"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/color"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/common"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/rule"
)

// Board encapsulates simulation logic
type Board struct {
	particlesByName map[string]*particle.ParticleList
	particleNames   []string

	width  int
	height int

	rulesOfAttraction map[string]rule.Rule

	paused    bool
	forwarded bool
	// reversed  bool
}

// New is a Board constructor
func New(w, h int) *Board {
	b := &Board{
		width:           w,
		height:          h,
		particlesByName: make(map[string]*particle.ParticleList),
		particleNames:   make([]string, 0),
	}

	return b
}

func (b *Board) randomX() int {
	return rand.Intn(b.width-50) + 25
}

func (b *Board) randomY() int {
	return rand.Intn(b.height-50) + 25
}

func (b *Board) createParticles(name string, numberOfParticles int, color image.Color) {
	if !slices.Contains(b.particleNames, name) {
		b.particleNames = append(b.particleNames, name)
	}

	b.particlesByName[name] = particle.NewList(name, color)
	for i := 0; i < numberOfParticles; i++ {
		p := particle.New(b.randomX(), b.randomY())
		b.particlesByName[name].Particles = append(b.particlesByName[name].Particles, p)
	}
}

// Setup prepares board
func (b *Board) Setup(numberOfParticles int) {
	b.createParticles("red", numberOfParticles, color.RED)
	b.createParticles("green", numberOfParticles, color.GREEN)
	b.createParticles("blue", numberOfParticles, color.BLUE)
	b.createParticles("yellow", numberOfParticles, color.YELLOW)
	b.createParticles("white", numberOfParticles, color.WHITE)
	b.createParticles("teal", numberOfParticles, color.TEAL)

	b.rulesOfAttraction = rule.GenerateRandomRules(b.particleNames)
	b.paused = false
}

// TogglePause toggles board pause
func (b *Board) TogglePause() {
	b.paused = !b.paused
}

// Forward sets forward
func (b *Board) Forward(forward bool) {
	b.forwarded = forward
}

// Update performs board updates
func (b *Board) Update() error {
	return nil
}

// Size returns board size
func (b *Board) Size() (w, h int) {
	return b.width, b.height
}

// Draw draws board
func (b *Board) Draw(boardImage *ebiten.Image) {
	b.drawParticles(boardImage)
}

func (b *Board) drawParticles(boardImage *ebiten.Image) {
	if !b.paused || b.forwarded {
		boardImage.Clear()
		b.applyRules()
		for _, pl := range b.particlesByName {
			for _, p := range pl.Particles {
				boardImage.Set(p.X, p.Y, pl.Color)
			}
		}
	}
}

func (b *Board) applyRules() {
	var rulesWg sync.WaitGroup
	rulesWg.Add(len(b.particleNames))

	for _, name := range b.particleNames {
		go func(name string) {
			defer rulesWg.Done()
			b.applyRule(name)
		}(name)
	}

	rulesWg.Wait()
}

func (b *Board) applyRule(p1Name string) {
	leftMouseIsPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	rightMouseIsPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	cursorPosX, cursorPosY := ebiten.CursorPosition()

	for i1, p1 := range b.particlesByName[p1Name].Particles {
		var fx, fy float32 = 0.0, 0.0

		chanfx := make(chan float32)
		chanfy := make(chan float32)

		for p2Name, pl := range b.particlesByName {
			go func(p1Name, p2Name string, pl *particle.ParticleList, chanfx, chanfy chan float32) {
				var fx, fy float32 = 0.0, 0.0
				g := b.rulesOfAttraction[p1Name][p2Name]

				for i2, p2 := range pl.Particles {
					if i1 == i2 && p1Name == p2Name {
						continue
					}

					dx := float32(p1.X - p2.X)
					dy := float32(p1.Y - p2.Y)

					if dx != 0 || dy != 0 {
						d := dx*dx + dy*dy
						if d < 6400 {
							F := g * common.FastInvSqrt32(d)
							fx += F * dx
							fy += F * dy
						}
					}
				}

				chanfx <- fx
				chanfy <- fy
			}(p1Name, p2Name, pl, chanfx, chanfy)
		}

		switch {
		case leftMouseIsPressed:
			var g float32 = -32.

			dx := float32(p1.X - cursorPosX)
			dy := float32(p1.Y - cursorPosY)

			if dx != 0 || dy != 0 {
				d := dx*dx + dy*dy
				F := g * common.FastInvSqrt32(d)
				fx += F * dx
				fy += F * dy
			}
		case rightMouseIsPressed:
			var g float32 = 32.

			dx := float32(p1.X - cursorPosX)
			dy := float32(p1.Y - cursorPosY)

			if dx != 0 || dy != 0 {
				d := dx*dx + dy*dy
				F := g * common.FastInvSqrt32(d)
				fx += F * dx
				fy += F * dy
			}
		}

		for range b.particlesByName {
			fx += <-chanfx
			fy += <-chanfy
		}

		var factor float32 = .2

		p1.Vx = (p1.Vx + fx) * factor
		if p1.Vx >= 1 || p1.Vx <= -1 {
			p1.X += int(p1.Vx)
			if p1.X <= 0 {
				p1.Vx *= -1
				p1.X = 0
			} else if p1.X >= b.width {
				p1.Vx *= -1
				p1.X = b.width - 1
			}
		}

		p1.Vy = (p1.Vy + fy) * factor
		if p1.Vy >= 1 || p1.Vy <= -1 {
			p1.Y += int(p1.Vy)
			if p1.Y <= 0 {
				p1.Vy *= -1
				p1.Y = 0
			} else if p1.Y >= b.height {
				p1.Vy *= -1
				p1.Y = b.height - 1
			}
		}
	}
}
