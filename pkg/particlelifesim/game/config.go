//go:build !wasm
// +build !wasm

package game

const (
	screenWidth       = 400
	screenHeight      = 400
	numberOfParticles = 1000

	wrapped        = false
	symmetricRules = false

	maxEffectDistance          float32 = .04
	terminalVelocity           float32 = .06
	particleDisplacementFactor float32 = .0003
	particleRepulsionFactor    float32 = .07

	mouseGravityForce          float32 = 258.
	mouseGravityEffectDistance float32 = .1
)
