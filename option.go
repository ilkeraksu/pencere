package pencere

//Option ...
type Option func(*Pencere) error

func LayoutFill() Option {
	return func(p *Pencere) error {

		p.Layout = func() error {
			parent := p.Parent()
			if parent == nil {
				return nil
			}
			p.Left = parent.Inner.Min.X
			p.Top = parent.Inner.Min.Y
			p.Width = parent.Inner.Dx()
			p.Height = parent.Inner.Dy()
			return nil
		}
		return nil
	}

}

func Position(left, top, width, height int) Option {
	return func(p *Pencere) error {
		p.Left = left
		p.Top = top
		p.Width = width
		p.Height = height
		return nil
	}

}

const (
	BORDER = 1 << iota
	SCROLLABLE
	ENABLED
	VISIBLE
	CANFOCUS
)

func SetFlag(flags int) Option {
	return func(p *Pencere) error {
		if flags&BORDER == BORDER {
			p.HasBorder = true
		}

		if flags&SCROLLABLE == SCROLLABLE {
			p.Scrollable = true
		}
		if flags&ENABLED == ENABLED {
			p.Enabled = true
		}

		if flags&VISIBLE == VISIBLE {
			p.Visible = true
		}

		if flags&CANFOCUS == CANFOCUS {
			p.CanFocus = true
		}

		return nil
	}
}

func ResetFlag(flags int) Option {
	return func(p *Pencere) error {
		if flags&BORDER == BORDER {
			p.HasBorder = false
		}

		if flags&SCROLLABLE == SCROLLABLE {
			p.Scrollable = false
		}
		if flags&ENABLED == ENABLED {
			p.Enabled = false
		}

		if flags&VISIBLE == VISIBLE {
			p.Visible = false
		}

		if flags&CANFOCUS == CANFOCUS {
			p.CanFocus = false
		}

		return nil
	}
}

// b3.Layout = func() error {
// 	b3.Left = b2.Inner.Min.X
// 	b3.Top = b2.Inner.Min.Y
// 	b3.Width = b2.Inner.Dx()
// 	b3.Height = b2.Inner.Dy()
// 	return nil
// }
