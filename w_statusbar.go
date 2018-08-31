package pencere

import (
	"fmt"
)

func NewStatusBar() *Pencere {
	p := NewPencere()
	p.HasBorder = true
	p.Render = func(buf *Buffer) error {
		Decorate(p, buf)
		buf.SetString(1, 1, fmt.Sprintf("%v", p.Data), p.Fg, p.Bg)
		return nil
	}

	p.Layout = func() error {

		p.Left = 1
		p.Height = 3

		p.Top = p.Parent().Height - p.Height - 1
		p.Width = p.Parent().Width - 3

		return nil
	}

	return p
}
