package pencere

import (
	"fmt"
)

func NewStatusBar() *Pencere {
	p := NewPencere()
	p.HasBorder = true
	p.Render = func(buf *Buffer) error {

		buf.SetString(1, 1, fmt.Sprintf("%v", p.Data), p.Fg, p.Bg)
		return nil
	}

	p.Layout = func() error {
		p.Left = p.Parent().Inner.Min.X
		p.Height = 3
		p.Top = p.Parent().Inner.Max.Y - p.Height
		p.Width = p.Parent().Inner.Dx()

		return nil
	}

	return p
}
