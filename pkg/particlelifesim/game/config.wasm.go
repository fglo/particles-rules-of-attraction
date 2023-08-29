//go:build wasm
// +build wasm

package game

const (
	screenWidth       = 250
	screenHeight      = 250
	numberOfParticles = 400

	wrapped        = true
	symmetricRules = true

	maxEffectDistance          float32 = .04
	terminalVelocity           float32 = .06
	particleDisplacementFactor float32 = .0003
	particleRepulsionFactor    float32 = .07

	mouseGravityForce          float32 = 258.
	mouseGravityEffectDistance float32 = .1
)
