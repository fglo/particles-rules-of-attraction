package rule

import "math/rand"

type Rule map[string]float64

var RULES = map[string]Rule{
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

func GenerateRandomRules(names []string) map[string]Rule {
	rules := make(map[string]Rule)

	for _, name := range names {
		rules[name] = make(Rule)
		for _, name2 := range names {
			rules[name][name2] = rand.Float64()*2 - 1
		}
	}

	return rules
}
