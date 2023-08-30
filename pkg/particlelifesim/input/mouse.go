package input

import "github.com/hajimehoshi/ebiten/v2"

type Mouse struct {
	LeftButtonPressed      bool
	LeftButtonJustPressed  bool
	RightButtonPressed     bool
	RightButtonJustPressed bool
	CursorPosX             int
	CursorPosY             int
	CursorPosXNormalized   float32
	CursorPosYNormalized   float32
}

func (m *Mouse) Update() {
	cursorPosX, cursorPosY := ebiten.CursorPosition()

	m.LeftButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	m.RightButtonPressed = ebiten.IsMouseButtonPressed(ebiten.MouseButtonRight)
	m.CursorPosX = cursorPosX
	m.CursorPosY = cursorPosY
}
