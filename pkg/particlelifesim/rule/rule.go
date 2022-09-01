package rule

import "math/rand"

type Rule map[string]float64

var RULES = map[string]Rule{
	"green": {
		"green":  rand.Float64()*3 - 1.5,
		"red":    rand.Float64()*3 - 1.5,
		"yellow": rand.Float64()*3 - 1.5,
		"blue":   rand.Float64()*3 - 1.5,
	},
	"red": {
		"green":  rand.Float64()*3 - 1.5,
		"red":    rand.Float64()*3 - 1.5,
		"yellow": rand.Float64()*3 - 1.5,
		"blue":   rand.Float64()*3 - 1.5,
	},
	"yellow": {
		"green":  rand.Float64()*3 - 1.5,
		"red":    rand.Float64()*3 - 1.5,
		"yellow": rand.Float64()*3 - 1.5,
		"blue":   rand.Float64()*3 - 1.5,
	},
	"blue": {
		"green":  rand.Float64()*3 - 1.5,
		"red":    rand.Float64()*3 - 1.5,
		"yellow": rand.Float64()*3 - 1.5,
		"blue":   rand.Float64()*3 - 1.5,
	},
}

var factor = 2.0

var RULES_PREDEFINED = map[string]Rule{
	"green": {
		"green":  0.878214014158254 * factor,
		"red":    0.383942932294564 * factor,
		"yellow": 0.3632328353781209 * factor,
		"blue":   0.4357079645785089 * factor,
	},
	"red": {
		"green":  -0.8131279812066854 * factor,
		"red":    0.8761564046567396 * factor,
		"yellow": -0.686246916739194 * factor,
		"blue":   -0.42403398294928163 * factor,
	},
	"yellow": {
		"green":  0.8283611643992606 * factor,
		"red":    -0.8050409003234531 * factor,
		"yellow": 0.8422661062679588 * factor,
		"blue":   -0.6206303204367405 * factor,
	},
	"blue": {
		"green":  -0.6276679142294777 * factor,
		"red":    -0.48726835984229977 * factor,
		"yellow": -0.8155039608681607 * factor,
		"blue":   0.49503848830455155 * factor,
	},
}
