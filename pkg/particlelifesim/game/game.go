package game

import (
	"errors"
	"image/color"
	"math/rand"
	"strings"
	"time"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/board"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	ScreenWidth  = 400
	ScreenHeight = 400
)

var (
	Terminated = errors.New("terminated")
)

// Game encapsulates game logic
type Game struct {
	// input      *Input
	board      *board.Board
	boardImage *ebiten.Image

	quitIsPressed    bool
	restartIsPressed bool
}

// New generates a new Game object.
func New() *Game {
	g := new(Game)
	// g.input =  NewInput()
	g.board = board.New(ScreenWidth, ScreenHeight)
	g.board.Setup()

	ebiten.SetWindowSize(ScreenWidth*2, ScreenHeight*2)
	ebiten.SetWindowTitle("Particles' Rules of Attraction")

	return g
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func (g *Game) restart() {
	g.board.Setup()
	g.boardImage.Clear()
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.checkRestartButton()
	// g.input.Update()
	if err := g.board.Update(); err != nil {
		return err
	}
	return g.checkQuitButton()
}

func (g *Game) checkQuitButton() error {
	if !g.quitIsPressed && inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		g.quitIsPressed = true
	}
	if g.quitIsPressed && inpututil.IsKeyJustReleased(ebiten.KeyQ) {
		g.quitIsPressed = false
		return Terminated
	}
	return nil
}

func (g *Game) checkRestartButton() {
	if !g.restartIsPressed && inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.restartIsPressed = true
	}
	if g.restartIsPressed && inpututil.IsKeyJustReleased(ebiten.KeyR) {
		g.restartIsPressed = false
		g.restart()
	}
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage = ebiten.NewImage(w, h)
	}

	screen.Fill(color.RGBA{9, 32, 42, 100})
	g.drawInstructions(screen)

	g.board.Draw(g.boardImage)
	op := &ebiten.DrawImageOptions{}
	sw, sh := screen.Size()
	bw, bh := g.boardImage.Size()
	x := (sw - bw) / 2
	y := (sh - bh) / 2
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.boardImage, op)
}

func (g *Game) drawInstructions(screen *ebiten.Image) {
	instructions := []string{
		" Press R to restart the simulation",
	}
	ebitenutil.DebugPrint(screen, strings.Join(instructions, "\n"))
}
