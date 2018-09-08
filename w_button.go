package pencere

func NewButton(label string) *Pencere {
	p := NewPencere()
	p.HasBorder = true
	p.Width = len([]rune(label)) + 4
	p.Height = 3
	//p.Label = label
	p.Text = label
	p.Bg = -1
	p.Fg = 230

	p.Render = func(buffer *Buffer) error {
		//Decorate(p, buffer)
		buffer.SetString(1, 1, p.Text, p.Fg, -1)
		return nil
	}

	return p
}
