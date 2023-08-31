package simulation

import (
	"math/rand"
)

func randFloat32() float32 {
	return (rand.Float32() - .5) / 2
}
