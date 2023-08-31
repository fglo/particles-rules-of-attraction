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

	for _, o := range options.opts {
		o(cb)
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
	rows := cb.height
	cols := cb.width * 4
	arr := make([]byte, rows*cols)

	lastRowId := rows - 1
	penultimateRowId := lastRowId - 1
	lastColId := cols - 4
	penultimateColId := lastColId - 4

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j += 4 {
			if i == 0 || i == lastRowId || j == 0 || j == lastColId || (cb.Checked && j > 4 && j < penultimateColId && i > 1 && i < penultimateRowId) {
				arr[j+cols*i] = 230
				arr[j+1+cols*i] = 230
				arr[j+2+cols*i] = 230
				arr[j+3+cols*i] = 255
			} else if !cb.Checked && j > 4 && j < 32 && i > 1 && i < 8 {
				arr[j+cols*i] = 9
				arr[j+1+cols*i] = 32
				arr[j+2+cols*i] = 42
				arr[j+3+cols*i] = 255
			}
		}
	}

	cb.image.WritePixels(arr)

	return cb.image
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
