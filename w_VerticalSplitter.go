package pencere

type Splitter struct {
	*Pencere
}

func NewVerticalSplitter(options ...Option) (*Pencere, error) {
	p, err := NewPencere(options...)
	if err != nil {
		return nil, err
	}
	p.HasBorder = false
	p.Width = 1
	p.Render = func(buf *Buffer) error {

		buf.SetCell(0, 0, Cell{Ch: HORIZONTAL_DOWN, Fg: 34, Bg: -1})
		for y := 1; y < p.Height-1; y++ {
			buf.SetCell(0, y, Cell{Ch: VERTICAL_LINE, Fg: 34, Bg: -1})
		}
		buf.SetCell(0, p.Height-1, Cell{Ch: HORIZONTAL_UP, Fg: 34, Bg: -1})
		return nil
	}

	p.Layout = func() error {
		p.Top = p.Parent().Inner.Min.Y
		p.Height = p.Parent().Inner.Dy()

		return nil
	}
	p.SetLayoutOrder(-100)
	return p, nil
}
