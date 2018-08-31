package pencere

func NewButton(label string) *Pencere {
	p := NewPencere()
	p.HasBorder = false
	p.Width = len(label) + 4
	p.Height = 1
	p.Label = label
	p.Bg = -1
	p.Fg = 230

	p.Render = func(buffer *Buffer) error {
		buffer.SetString(0, 0, "[ "+p.Label+" ]", p.Fg, -1)
		return nil
	}

	return p
}
