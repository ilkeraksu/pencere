package pencere

func Decorate(this *Pencere, buf *Buffer) {
	if this.HasBorder {
		drawBorder(this, buf)
	}
	if this.Label != "" {
		drawLabel(this, buf)
	}
}

func drawBorder(this *Pencere, buf *Buffer) {
	x := this.Width + 1
	y := this.Height + 1

	// draw lines
	buf.Merge(NewFilledBuffer(0, 0, x, 1, Cell{HORIZONTAL_LINE, this.BorderFg, this.BorderBg}))
	buf.Merge(NewFilledBuffer(0, y, x, y+1, Cell{HORIZONTAL_LINE, this.BorderFg, this.BorderBg}))
	buf.Merge(NewFilledBuffer(0, 0, 1, y+1, Cell{VERTICAL_LINE, this.BorderFg, this.BorderBg}))
	buf.Merge(NewFilledBuffer(x, 0, x+1, y+1, Cell{VERTICAL_LINE, this.BorderFg, this.BorderBg}))

	// draw corners
	buf.SetCell(0, 0, Cell{TOP_LEFT, this.BorderFg, this.BorderBg})
	buf.SetCell(x, 0, Cell{TOP_RIGHT, this.BorderFg, this.BorderBg})
	buf.SetCell(0, y, Cell{BOTTOM_LEFT, this.BorderFg, this.BorderBg})
	buf.SetCell(x, y, Cell{BOTTOM_RIGHT, this.BorderFg, this.BorderBg})
}

func drawLabel(this *Pencere, buf *Buffer) {
	r := MaxString(this.Label, (this.Width-3)-1)
	buf.SetString(3, 0, r, this.LabelFg, this.LabelBg)
	if this.Label == "" {
		return
	}
	c := Cell{' ', this.Fg, this.Bg}
	buf.SetCell(2, 0, c)
	if len(this.Label)+3 < this.Width {
		buf.SetCell(len(this.Label)+3, 0, c)
	} else {
		buf.SetCell(this.Width-1, 0, c)
	}
}
