package rule

import "math/rand"

type Rule map[string]float64

func GetRules() map[string]Rule {

	return map[string]Rule{
		"green": {
			"green":  rand.Float64()*2 - 1,
			"red":    rand.Float64()*2 - 1,
			"yellow": rand.Float64()*2 - 1,
			"blue":   rand.Float64()*2 - 1,
		},
		"red": {
			"green":  rand.Float64()*2 - 1,
			"red":    rand.Float64()*2 - 1,
			"yellow": rand.Float64()*2 - 1,
			"blue":   rand.Float64()*2 - 1,
		},
		"yellow": {
			"green":  rand.Float64()*2 - 1,
			"red":    rand.Float64()*2 - 1,
			"yellow": rand.Float64()*2 - 1,
			"blue":   rand.Float64()*2 - 1,
		},
		"blue": {
			"green":  rand.Float64()*2 - 1,
			"red":    rand.Float64()*2 - 1,
			"yellow": rand.Float64()*2 - 1,
			"blue":   rand.Float64()*2 - 1,
		},
	}
}
