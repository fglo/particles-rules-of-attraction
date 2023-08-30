package gui

import (
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Component interface {
	Draw() *ebiten.Image
	Toggle()
	Position() (float64, float64)
	Size() (int, int)
}

type Gui struct {
	components []Component

	leftButtonIsPressed bool
}

func New() *Gui {
	return &Gui{
		components: make([]Component, 0),
	}
}

func (gui *Gui) AddComponent(component Component) {
	gui.components = append(gui.components, component)
}

func (gui *Gui) Update(mouse input.Mouse) {
	if mouse.LeftButtonPressed && !gui.leftButtonIsPressed {
		for _, component := range gui.components {
			posX, posY := component.Position()
			w, h := component.Size()
			if mouse.CursorPosX > int(posX) && mouse.CursorPosX < int(posX)+w && mouse.CursorPosY > int(posY) && mouse.CursorPosY < int(posY)+h {
				component.Toggle()
			}
		}
	}

	gui.leftButtonIsPressed = mouse.LeftButtonPressed
}

func (gui *Gui) Draw(guiImage *ebiten.Image) {
	for _, component := range gui.components {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(component.Position())
		guiImage.DrawImage(component.Draw(), op)
	}
}

type CheckBox struct {
	checked bool
	toggled bool
	image   *ebiten.Image
	width   int
	height  int
	posX    float64
	posY    float64
}

func NewCheckBox(posX, posY float64) *CheckBox {
	return &CheckBox{
		checked: false,
		image:   ebiten.NewImage(10, 10),
		width:   10,
		height:  10,
		posX:    posX,
		posY:    posY,
	}
}

func (cb *CheckBox) Set(checked bool) {
	cb.checked = checked
}

func (cb *CheckBox) Toggle() {
	cb.checked = !cb.checked
}

func (cb *CheckBox) Draw() *ebiten.Image {
	arr := make([]byte, 400)

	for i := 0; i < 10; i++ {
		for j := 0; j < 40; j += 4 {
			if i == 0 || i == 9 || j == 0 || j == 36 || (cb.checked && j > 4 && j < 32 && i > 1 && i < 8) {
				arr[j+40*i] = 230
				arr[j+1+40*i] = 230
				arr[j+2+40*i] = 230
				arr[j+3+40*i] = 255
			} else if !cb.checked && j > 4 && j < 32 && i > 1 && i < 8 {
				arr[j+40*i] = 9
				arr[j+1+40*i] = 32
				arr[j+2+40*i] = 42
				arr[j+3+40*i] = 255
			}
		}
	}

	cb.image.WritePixels(arr)

	return cb.image
}

func (cb *CheckBox) Position() (float64, float64) {
	return cb.posX, cb.posY
}

func (cb *CheckBox) Size() (int, int) {
	return cb.width, cb.height
}
