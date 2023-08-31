package gui

import (
	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Gui struct {
	components []Widget
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
	for _, component := range gui.components {
		component.FireEvents(mouse)
	}
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
