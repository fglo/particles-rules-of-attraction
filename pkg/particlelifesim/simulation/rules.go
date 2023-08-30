package simulation

// GenerateRules creates 2D array to store particle rules
func GenerateRules(numberOfParticleGroups int) [][]float32 {
	rules := make([][]float32, numberOfParticleGroups)
	for i := range rules {
		rules[i] = make([]float32, numberOfParticleGroups)
	}
	return rules
}

// GenerateRandomSymmetricRules generates random rules for every pair of particles in a symmetric configuration
func GenerateRandomSymmetricRules(numberOfParticleGroups int) [][]float32 {
	rules := GenerateRules(numberOfParticleGroups)

	for i := 0; i < numberOfParticleGroups; i++ {
		for j := 0; j < numberOfParticleGroups; j++ {
			rules[i][j] = randFloat32()
			// if j != i {
			rules[j][i] = rules[i][j]
			// }
		}
	}

	return rules
}

// GenerateRandomAsymmetricRules generates random rules for every pair of particles in a asymmetric configuration
func GenerateRandomAsymmetricRules(numberOfParticleGroups int) [][]float32 {
	rules := GenerateRules(numberOfParticleGroups)

	for i := 0; i < numberOfParticleGroups; i++ {
		for j := 0; j < numberOfParticleGroups; j++ {
			rules[i][j] = randFloat32()
			// if j != i {
			rules[j][i] = randFloat32()
			// }
		}
	}

	return rules
}
