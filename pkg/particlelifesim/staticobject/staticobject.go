package staticobject

type Object struct {
	Y int
	X int

	Width  int
	Height int
}

func New(x, y, w, h int) *Object {
	return &Object{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}
