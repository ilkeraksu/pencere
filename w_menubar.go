package pencere

import "fmt"

func NewMenuBar() *Pencere {
	p := NewPencere()
	p.HasBorder = false
	p.CanFocus = true
	p.Render = func(buf *Buffer) error {
		//Decorate(p, buf)
		buf.SetString(0, 0, fmt.Sprintf("%v", p.Data), p.Fg, p.Bg)
		return nil
	}
	p.Height = 1

	// p.Layout = func() error {

	// 	p.Left = 1
	// 	p.Height = 3

	// 	p.Top = p.Parent().Height - p.Height - 1
	// 	p.Width = p.Parent().Width - 2

	// 	return nil
	// }

	return p
}
