package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/game"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func run() {
	g := game.New()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
