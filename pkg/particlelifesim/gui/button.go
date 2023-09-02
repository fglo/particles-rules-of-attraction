package gui

import (
	"image/color"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	widget

	pressed  bool
	hovering bool

	PressedEvent  *Event
	ReleasedEvent *Event
	ClickedEvent  *Event

	Label *Label
}

type ButtonOpt func(b *Button)
type ButtonOptions struct {
	opts []ButtonOpt
}

type ButtonPressedEventArgs struct {
	Button *Button
}

type ButtonReleasedEventArgs struct {
	Button *Button
	Inside bool
}

type ButtonClickedEventArgs struct {
	Button *Button
}

type ButtonPressedHandlerFunc func(args *ButtonPressedEventArgs)

type ButtonReleasedHandlerFunc func(args *ButtonReleasedEventArgs)

type ButtonClickedHandlerFunc func(args *ButtonClickedEventArgs)

func NewButton(posX, posY float64, options *ButtonOptions) *Button {
	b := &Button{
		PressedEvent:  &Event{},
		ReleasedEvent: &Event{},
		ClickedEvent:  &Event{},
	}

	if options != nil {
		for _, o := range options.opts {
			o(b)
		}
	}

	b.widget = b.createWidget(posX, posY, 45, 15)

	return b
}

func (o *ButtonOptions) Text(posX, posY float64, text string, color color.RGBA) *ButtonOptions {
	label := NewLabel(posX, 15, text, color, &LabelOptions{})

	o.PressedHandler(func(args *ButtonPressedEventArgs) {
		label.Inverted = true
	})

	o.ReleasedHandler(func(args *ButtonReleasedEventArgs) {
		label.Inverted = false
	})

	o.opts = append(o.opts, func(b *Button) {
		b.Label = label
	})

	return o
}

func (o *ButtonOptions) PressedHandler(f ButtonPressedHandlerFunc) *ButtonOptions {
	o.opts = append(o.opts, func(b *Button) {
		b.PressedEvent.AddHandler(func(args interface{}) {
			f(args.(*ButtonPressedEventArgs))
		})
	})

	return o
}

func (o *ButtonOptions) ReleasedHandler(f ButtonReleasedHandlerFunc) *ButtonOptions {
	o.opts = append(o.opts, func(b *Button) {
		b.ReleasedEvent.AddHandler(func(args interface{}) {
			f(args.(*ButtonReleasedEventArgs))
		})
	})

	return o
}

func (o *ButtonOptions) ClickedHandler(f ButtonClickedHandlerFunc) *ButtonOptions {
	o.opts = append(o.opts, func(b *Button) {
		b.ClickedEvent.AddHandler(func(args interface{}) {
			f(args.(*ButtonClickedEventArgs))
		})
	})

	return o
}

func (b *Button) Draw() *ebiten.Image {
	if b.pressed {
		b.image.WritePixels(b.drawPressed())
	} else if b.hovering {
		b.image.WritePixels(b.drawHovered())
	} else {
		b.image.WritePixels(b.draw())
	}

	if b.Label != nil {
		b.Label.DrawCentered(b.image, (b.Rect.Max.X-b.Rect.Min.X)/2, (b.Rect.Max.Y-b.Rect.Min.Y)/2+1)
		// opts := &ebiten.DrawImageOptions{}
		// opts.GeoM.Translate(float64((b.Rect.Max.X-b.Rect.Min.X)/2), float64((b.Rect.Max.Y-b.Rect.Min.Y)/2+1))
		// b.image.DrawImage(b.Label.Draw(), opts)
	}

	return b.image
}

