package pencere

import (
	tb "github.com/nsf/termbox-go"
)

// Loop gets events from termbox and passes them off to handleEvent.
// Stops when StopLoop is called.
func Loop() {
	root.Width, root.Height = tb.Size()

	layoutPencere(root)
	render()
	go func() {
		for {
			eventStream.eventQueue <- tb.PollEvent()
		}
	}()

	for {
		select {
		case <-eventStream.stopLoop:
			return

		case _ = <-eventStream.shouldRender:
			render()
		case e := <-eventStream.eventQueue:
			eventStream.HandleEvent(e)
		}
	}
}

// StopLoop stops the event loop.
func StopLoop() {
	eventStream.stopLoop <- true
}
