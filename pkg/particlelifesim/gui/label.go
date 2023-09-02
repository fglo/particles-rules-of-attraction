package gui

import (
	"image"
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

const (
	fontRegular = "pkg/particlelifesim/gui/fonts/Minecraftia-Regular.ttf"
	// fontRegular = "pkg/particlelifesim/gui/fonts/rainyhearts.ttf"
	// fontRegular = "../../pkg/particlelifesim/gui/fonts/rainyhearts.ttf"
)

type position int

const (
	left     position = 0
	centered position = 1
	right    position = 2
)

type Label struct {
	widget
	text  string
	color color.RGBA
	font  font.Face

	position position
	textPosX int
	textPosY int

	textBounds image.Rectangle

	Inverted bool
}

type LabelOpt func(b *Label)
type LabelOptions struct {
	opts []LabelOpt
}

func NewLabel(posX, posY float64, labelText string, color color.RGBA, options *LabelOptions) *Label {
	fontFace, _ := loadFont(fontRegular, 8)

	bounds := text.BoundString(fontFace, labelText)

	width := bounds.Dx()
	height := bounds.Dy()

	lbl := &Label{
		text:       labelText,
		color:      color,
		font:       fontFace,
		position:   left,
		textPosX:   0,
		textPosY:   -bounds.Min.Y,
		textBounds: bounds,
		Inverted:   false,
	}

	if options != nil {
		for _, o := range options.opts {
			o(lbl)
		}
	}

	lbl.widget = lbl.createWidget(posX, posY, width, height)

	return lbl
}

func (o *LabelOptions) Centered(cx, cy int) *LabelOptions {
	o.opts = append(o.opts, func(l *Label) {
		l.position = centered
		l.textPosX, l.textPosY = cx-l.textBounds.Min.X-l.textBounds.Dx()/2, cy-l.textBounds.Min.Y-l.textBounds.Dy()/2
	})

	return o
}

func (o *LabelOptions) CenteredHorizontally(cx int) *LabelOptions {
	o.opts = append(o.opts, func(l *Label) {
		l.position = centered
		l.textPosX = cx - l.textBounds.Min.X - l.textBounds.Dx()/2
	})

	return o
}

func (o *LabelOptions) CenteredVertically(cy int) *LabelOptions {
	o.opts = append(o.opts, func(l *Label) {
		l.position = centered
		l.textPosY = cy - l.textBounds.Min.Y - l.textBounds.Dy()/2
	})

	return o
}

func (o *LabelOptions) Left(lx int) *LabelOptions {
	o.opts = append(o.opts, func(l *Label) {
		l.position = left
		l.textPosX = lx
	})

	return o
}

func (o *LabelOptions) Right(rx int) *LabelOptions {
	o.opts = append(o.opts, func(l *Label) {
		l.position = right
		l.textPosX = rx - l.textBounds.Dx()/2
	})

	return o
}

func (l *Label) Invert() {
	l.Inverted = !l.Inverted
}

func (l *Label) Draw() *ebiten.Image {
	if l.Inverted {
		text.Draw(l.image, l.text, l.font, 0, l.textPosY, color.RGBA{255 - l.color.R, 255 - l.color.G, 255 - l.color.B, l.color.A})
	} else {
		text.Draw(l.image, l.text, l.font, 0, l.textPosY, l.color)
	}

	return l.image
}

func (l *Label) DrawCentered(image *ebiten.Image, cx, cy int) {
	bounds := text.BoundString(l.font, l.text)
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2

	if l.Inverted {
		text.Draw(image, l.text, l.font, x, y, color.RGBA{255 - l.color.R, 255 - l.color.G, 255 - l.color.B, l.color.A})
	} else {
		text.Draw(image, l.text, l.font, x, y, l.color)
	}
}

func (l *Label) createWidget(posX, posY float64, width, height int) widget {
	widgetOptions := &WidgetOptions{}

	return *NewWidget(posX, posY, width, height, widgetOptions)
}