func (b *Button) draw() []byte {
	arr := make([]byte, b.pixelRows*b.pixelCols)

	for i := 0; i < b.pixelRows; i++ {
		for j := 0; j < b.pixelCols; j += 4 {
			if i == 0 && (j == 0 || j == b.lastPixelColId) || i == b.lastPixelRowId && (j == 0 || j == b.lastPixelColId) {
				continue
			} else if i == 0 || i == b.lastPixelRowId || j == 0 || j == b.lastPixelColId {
				arr[j+b.pixelCols*i] = 230
				arr[j+1+b.pixelCols*i] = 230
				arr[j+2+b.pixelCols*i] = 230
			} else if j > 4 && j < b.penultimatePixelColId && i > 1 && i < b.penultimatePixelRowId {
				arr[j+b.pixelCols*i] = 230
				arr[j+1+b.pixelCols*i] = 230
				arr[j+2+b.pixelCols*i] = 230
			}
			arr[j+3+b.pixelCols*i] = 255
		}
	}

	return arr
}

func (b *Button) drawPressed() []byte {
	arr := make([]byte, b.pixelRows*b.pixelCols)

	for i := 0; i < b.pixelRows; i++ {
		for j := 0; j < b.pixelCols; j += 4 {
			if i == 0 && (j == 0 || j == b.lastPixelColId) || i == b.lastPixelRowId && (j == 0 || j == b.lastPixelColId) {
				continue
			} else if i == 0 || i == b.lastPixelRowId || j == 0 || j == b.lastPixelColId {
				arr[j+b.pixelCols*i] = 200
				arr[j+1+b.pixelCols*i] = 200
				arr[j+2+b.pixelCols*i] = 200
			} else if j > 4 && j < b.penultimatePixelColId && i > 1 && i < b.penultimatePixelRowId {
				arr[j+b.pixelCols*i] = 9
				arr[j+1+b.pixelCols*i] = 32
				arr[j+2+b.pixelCols*i] = 42
			}
			arr[j+3+b.pixelCols*i] = 255
		}
	}

	return arr
}

func (b *Button) drawHovered() []byte {
	arr := make([]byte, b.pixelRows*b.pixelCols)

	for i := 0; i < b.pixelRows; i++ {
		for j := 0; j < b.pixelCols; j += 4 {
			if i == 0 && (j == 0 || j == b.lastPixelColId) || i == b.lastPixelRowId && (j == 0 || j == b.lastPixelColId) {
				continue
			} else if i == 0 || i == b.lastPixelRowId || j == 0 || j == b.lastPixelColId {
				arr[j+b.pixelCols*i] = 220
				arr[j+1+b.pixelCols*i] = 220
				arr[j+2+b.pixelCols*i] = 220
			} else if j > 4 && j < b.penultimatePixelColId && i > 1 && i < b.penultimatePixelRowId {
				arr[j+b.pixelCols*i] = 200
				arr[j+1+b.pixelCols*i] = 200
				arr[j+2+b.pixelCols*i] = 200
			}
			arr[j+3+b.pixelCols*i] = 255
		}
	}

	return arr
}

func (b *Button) createWidget(posX, posY float64, width, height int) widget {
	widgetOptions := &WidgetOptions{}

	widgetOptions.CursorEnterHandler(func(args *WidgetCursorEnterEventArgs) {
		if !b.disabled {
			b.hovering = true
		}
	})

	widgetOptions.CursorExitHandler(func(args *WidgetCursorExitEventArgs) {
		b.hovering = false
	})

	widgetOptions.MouseButtonPressedHandler(func(args *WidgetMouseButtonPressedEventArgs) {
		if !b.disabled && args.Button == ebiten.MouseButtonLeft {
			b.pressed = true
			b.PressedEvent.Fire(&ButtonPressedEventArgs{
				Button: b,
			})
		}
	})

	widgetOptions.MouseButtonReleasedHandler(func(args *WidgetMouseButtonReleasedEventArgs) {
		if !b.disabled && args.Button == ebiten.MouseButtonLeft {
			b.pressed = false
			b.ReleasedEvent.Fire(&ButtonReleasedEventArgs{
				Button: b,
				Inside: args.Inside,
			})

			b.ClickedEvent.Fire(&ButtonClickedEventArgs{
				Button: b,
			})
		}
	})

	return *NewWidget(posX, posY, width, height, widgetOptions)
}
