package gui

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type CheckBox struct {
	widget
	Checked bool

	ToggledEvent *Event
}

type CheckBoxOpt func(b *CheckBox)
type CheckBoxOptions struct {
	opts []CheckBoxOpt
}

type CheckBoxToggledEventArgs struct {
	CheckBox *CheckBox
}

type CheckBoxToggledHandlerFunc func(args *CheckBoxToggledEventArgs)

func NewCheckBox(posX, posY float64, options *CheckBoxOptions) *CheckBox {
	cb := &CheckBox{
		Checked:      false,
		ToggledEvent: &Event{},
	}

	if options != nil {
		for _, o := range options.opts {
			o(cb)
		}
	}

	cb.widget = cb.createWidget(posX, posY, 10, 10)

	return cb
}

func (o *CheckBoxOptions) ToggledHandler(f CheckBoxToggledHandlerFunc) *CheckBoxOptions {
	o.opts = append(o.opts, func(b *CheckBox) {
		b.ToggledEvent.AddHandler(func(args interface{}) {
			f(args.(*CheckBoxToggledEventArgs))
		})
	})

	return o
}

func (cb *CheckBox) Set(checked bool) {
	cb.Checked = checked
}

func (cb *CheckBox) Toggle() {
	cb.Checked = !cb.Checked
	cb.ToggledEvent.Fire(&CheckBoxToggledEventArgs{
		CheckBox: cb,
	})
}

func (cb *CheckBox) Draw() *ebiten.Image {
	if cb.Checked {
		cb.image.WritePixels(cb.drawChecked())
	} else {
		cb.image.WritePixels(cb.drawUnchecked())
	}

	return cb.image
}

func (cb *CheckBox) drawUnchecked() []byte {
	arr := make([]byte, cb.pixelRows*cb.pixelCols)

	for i := 0; i < cb.pixelRows; i++ {
		for j := 0; j < cb.pixelCols; j += 4 {
			if i == 0 || i == cb.lastPixelRowId || j == 0 || j == cb.lastPixelColId {
				arr[j+cb.pixelCols*i] = 230
				arr[j+1+cb.pixelCols*i] = 230
				arr[j+2+cb.pixelCols*i] = 230
			} else {
				arr[j+cb.pixelCols*i] = 9
				arr[j+1+cb.pixelCols*i] = 32
				arr[j+2+cb.pixelCols*i] = 42
			}
			arr[j+3+cb.pixelCols*i] = 255
		}
	}

	return arr
}

func (cb *CheckBox) drawChecked() []byte {
	arr := make([]byte, cb.pixelRows*cb.pixelCols)

	for i := 0; i < cb.pixelRows; i++ {
		for j := 0; j < cb.pixelCols; j += 4 {
			if i == 0 || i == cb.lastPixelRowId || j == 0 || j == cb.lastPixelColId || (j > 4 && j < cb.penultimatePixelColId && i > 1 && i < cb.penultimatePixelRowId) {
				arr[j+cb.pixelCols*i] = 230
				arr[j+1+cb.pixelCols*i] = 230
				arr[j+2+cb.pixelCols*i] = 230
			}
			arr[j+3+cb.pixelCols*i] = 255
		}
	}

	return arr
}

func (cb *CheckBox) createWidget(posX, posY float64, width, height int) widget {
	widgetOptions := &WidgetOptions{}

	widgetOptions.MouseButtonReleasedHandler(func(args *WidgetMouseButtonReleasedEventArgs) {
		if !cb.disabled && args.Inside {
			cb.Checked = !cb.Checked
			cb.ToggledEvent.Fire(&CheckBoxToggledEventArgs{
				CheckBox: cb,
			})
		}
	})

	return *NewWidget(posX, posY, width, height, widgetOptions)
}
