package game

import (
	"errors"
	"fmt"
	imgColor "image/color"
	"math/rand"
	"strings"
	"time"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/board"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/simulation"
	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var (
	Terminated = errors.New("terminated")
)

// Game encapsulates game logic
type Game struct {
	board      *board.Board
	boardImage *ebiten.Image
	guiImage   *ebiten.Image

	mouse *input.Mouse

	screenWidth  int
	screenHeight int

	quitIsPressed    bool
	restartIsPressed bool
	resetIsPressed   bool
	forwardIsPressed bool
	debugIsToggled   bool
}

const guiWidth = 60

// New generates a new Game object.
func New() *Game {
	g := &Game{
		screenWidth:  screenWidth + guiWidth,
		screenHeight: screenHeight,
		mouse:        input.NewMouse(),
	}

	simEngine := simulation.New(numberOfParticles, symmetricRules, maxEffectDistance, terminalVelocity, particleDisplacementFactor, particleRepulsionFactor, mouseGravityForce, mouseGravityEffectDistance, wrapped)
	board := board.New(screenWidth, screenHeight, simEngine)
	g.board = board

	ebiten.SetWindowSize(g.getWindowSize())
	ebiten.SetWindowTitle("Particles' Rules of Attraction")

	return g
}

func (g *Game) getWindowSize() (int, int) {
	var factor float32 = 1.8
	return int(float32(g.screenWidth) * factor), int(float32(g.screenHeight) * factor)
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) restart() {
	g.boardImage.Clear()
	g.board.Restart()
}

func (g *Game) reset() {
	g.boardImage.Clear()
	g.board.Reset()
}

// Update updates the current game state.
func (g *Game) Update() error {
	g.checkRestartButton()
	g.checkResetButton()
	g.checkPauseButton()
	g.checkWrappedButton()
	g.checkForwardButton()
	g.checkDebugButton()

	g.mouse.Update()

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

func (g *Game) checkResetButton() {
	if !g.resetIsPressed && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		g.resetIsPressed = true
	}
	if g.resetIsPressed && inpututil.IsKeyJustReleased(ebiten.KeyS) {
		g.resetIsPressed = false
		g.reset()
	}
}

func (g *Game) checkPauseButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyP) || inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.board.TogglePause()
	}
}

func (g *Game) checkForwardButton() {
	if !g.forwardIsPressed && (inpututil.IsKeyJustPressed(ebiten.KeyF) || inpututil.IsKeyJustPressed(ebiten.KeyArrowRight)) {
		g.forwardIsPressed = true
		g.board.Forward(true)
	}
	if g.forwardIsPressed && (inpututil.IsKeyJustReleased(ebiten.KeyF) || inpututil.IsKeyJustReleased(ebiten.KeyArrowRight)) {
		g.forwardIsPressed = false
		g.board.Forward(false)
	}
}

func (g *Game) checkWrappedButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyW) {
		g.board.ToggleWrapped()
	}
}

func (g *Game) checkDebugButton() {
	if inpututil.IsKeyJustPressed(ebiten.KeyD) {
		g.debugIsToggled = !g.debugIsToggled
	}
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	g.mouse.Draw()

	if g.boardImage == nil {
		w, h := g.board.Size()
		g.boardImage = ebiten.NewImage(w, h)
	}

	if g.guiImage == nil {
		_, h := g.board.Size()
		g.guiImage = ebiten.NewImage(guiWidth, h)
	}

	screen.Fill(imgColor.RGBA{9, 32, 42, 255})
	g.drawInstructions(screen)

	g.board.Draw(screen, g.boardImage, g.guiImage, g.debugIsToggled, g.mouse)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(guiWidth, 0)
	screen.DrawImage(g.boardImage, op)

	op = &ebiten.DrawImageOptions{}
	screen.DrawImage(g.guiImage, op)
}

func (g *Game) drawInstructions(screen *ebiten.Image) {
	instructions := []string{
		" P: pause/unpause",
		" F: play paused sim",
		" R: restart",
		" S: reset",
		" W: wrap board",
		" Q: quit",
	}

	if g.debugIsToggled {
		instructions = append([]string{fmt.Sprintf(" TPS: %0.2f\n FPS: %0.2f", ebiten.ActualTPS(), ebiten.ActualFPS())}, instructions...)
	}

	ebitenutil.DebugPrint(screen, strings.Join(instructions, "\n"))
}
