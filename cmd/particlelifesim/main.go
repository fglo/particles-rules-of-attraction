package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/board"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func run() {
	numberOfParticles := 1000

	b := board.New()
	b.CreateParticles("red", numberOfParticles, particle.RED)
	b.CreateParticles("green", numberOfParticles, particle.GREEN)
	b.CreateParticles("blue", numberOfParticles, particle.BLUE)
	b.CreateParticles("yellow", numberOfParticles, particle.YELLOW)
	b.CreateParticles("white", numberOfParticles, particle.WHITE)
	b.CreateParticles("teal", numberOfParticles, particle.TEAL)

	b.Init()
	if err := ebiten.RunGame(b); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
