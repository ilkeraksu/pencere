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
var mouseDragOverWindow *Pencere
var mouseDragContext *DragContext

func isDropable(p *Pencere, event DropEvent) bool {
	if p.CanDrop == nil || p.OnDrop == nil {
		return false
	}

	canDrop, err := p.CanDrop(event)
	_ = err
	return canDrop
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

		case tb.MouseRelease:
			if mouseDragging {
				// we were dragging when mouserelase

				// First we need to call DragLeave on previous dragover window
				if mouseDragOverWindow != nil && mouseDragOverWindow.OnDragLeave != nil {
					event := DragLeaveEvent{
						Target:      mouseDraggingWindow,
						DragContext: globalDragContext,
						GlobalX:     e.MouseX,
						GlobalY:     e.MouseY,
					}

					mouseDragOverWindow.OnDragLeave(event)
				}

				// no longer dragging
				mouseDragOverWindow = nil

				var dropedPencere *Pencere
				p, x, y := getPencereAt(e.MouseX, e.MouseY)
				// check if we hit a window on mouse release
				if p != nil {
					// we have a window on releaase point

					// send OnDrag
					event := DropEvent{
						DragContext: globalDragContext,
						Target:      mouseDraggingWindow,
						GlobalX:     e.MouseX,
						GlobalY:     e.MouseY,
						X:           x,
						Y:           y,
					}
					if isDropable(p, event) {
						p.OnDrop(event)
						dropedPencere = p
					}

				} else { // No dragging nothing to do

				}

				if mouseDraggingWindow != nil {
					if mouseDraggingWindow.OnDragEnd != nil {
						event := DragEndEvent{
							DragContext:   globalDragContext,
							DropedPencere: dropedPencere,
							Target:        mouseDraggingWindow,
							GlobalX:       e.MouseX,
							GlobalY:       e.MouseY,
						}
						mouseDraggingWindow.OnDragEnd(event)

					}
				}
				globalDragContext = nil
			}

			mouseDragging = false
			mouseReleaseWaiting = false

			Render()
			return

		case tb.MouseLeft:
			if mouseDragging {
				if mouseDraggingWindow != nil {
					if mouseDraggingWindow.OnDragging != nil {
						event := DraggingEvent{
							Target:  mouseDraggingWindow,
							GlobalX: e.MouseX,
							GlobalY: e.MouseY,
						}
						mouseDraggingWindow.OnDragging(event)

					}

					if globalDragContext != nil && globalDragContext.DraggingIcon != nil {
						globalDragContext.DraggingIcon.Left, globalDragContext.DraggingIcon.Top = e.MouseX, e.MouseY
					}
				}

				p, x, y := getPencereAt(e.MouseX, e.MouseY)

				// we are on some window
				if p != nil {

					// check we are already dragging over
					if mouseDragOverWindow == nil {
						dropEvent := DropEvent{
							DragContext: globalDragContext,
							Target:      mouseDraggingWindow,
							GlobalX:     e.MouseX,
							GlobalY:     e.MouseY,
							X:           x,
							Y:           y,
						}

						if isDropable(p, dropEvent) {
							// we are not already dragging over so need to call OnDragEnter
							mouseDragOverWindow = p
							if mouseDragOverWindow.OnDragEnter != nil {
								event := DragEnterEvent{
									Target:      mouseDraggingWindow,
									DragContext: globalDragContext,
									GlobalX:     e.MouseX,
									GlobalY:     e.MouseY,
									X:           x,
									Y:           y,
								}
								mouseDragOverWindow.OnDragEnter(event)
							}
						}

					} else { // we are already dragging over

						// check if it is same window we already dragging over
						if mouseDragOverWindow == p {
							// yes same window so call OnDragOver
							if mouseDragOverWindow.OnDragOver != nil {
								event := DragOverEvent{
									Target:  mouseDraggingWindow,
									GlobalX: e.MouseX,
									GlobalY: e.MouseY,
									X:       x,
									Y:       y,
								}
								mouseDragOverWindow.OnDragOver(event)
							}
						} else { // this is new window we are dragging over

							// we need to leave first
							if mouseDragOverWindow.OnDragLeave != nil {
								event := DragLeaveEvent{
									Target:      mouseDraggingWindow,
									DragContext: globalDragContext,
									GlobalX:     e.MouseX,
									GlobalY:     e.MouseY,
								}

								mouseDragOverWindow.OnDragLeave(event)
							}
							mouseDragOverWindow = p

							dropEvent := DropEvent{
								DragContext: globalDragContext,
								Target:      mouseDraggingWindow,
								GlobalX:     e.MouseX,
								GlobalY:     e.MouseY,
								X:           x,
								Y:           y,
							}

							if isDropable(p, dropEvent) {
								// and lets DragEnter
								if mouseDragOverWindow.OnDragEnter != nil {
									event := DragEnterEvent{
										Target:      mouseDraggingWindow,
										DragContext: globalDragContext,
										GlobalX:     e.MouseX,
										GlobalY:     e.MouseY,
										X:           x,
										Y:           y,
									}
									mouseDragOverWindow.OnDragEnter(event)
								}
							}
						}
					}

				} else { // we are not on a window
					// check if we were on a window and dragover
					// if so we need to call OnDragLeave
					if mouseDragOverWindow != nil && mouseDragOverWindow.OnDragLeave != nil {
						event := DragLeaveEvent{
							Target:      mouseDraggingWindow,
							DragContext: globalDragContext,
							GlobalX:     e.MouseX,
							GlobalY:     e.MouseY,
						}

						mouseDragOverWindow.OnDragLeave(event)
						mouseDragOverWindow = nil
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
					event := DragBeginEvent{
						Target:  p,
						GlobalX: e.MouseX,
						GlobalY: e.MouseY,
						X:       x,
						Y:       y,
					}
					drag, dragContext, err := p.OnDragBegin(event)
					panicif(err)
					if drag && dragContext != nil {
						if dragContext.DraggingIcon != nil {
							dragContext.DraggingIcon.Left = event.GlobalX
							dragContext.DraggingIcon.Top = event.GlobalY
						}
						globalDragContext = dragContext
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

		}
	}
}

type MouseEvent struct {
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}

type DragBeginEvent struct {
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}

type DraggingEvent struct {
	DragContext      *DragContext
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}
type DragEndEvent struct {
	DragContext      *DragContext
	DropedPencere    *Pencere
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}
type DragEnterEvent struct {
	DragContext      *DragContext
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}
type DragLeaveEvent struct {
	DragContext      *DragContext
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}
type DragOverEvent struct {
	DragContext      *DragContext
	Target           *Pencere
	GlobalX, GlobalY int
	X, Y             int
	Type             int
}
type DropEvent struct {
	DragContext      *DragContext
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
