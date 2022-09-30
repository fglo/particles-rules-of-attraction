package rule

import "math/rand"

type Rule map[string]float64

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
