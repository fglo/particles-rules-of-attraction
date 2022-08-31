package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/fglo/particlelifesim/pkg/particlelifesim/board"
	"github.com/fglo/particlelifesim/pkg/particlelifesim/particle"
	"github.com/hajimehoshi/ebiten"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func run() {
	b := board.New()
	b.CreateParticles("red", 500, particle.RED)
	b.CreateParticles("green", 500, particle.GREEN)
	b.CreateParticles("blue", 500, particle.BLUE)
	b.CreateParticles("yellow", 500, particle.YELLOW)

	if err := ebiten.RunGame(b); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// var sourceCodePath string
	// var prettyPrint bool
	// flag.StringVar(&sourceCodePath, "f", "", "Source code file path.")
	// flag.BoolVar(&prettyPrint, "p", false, "Pretty print the AST.")
	// flag.Parse()

	run()
}
