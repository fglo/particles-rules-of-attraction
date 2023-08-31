package gui

import (
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Button struct {
	widget

	pressed  bool
	hovering bool

	PressedEvent  *Event
	ReleasedEvent *Event
	ClickedEvent  *Event
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

	for _, o := range options.opts {
		o(b)
	}

	b.widget = b.createWidget(posX, posY, 45, 15)

	return b
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
	cols := b.width * 4
	rows := b.height
	arr := make([]byte, rows*cols)

	lastRowId := rows - 1
	lastColId := cols - 4

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j += 4 {
			if i == 0 || i == lastRowId || j == 0 || j == lastColId {
				arr[j+cols*i] = 230
				arr[j+1+cols*i] = 230
				arr[j+2+cols*i] = 230
			} else if b.pressed {
				arr[j+cols*i] = 200
				arr[j+1+cols*i] = 200
				arr[j+2+cols*i] = 200
			} else if b.hovering {
				arr[j+cols*i] = 220
				arr[j+1+cols*i] = 220
				arr[j+2+cols*i] = 220
			} else {
				arr[j+cols*i] = 240
				arr[j+1+cols*i] = 240
				arr[j+2+cols*i] = 240
			}
			arr[j+3+cols*i] = 255
		}
	}

	b.image.WritePixels(arr)

	return b.image
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
		if !b.disabled {
			b.pressed = true
			b.PressedEvent.Fire(&ButtonPressedEventArgs{
				Button: b,
			})
		}
	})

	widgetOptions.MouseButtonReleasedHandler(func(args *WidgetMouseButtonReleasedEventArgs) {
		if !b.disabled {
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
