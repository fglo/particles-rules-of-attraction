package rule

import "math/rand"

type Rule map[string]float64

// GenerateRandomRules generates random rules for every pair of particles
func GenerateRandomRules(names []string) map[string]Rule {
	rules := make(map[string]Rule)

	for _, name := range names {
		rules[name] = make(Rule)
	}

	for _, name := range names {
		rules[name] = make(Rule)
		for _, name2 := range names {
			rules[name][name2] = rand.Float64()*2 - 1
			rules[name2][name] = rules[name][name2]
		}
	}

	return rules
}
