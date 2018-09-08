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
		return
	}

	if e.Type == tb.EventKey {
		if focusedPencere != nil {

			event := KeyEvent{
				Key:  e.Key,
				Rune: e.Ch,
				Mod:  e.Mod,
			}
			if focusedPencere.OnKeyEvent != nil {
				focusedPencere.OnKeyEvent(event)
			}
		}

		Render()
		return
	}

	if e.Type == tb.EventResize {

		root.Width = e.Width
		root.Height = e.Height

		tb.Clear(0, 0)
		tb.Flush()
		//panic(fmt.Sprintf("RESIZE %v %v", e.Width, e.Height))
		layoutPencere(root)

		Render()

		return
	}

	if e.Type == tb.EventMouse {
		switch e.Key {
		case tb.MouseLeft:

			p, x, y := getPencereAt(e.MouseX, e.MouseY)

			event := MouseEvent{
				Target:  p,
				GlobalX: e.MouseX,
				GlobalY: e.MouseY,
				X:       x,
				Y:       y,
			}

			cur := p
			for {
				if cur == nil {
					break
				}

				if cur.OnMouseLeftClick != nil {
					cur.OnMouseLeftClick(event)
					break
				} else {
					cur = cur.parent
				}
			}

			cur = p
			for {
				if cur == nil {
					break
				}

				if cur.CanFocus {

					SetFocus(cur)
					break
				} else {
					cur = cur.parent
				}
			}

			Render()
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

// Modifier is a mask of modifier keys.
type Modifier int16

type Key uint16

type KeyEvent struct {
	Target *Pencere
	Mod    tb.Modifier // one of Mod* constants or 0
	Key    tb.Key      // one of Key* constants, invalid if 'Ch' is not 0
	Rune   rune
}
