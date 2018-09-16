//
// Display termbox 256 colors.
// Usage:
//     go run termbox-256colors.go
//
// Press Esc to quit.
//

package main

import (
	"strconv"
	"time"

	termbox "github.com/nsf/termbox-go"
)

const textColor = termbox.ColorWhite
const backgroundColor = termbox.ColorBlack

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.SetOutputMode(termbox.Output256)
	termbox.Clear(backgroundColor, backgroundColor)

	for i := 1; i < 256; i++ {
		z := i / 100
		y := (i % 100) / 10
		x := (i % 100) % 10 / 1

		Render256(i, x, y, z)
	}

	termbox.Flush()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		ev := <-eventQueue
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Key == termbox.KeyEsc:
				return
			}
		}
		time.Sleep(1 * time.Second)
	}

}

func Render256(i, x, y, z int) {
	row := i % 16
	col := i / 16

	color := termbox.Attribute(z*100 + y*10 + x)
	termbox.SetCell(col*3, row, []rune(strconv.Itoa(z))[0], textColor, color)
	termbox.SetCell(col*3+1, row, []rune(strconv.Itoa(y))[0], textColor, color)
	termbox.SetCell(col*3+2, row, []rune(strconv.Itoa(x))[0], textColor, color)
}
