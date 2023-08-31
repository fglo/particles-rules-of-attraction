package gui

import (
	"image"

	"github.com/fglo/particles-rules-of-attraction/pkg/particlelifesim/input"
	ebiten "github.com/hajimehoshi/ebiten/v2"
)

type Widget interface {
	Draw() *ebiten.Image
	Position() (float64, float64)
	Size() (int, int)
	FireEvents(mouse *input.Mouse)
	Disable()
	Enable()
}

type widget struct {
	image *ebiten.Image

	Rect image.Rectangle

	disabled bool

	width  int
	height int

	posX float64
	posY float64

	lastUpdateMouseLeftButtonPressed  bool
	lastUpdateMouseRightButtonPressed bool
	lastUpdateCursorEntered           bool

	MouseButtonPressedEvent  Event
	MouseButtonReleasedEvent Event
	CursorEnterEvent         Event
	CursorExitEvent          Event
}

type WidgetOpt func(w *widget)
type WidgetOptions struct {
	opts []WidgetOpt
}

func NewWidget(posX, posY float64, width, height int, options *WidgetOptions) *widget {
	w := &widget{
		image:  ebiten.NewImage(width, height),
		width:  width,
		height: height,
		posX:   posX,
		posY:   posY,
		Rect:   image.Rectangle{Min: image.Point{int(posX), int(posY)}, Max: image.Point{int(posX) + width, int(posY) + height}},
	}
	for _, o := range options.opts {
		o(w)
	}

	return w
}

func (o *WidgetOptions) Disabled() *WidgetOptions {
	o.opts = append(o.opts, func(w *widget) {
		w.disabled = true
	})

	return o
}

func (w *widget) Disable() {
	w.disabled = true
}

func (w *widget) Enable() {
	w.disabled = false
}

func (w *widget) Position() (float64, float64) {
	return w.posX, w.posY
}

func (w *widget) Size() (int, int) {
	return w.width, w.height
}

func (w *widget) FireEvents(mouse *input.Mouse) {
	// posX, posY := w.Position()
	// width, height := w.Size()
	// mouseEntered := mouse.CursorPosX > int(posX) && mouse.CursorPosX < int(posX)+width && mouse.CursorPosY > int(posY) && mouse.CursorPosY < int(posY)+height
	p := image.Point{mouse.CursorPosX, mouse.CursorPosY}
	mouseEntered := p.In(w.Rect)

	if mouseEntered {
		w.lastUpdateCursorEntered = true

		if mouse.LeftButtonJustPressed {
			w.lastUpdateMouseLeftButtonPressed = true
			w.MouseButtonPressedEvent.Fire(&WidgetMouseButtonPressedEventArgs{
				Widget: w,
			})
		} else {
			w.CursorEnterEvent.Fire(&WidgetCursorEnterEventArgs{
				Widget: w,
			})
		}
	} else {
		w.lastUpdateCursorEntered = false
		w.CursorExitEvent.Fire(&WidgetCursorExitEventArgs{
			Widget: w,
		})
	}

	if !mouse.LeftButtonPressed && w.lastUpdateMouseLeftButtonPressed {
		w.lastUpdateMouseLeftButtonPressed = false
		w.MouseButtonReleasedEvent.Fire(&WidgetMouseButtonReleasedEventArgs{
			Widget: w,
			Inside: mouseEntered,
		})
	}
}

// WidgetMouseButtonPressedHandlerFunc is a function that handles mouse button press events.
type WidgetMouseButtonPressedHandlerFunc func(args *WidgetMouseButtonPressedEventArgs) //nolint:golint
// WidgetMouseButtonPressedEventArgs are the arguments for mouse button press events.
type WidgetMouseButtonPressedEventArgs struct { //nolint:golint
	Widget *widget
	Button ebiten.MouseButton
}

func (o *WidgetOptions) MouseButtonPressedHandler(f WidgetMouseButtonPressedHandlerFunc) *WidgetOptions {
	o.opts = append(o.opts, func(w *widget) {
		w.MouseButtonPressedEvent.AddHandler(func(args interface{}) {
			f(args.(*WidgetMouseButtonPressedEventArgs))
		})
	})

	return o
}

// WidgetMouseButtonReleasedHandlerFunc is a function that handles mouse button release events.
type WidgetMouseButtonReleasedHandlerFunc func(args *WidgetMouseButtonReleasedEventArgs) //nolint:golint
// WidgetMouseButtonReleasedEventArgs are the arguments for mouse button release events.
type WidgetMouseButtonReleasedEventArgs struct { //nolint:golint
	Widget *widget
	Button ebiten.MouseButton
	Inside bool
}

func (o *WidgetOptions) MouseButtonReleasedHandler(f WidgetMouseButtonReleasedHandlerFunc) *WidgetOptions {
	o.opts = append(o.opts, func(w *widget) {
		w.MouseButtonReleasedEvent.AddHandler(func(args interface{}) {
			f(args.(*WidgetMouseButtonReleasedEventArgs))
		})
	})

	return o
}

// WidgetCursorEnterHandlerFunc is a function that handles cursor enter events.
type WidgetCursorEnterHandlerFunc func(args *WidgetCursorEnterEventArgs) //nolint:golint
// WidgetCursorEnterEventArgs are the arguments for cursor enter events.
type WidgetCursorEnterEventArgs struct { //nolint:golint
	Widget *widget
}

func (o *WidgetOptions) CursorEnterHandler(f WidgetCursorEnterHandlerFunc) *WidgetOptions {
	o.opts = append(o.opts, func(w *widget) {
		w.CursorEnterEvent.AddHandler(func(args interface{}) {
			f(args.(*WidgetCursorEnterEventArgs))
		})
	})

	return o
}

// WidgetCursorExitHandlerFunc is a function that handles cursor exit events.
type WidgetCursorExitHandlerFunc func(args *WidgetCursorExitEventArgs) //nolint:golint
// WidgetCursorExitEventArgs are the arguments for cursor exit events.
type WidgetCursorExitEventArgs struct { //nolint:golint
	Widget *widget
}

func (o *WidgetOptions) CursorExitHandler(f WidgetCursorExitHandlerFunc) *WidgetOptions {
	o.opts = append(o.opts, func(w *widget) {
		w.CursorExitEvent.AddHandler(func(args interface{}) {
			f(args.(*WidgetCursorExitEventArgs))
		})
	})

	return o
}
