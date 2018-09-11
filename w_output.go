package pencere

import (
	"strings"
)

func NewOutput(output chan string, options ...Option) (*Pencere, error) {
	p, err := NewPencere(options...)
	if err != nil {
		return nil, err
	}
	p.HasBorder = true
	p.CanFocus = true
	p.Scrollable = true
	p.Render = func(buf *Buffer) error {

		lines := strings.Split(p.Text, "\n")

		for i, line := range lines {
			buf.SetString(p.Inner.Min.X, p.Inner.Min.Y+i, line, p.Fg, p.Bg)
		}

		// for i := 0 ; i< p.Inner.Dy();i++{

		// }
		p.ContentY = len(lines)
		return nil
	}

	go func() {
		for {
			select {
			case s := <-output:
				p.Text = p.Text + "\n" + s
			}

		}
	}()

	return p, nil
}
