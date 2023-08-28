package board

import (
	image "image/color"
	"math/rand"

	"golang.org/x/exp/slices"

	ebiten "github.com/hajimehoshi/ebiten/v2"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/particle"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/simulation"
)

// Board encapsulates simulation logic
type Board struct {
	width  int
	height int

	paused    bool
	forwarded bool
	// reversed  bool

	particleNames []string

	se *simulation.SimulationEngine
}

// New is a Board constructor
func New(w, h int, se *simulation.SimulationEngine) *Board {
	b := &Board{
		width:  w,
		height: h,
		se:     se,
	}

	return b
}

func (b *Board) randomX() int {
	return rand.Intn(b.width-50) + 25
}

func (b *Board) randomY() int {
	return rand.Intn(b.height-50) + 25
}

func (b *Board) createParticles(name string, numberOfParticles int, color image.Color) {
	if !slices.Contains(b.particleNames, name) {
		b.particleNames = append(b.particleNames, name)
	}

	particleGroup := particle.NewGroup(name, color)

	for i := 0; i < numberOfParticles; i++ {
		p := particle.New(b.randomX(), b.randomY())
		particleGroup.Particles = append(particleGroup.Particles, p)
	}

	b.particleGroups = append(b.particleGroups, particleGroup)
}

// Setup prepares board
func (b *Board) Setup(numberOfParticles int) {
	b.se.Setup(b.particleGroups)
	b.paused = false
}

// TogglePause toggles board pause
func (b *Board) TogglePause() {
	b.paused = !b.paused
}

// Forward sets forward
func (b *Board) Forward(forward bool) {
	b.forwarded = forward
}

// Update performs board updates
func (b *Board) Update() error {
	return nil
}

// Size returns board size
func (b *Board) Size() (w, h int) {
	return b.width, b.height
}

// Draw draws board
func (b *Board) Draw(boardImage *ebiten.Image) {
	b.drawParticles(boardImage)
}

func (b *Board) drawParticles(boardImage *ebiten.Image) {
	// leftMouseIsPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	// rightMouseIsPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	// cursorPosX, cursorPosY := ebiten.CursorPosition()
	if !b.paused || b.forwarded {
		boardImage.Clear()
		b.applyRules()
		for _, pl := range b.particleGroups {
			for _, p := range pl.Particles {
				boardImage.Set(p.X, p.Y, pl.Color)
			}
		}
	}
}
