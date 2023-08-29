package simulation

import (
	"math"
	"math/rand"
	"sync"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/color"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/common"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
)

// SimulationEngine encapsulates simulation logic
type SimulationEngine struct {
	particleGroups []*particle.ParticleGroup

	rules [][]float32

	maxEffectDistance          float32
	terminalVelocity           float32
	particleDisplacementFactor float32
	particleRepulsionFactor    float32
	mouseGravityForce          float32
	mouseGravityEffectDistance float32

	numberOfParticles int
	symmetricRules    bool
	wrapped           bool

	Paused    bool
	Forwarded bool
	// reversed  bool
}

// New is a SimulationEngine constructor
func New(numberOfParticles int, symmetricRules bool, maxEffectDistance, terminalVelocity, conservationOfMomentum, particleRepulsionFactor, mouseGravityForce, mouseGravityEffectDistance float32, wrapped bool) *SimulationEngine {
	se := &SimulationEngine{
		numberOfParticles:          numberOfParticles,
		symmetricRules:             symmetricRules,
		particleDisplacementFactor: conservationOfMomentum,
		maxEffectDistance:          maxEffectDistance,
		terminalVelocity:           terminalVelocity,
		particleRepulsionFactor:    particleRepulsionFactor,
		mouseGravityForce:          mouseGravityForce,
		mouseGravityEffectDistance: mouseGravityEffectDistance,
		wrapped:                    wrapped,
	}

	se.Setup()

	return se
}

// Setup prepares particles and rules
func (se *SimulationEngine) Setup() {
	se.Paused = false

	se.particleGroups = []*particle.ParticleGroup{
		particle.NewGroup("red", color.RED, se.numberOfParticles),
		particle.NewGroup("green", color.GREEN, se.numberOfParticles),
		particle.NewGroup("blue", color.BLUE, se.numberOfParticles),
		particle.NewGroup("yellow", color.YELLOW, se.numberOfParticles),
		particle.NewGroup("white", color.WHITE, se.numberOfParticles),
		particle.NewGroup("teal", color.TEAL, se.numberOfParticles),
	}

	if se.symmetricRules {
		se.rules = GenerateRandomSymmetricRules(len(se.particleGroups))
	} else {
		se.rules = GenerateRandomAsymmetricRules(len(se.particleGroups))
	}
}

// Clear removes all particles
func (se *SimulationEngine) Clear() error {
	se.particleGroups = nil
	se.rules = nil
	return nil
}

// TogglePause toggles board pause
func (se *SimulationEngine) TogglePause() {
	se.Paused = !se.Paused
}

// ToggleWrapped toggles wrapped board
func (se *SimulationEngine) ToggleWrapped() {
	se.wrapped = !se.wrapped
}

// Forward sets forward
func (se *SimulationEngine) Forward(forward bool) {
	se.Forwarded = forward
}

// Reset places particles back on initial positions
func (se *SimulationEngine) Reset() {
	for _, pg := range se.particleGroups {
		pg.ResetPosition()
	}
}

func (se *SimulationEngine) NextFrame(mouse input.Mouse) []*particle.ParticleGroup {
	se.Update(mouse)
	return se.particleGroups
}

// Update performs simulation update
func (se *SimulationEngine) Update(mouse input.Mouse) {
	var rulesWg sync.WaitGroup
	rulesWg.Add(len(se.particleGroups))

	for pgIndex := range se.particleGroups {
		go func(pgIndex int) {
			defer rulesWg.Done()
			se.applyRule(pgIndex, mouse)
		}(pgIndex)
	}

	rulesWg.Wait()
}

func (se *SimulationEngine) applyRule(pg1Index int, mouse input.Mouse) {
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

					dx := se.calcDisplacement(p1.X, p2.X)
					dy := se.calcDisplacement(p1.Y, p2.Y)

					if dx != 0 || dy != 0 {
						dSquared := dx*dx + dy*dy
						if dSquared < se.maxEffectDistance {
							F := g * common.FastInvSqrt32(dSquared)
							fx += F * dx
							fy += F * dy
						}
					} else {
						repultion := rand.Float32() * se.particleRepulsionFactor

						if fx > 0 {
							fx += repultion
						} else if fx < 0 {
							fx -= repultion
						}

						if fy > 0 {
							fy += repultion
						} else if fy < 0 {
							fy -= repultion
						}
					}
				}

				chanfx <- fx
				chanfy <- fy
			}(pg1Index, pg2Index, pl, chanfx, chanfy)
		}

		switch {
		case mouse.LeftButtonPressed:
			dx := float32(p1.X - mouse.CursorPosXNormalized)
			dy := float32(p1.Y - mouse.CursorPosYNormalized)

			if dx != 0 || dy != 0 {
				dSquared := dx*dx + dy*dy
				if dSquared < se.mouseGravityEffectDistance {
					F := -se.mouseGravityForce * common.FastInvSqrt32(dSquared)
					fx += F * dx
					fy += F * dy
				}
			}
		case mouse.RightButtonPressed:
			dx := float32(p1.X - mouse.CursorPosXNormalized)
			dy := float32(p1.Y - mouse.CursorPosYNormalized)

			if dx != 0 || dy != 0 {
				dSquared := dx*dx + dy*dy
				if dSquared < se.mouseGravityEffectDistance {
					F := se.mouseGravityForce * common.FastInvSqrt32(dSquared)
					fx += F * dx
					fy += F * dy
				}
			}
		}

		for range se.particleGroups {
			fx += <-chanfx
			fy += <-chanfy
		}

		p1.X, p1.Vx = se.moveParticle(p1.X, p1.Vx, fx)
		p1.Y, p1.Vy = se.moveParticle(p1.Y, p1.Vy, fy)
	}
}

func (se *SimulationEngine) calcDisplacement(coord1, coord2 float32) float32 {
	d := coord1 - coord2
	if se.wrapped {
		dx2 := coord1 + (1 - coord2)
		if math.Abs(float64(d)) > math.Abs(float64(dx2)) {
			d = dx2
		}
	}

	return d
}

func (se *SimulationEngine) moveParticle(coord, velocity, force float32) (float32, float32) {
	velocity = (velocity + force) * se.particleDisplacementFactor

	if velocity > se.terminalVelocity {
		velocity = se.terminalVelocity
	} else if velocity < -se.terminalVelocity {
		velocity = -se.terminalVelocity
	}

	if velocity != 0 {
		coord += velocity
		if coord < 0 {
			if se.wrapped {
				coord += 1
			} else {
				velocity *= -1
				coord = 0
			}
		} else if coord > 1 {
			if se.wrapped {
				coord -= 1
			} else {
				velocity *= -1
				coord = 1
			}
		}
	}

	return coord, velocity
}
