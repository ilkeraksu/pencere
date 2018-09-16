package pencere

import (
	"time"

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

var mouseClickTime time.Time
var mouseDragging bool
var mouseReleaseWaiting bool

var mouseLastClicked *Pencere
var mouseDraggingWindow *Pencere

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

		case tb.MouseRelease:
			if mouseDragging {
				if mouseDraggingWindow != nil {
					if mouseDraggingWindow.OnDragEnd != nil {
						event := MouseEvent{
							Target:  mouseDraggingWindow,
							GlobalX: e.MouseX,
							GlobalY: e.MouseY,
						}
						mouseDraggingWindow.OnDragEnd(event)

					}
				}

			}

			mouseDragging = false
			mouseReleaseWaiting = false
			Render()
			return

		case tb.MouseLeft:
			if mouseDragging {
				if mouseDraggingWindow != nil {
					if mouseDraggingWindow.OnDragging != nil {
						event := MouseEvent{
							Target:  mouseDraggingWindow,
							GlobalX: e.MouseX,
							GlobalY: e.MouseY,
						}
						mouseDraggingWindow.OnDragging(event)

					}
				}
				Render()
				return
			}

			//click
			//if mouseReleaseWaiting && time.Now().After(mouseClickTime.Add(time.Millisecond*200)) {
			if mouseReleaseWaiting {

				p, x, y := getPencereAt(e.MouseX, e.MouseY)

				_, _ = x, y

				if p.OnDragBegin != nil {
					event := MouseEvent{
						Target:  p,
						GlobalX: e.MouseX,
						GlobalY: e.MouseY,
						X:       x,
						Y:       y,
					}
					drag, err := p.OnDragBegin(event)
					panicif(err)
					if drag {
						mouseDraggingWindow = p
						mouseDragging = true
					}

					Render()
					return
				} else {

				}
			}

			p, x, y := getPencereAt(e.MouseX, e.MouseY)

			mouseLastClicked = p

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

			mouseReleaseWaiting = true

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
