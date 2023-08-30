package board

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/simulation"
)

// Board encapsulates simulation logic
type Board struct {
	width  int
	height int

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

// TogglePause toggles board pause
func (b *Board) TogglePause() {
	b.se.TogglePause()
}

// ToggleWrapped toggles wrapped board
func (b *Board) ToggleWrapped() {
	b.se.ToggleWrapped()
}

// Forward sets forward
func (b *Board) Forward(forward bool) {
	b.se.Forward(forward)
}

// Setup prepares board
func (b *Board) Setup() {
	b.se.Setup()
}

func (b *Board) Restart() {
	b.se.Setup()
}

func (b *Board) Reset() {
	b.se.Reset()
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
func (b *Board) Draw(boardImage *ebiten.Image, debugIsToggled bool, mouse input.Mouse) {
	b.drawParticles(boardImage, mouse)
}

func (b *Board) drawParticles(boardImage *ebiten.Image, mouse input.Mouse) {
	if !b.se.Paused || b.se.Forwarded {
		mouse.CursorPosXNormalized = float32(mouse.CursorPosX) / float32(b.width)
		mouse.CursorPosYNormalized = float32(mouse.CursorPosY) / float32(b.height)

		boardImage.Clear()
		for _, pg := range b.se.NextFrame(mouse) {
			for _, p := range pg.Particles {
				boardImage.Set(int(p.X*float32(b.width)), int(p.Y*float32(b.height)), pg.Color)
			}
		}
	}
}
