package gui

import (
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Gui struct {
	components []Widget

	leftButtonIsPressed bool
}

func New() *Gui {
	return &Gui{
		components: make([]Widget, 0),
	}
}

func (gui *Gui) AddComponent(component Widget) {
	gui.components = append(gui.components, component)
}

func (gui *Gui) Update(mouse *input.Mouse) {
	// if mouse.LeftButtonPressed && !gui.leftButtonIsPressed {
	// 	for _, component := range gui.components {
	// 		posX, posY := component.Position()
	// 		w, h := component.Size()
	// 		if mouse.CursorPosX > int(posX) && mouse.CursorPosX < int(posX)+w && mouse.CursorPosY > int(posY) && mouse.CursorPosY < int(posY)+h {
	// 			component.Toggle()
	// 		}
	// 	}
	// } else if !mouse.LeftButtonPressed && gui.leftButtonIsPressed {
	// 	for _, component := range gui.components {
	// 		posX, posY := component.Position()
	// 		w, h := component.Size()
	// 		if mouse.CursorPosX > int(posX) && mouse.CursorPosX < int(posX)+w && mouse.CursorPosY > int(posY) && mouse.CursorPosY < int(posY)+h {
	// 			component.Toggle()
	// 		}
	// 	}
	// }

	// gui.leftButtonIsPressed = mouse.LeftButtonPressed
}

func (gui *Gui) Draw(guiImage *ebiten.Image, mouse *input.Mouse) {
	ExecuteDeferred()

	for _, component := range gui.components {
		component.FireEvents(mouse)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(component.Position())
		guiImage.DrawImage(component.Draw(), op)
	}
}
