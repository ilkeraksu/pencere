package pencere

import (
	tb "github.com/nsf/termbox-go"
)

var eventStream = EventStream{
	//make(map[string]func(Event)),
	//pre"",
	shouldRender: make(chan bool, 1),
	stopLoop:     make(chan bool, 1),
	eventQueue:   make(chan tb.Event),
}

type EventStream struct {
	//eventHandlers map[string]func(Event)
	//prevKey       string // previous keypress
	shouldRender chan bool
	stopLoop     chan bool
	eventQueue   chan tb.Event // list of events from termbox
}

func (this EventStream) HandleEvent(e tb.Event) {
	if e.Ch == 'q' {
		StopLoop()
	}

	if e.Type == tb.EventResize {

		root.Width = e.Width
		root.Height = e.Height

		//panic(fmt.Sprintf("RESIZE %v %v", e.Width, e.Height))
		Layout()
		Render()
	}

	if e.Type == tb.EventMouse {
		switch e.Key {
		case tb.MouseLeft:
			root.dispatchMouseLeftClick(MouseEvent{
				GlobalX: e.MouseX,
				GlobalY: e.MouseY,
				X:       e.MouseX,
				Y:       e.MouseY,
			})
			//return "<MouseLeft>"
			// case tb.MouseMiddle:
			// 	return "<MouseMiddle>"
			// case tb.MouseRight:
			// 	return "<MouseRight>"
			// case tb.MouseWheelUp:
			// 	return "<MouseWheelUp>"
			// case tb.MouseWheelDown:
			// 	return "<MouseWheelDown>"
			// case tb.MouseRelease:
			// 	return "<MouseRelease>"
		}
	}
}

type MouseEvent struct {
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}
