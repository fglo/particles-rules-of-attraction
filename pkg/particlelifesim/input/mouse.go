package input

import ebiten "github.com/hajimehoshi/ebiten/v2"

type Mouse struct {
	CursorPosX           int
	CursorPosY           int
	CursorPosXNormalized float32
	CursorPosYNormalized float32

	LeftButtonPressed     bool
	LeftButtonJustPressed bool
	LastLeftButtonPressed bool

	RightButtonPressed     bool
	RightButtonJustPressed bool
	LastRightButtonPressed bool
}

func NewMouse() *Mouse {
	mouse := &Mouse{}
	mouse.Update()

	return mouse
}

func (m *Mouse) Update() {
	m.CursorPosX, m.CursorPosY = ebiten.CursorPosition()

	m.LeftButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	m.RightButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
}

func (m *Mouse) Draw() {
	m.LeftButtonJustPressed = m.LeftButtonPressed != m.LastLeftButtonPressed

	m.RightButtonJustPressed = m.RightButtonPressed != m.LastRightButtonPressed
}
