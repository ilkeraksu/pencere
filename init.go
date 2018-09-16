package pencere

import (
	"time"

	tb "github.com/nsf/termbox-go"
)

// Init initializes termui library. This function should be called before any others.
// After initialization, the library must be finalized by 'Close' function.
func Init() error {
	if err := tb.Init(); err != nil {
		return err
	}
	tb.SetInputMode(tb.InputEsc | tb.InputMouse)
	tb.SetOutputMode(tb.Output256)
	//tb.SetOutputMode(tb.OutputNormal)
	//tb.Clear(tb.ColorBlack, tb.ColorBlack)
	//tb.SetOutputMode(tb.OutputNormal)

	// Body = NewGrid()
	// Body.Width, Body.Height = tb.Size()

	root.Width, root.Height = tb.Size()

	timer1 := time.NewTimer(100 * time.Millisecond)
	go func() {
		<-timer1.C
		if isRenderDirty {
			Render()
		}
	}()
	return nil
}

// Close finalizes termui library.
// It should be called after successful initialization when termui's functionality isn't required anymore.
func Close() {
	tb.Close()
}
