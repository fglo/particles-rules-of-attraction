package simulation

import (
	"sync"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/common"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
)

// SimulationEngine encapsulates simulation logic
type SimulationEngine struct {
	particleGroups []*particle.ParticleGroup
	particleNames  []string

	rules [][]float32

	maxEffectDistance       float32
	terminalVelocity        float32
	conservationOfMomentum  float32
	particleRepulsionFactor float32

	wrapped bool

	paused    bool
	forwarded bool
	// reversed  bool
}

// New is a SimulationEngine constructor
func New(w, h int, particleGroups []*particle.ParticleGroup) *SimulationEngine {
	se := &SimulationEngine{
		particleGroups:         particleGroups,
		particleNames:          make([]string, 0),
		conservationOfMomentum: .2,
		maxEffectDistance:      6400,
	}

	return se
}

// Setup prepares particles and rules
func (se *SimulationEngine) Setup() {
	se.rules = GenerateRandomAsymmetricRules(len(se.particleGroups))
	se.paused = false
}

// Clear removes all particles
func (se *SimulationEngine) Clear() error {
	se.particleGroups = nil
	se.rules = nil
	return nil
}

// TogglePause toggles board pause
func (se *SimulationEngine) TogglePause() {
	se.paused = !se.paused
}

// Forward sets forward
func (se *SimulationEngine) Forward(forward bool) {
	se.forwarded = forward
}

// Reset places particles back on initial positions
func (se *SimulationEngine) Reset() {
	for _, pg := range se.particleGroups {
		pg.ResetPosition()
	}
}

// Update performs simulation update
func (se *SimulationEngine) Update() {
	var rulesWg sync.WaitGroup
	rulesWg.Add(len(se.particleNames))

	for pgIndex := range se.particleNames {
		go func(pgIndex int) {
			defer rulesWg.Done()
			se.applyRule(pgIndex)
		}(pgIndex)
	}

	rulesWg.Wait()
}

func (se *SimulationEngine) applyRule(pg1Index int) {
	for i1, p1 := range se.particleGroups[pg1Index].Particles {
		var fx, fy float32 = 0.0, 0.0

		chanfx := make(chan float32)
		chanfy := make(chan float32)

		for pg2Index, pl := range se.particleGroups {
			go func(pg1Index, pg2Index int, pl *particle.ParticleGroup, chanfx, chanfy chan float32) {
				var fx, fy float32 = 0.0, 0.0
				g := se.rules[pg1Index][pg2Index]

				for i2, p2 := range pl.Particles {
					if i1 == i2 && pg1Index == pg2Index {
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
			}(pg1Index, pg2Index, pl, chanfx, chanfy)
		}

		// switch {
		// case leftMouseIsPressed:
		// 	var g float32 = -32.

		// 	dx := float32(p1.X - cursorPosX)
		// 	dy := float32(p1.Y - cursorPosY)

		// 	if dx != 0 || dy != 0 {
		// 		d := dx*dx + dy*dy
		// 		F := g * common.FastInvSqrt32(d)
		// 		fx += F * dx
		// 		fy += F * dy
		// 	}
		// case rightMouseIsPressed:
		// 	var g float32 = 32.

		// 	dx := float32(p1.X - cursorPosX)
		// 	dy := float32(p1.Y - cursorPosY)

		// 	if dx != 0 || dy != 0 {
		// 		d := dx*dx + dy*dy
		// 		F := g * common.FastInvSqrt32(d)
		// 		fx += F * dx
		// 		fy += F * dy
		// 	}
		// }

		for range se.particleGroups {
			fx += <-chanfx
			fy += <-chanfy
		}

		var factor float32 = se.conservationOfMomentum

		p1.Vx = (p1.Vx + fx) * factor
		if p1.Vy != 0 {
			p1X := p1.X + p1.Vx
			if p1X <= 0 {
				if se.wrapped {
					p1X += 1
				} else {
					p1.Vx *= -1
					p1X = 0
				}
			} else if p1X >= 1 {
				if se.wrapped {
					p1X -= 1
				} else {
					p1.Vx *= -1
					p1X = 1
				}
			}
			p1.X = p1X
		}

		p1.Vy = (p1.Vy + fy) * factor
		// if math.Abs(float64(p1.Vy)) > float64(se.terminalVelocity) {
		// 	negativeY := math.Signbit(float64(p1.Vy))
		// 	p1.Vy = se.terminalVelocity
		// 	if negativeY {
		// 		p1.Vy *= -1
		// 	}
		// }

		if p1.Vy != 0 {
			p1Y := p1.Y + p1.Vy
			if p1Y <= 0 {
				if se.wrapped {
					p1Y += 1
				} else {
					p1.Vy *= -1
					p1Y = 0
				}
			} else if p1Y >= 1 {
				if se.wrapped {
					p1Y -= 1
				} else {
					p1.Vy *= -1
					p1Y = 1
				}
			}
			p1.Y = p1Y
		}
	}
}
