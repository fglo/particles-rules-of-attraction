package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/board"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func run() {
	b := board.New()
	b.Setup()

	if err := ebiten.RunGame(b); err != nil {
		log.Fatal(err)
	}
}

func main() {
	run()
}
