package gui

type Event struct {
	idCounter uint32
	handlers  []handler
}

// A HandlerFunc is a function that receives and handles an event. When firing an event using
// Event.Fire, arbitrary event arguments may be passed that are in turn passed on to the handler function.
type HandlerFunc func(args interface{})

// RemoveHandlerFunc is a function that removes a handler from an event.
type RemoveHandlerFunc func()

type handler struct {
	id uint32
	h  HandlerFunc
}

// A DeferredAction is an action that is executed at a later time.
type DeferredAction interface {
	// Do executes the action.
	Do()
}

type deferredEvent struct {
	event *Event
	args  interface{}
}

// AddHandler registers event handler h with e. It returns a function to remove h from e if desired.
func (e *Event) AddHandler(h HandlerFunc) RemoveHandlerFunc {
	e.idCounter++

	id := e.idCounter

	e.handlers = append(e.handlers, handler{
		id: id,
		h:  h,
	})

	return func() {
		e.removeHandler(id)
	}
}

func (e *Event) removeHandler(id uint32) {
	index := -1
	for i, h := range e.handlers {
		if h.id == id {
			index = i
			break
		}
	}

	if index < 0 {
		return
	}

	e.handlers = append(e.handlers[:index], e.handlers[index+1:]...)
}

// Fire fires an event to all registered handlers. Arbitrary event arguments may be passed
// which are in turn passed on to event handlers.
//
// Events are not fired directly, but are put into a deferred queue. This queue is then
// processed by the UI.
func (e *Event) Fire(args interface{}) {
	AddDeferred(&deferredEvent{
		event: e,
		args:  args,
	})
}

func (e *Event) handle(args interface{}) {
	for _, h := range e.handlers {
		h.h(args)
	}
}

// Do implements DeferredAction.
func (e *deferredEvent) Do() {
	e.event.handle(e.args)
}

// AddEventHandlerOneShot registers event handler h with e. When e fires an event, h is removed from e immediately.
func AddEventHandlerOneShot(e *Event, h HandlerFunc) {
	var r RemoveHandlerFunc
	rh := func(args interface{}) {
		r()
		h(args)
	}
	r = e.AddHandler(rh)
}

var deferredActions []DeferredAction

// AddDeferred adds d to the queue of deferred actions.
func AddDeferred(d *deferredEvent) {
	deferredActions = append(deferredActions, d)
}

// ExecuteDeferred processes the queue of deferred actions and executes them.
func ExecuteDeferred() {
	defer func(d []DeferredAction) {
		deferredActions = d[:0]
	}(deferredActions)

	for len(deferredActions) > 0 {
		a := deferredActions[0]
		deferredActions = deferredActions[1:]

		a.Do()
	}
}
