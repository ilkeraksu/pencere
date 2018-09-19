package pencere

type VerticalSplitter struct {
	*Pencere
	LeftSide  *Pencere
	RightSide *Pencere
}

func NewVerticalSplitter(options ...Option) (*VerticalSplitter, error) {
	p, err := NewPencere(options...)
	if err != nil {
		return nil, err
	}
	p.HasBorder = false
	p.Width = 1
	p.SetLayoutOrder(-100)
	this := &VerticalSplitter{
		Pencere: p,
	}

	p.Render = this.Render
	p.Layout = this.Layout
	p.OnDragBegin = this.OnDragBegin
	return this, nil
}

func (this *VerticalSplitter) Render(buf *Buffer) error {

	buf.SetCell(0, 0, Cell{Ch: HORIZONTAL_DOWN, Fg: 34, Bg: -1})
	for y := 1; y < this.Height-1; y++ {
		buf.SetCell(0, y, Cell{Ch: VERTICAL_LINE, Fg: 34, Bg: -1})
	}
	buf.SetCell(0, this.Height-1, Cell{Ch: HORIZONTAL_UP, Fg: 34, Bg: -1})
	return nil

}

func (this *VerticalSplitter) Layout() error {
	this.Top = this.Parent().Inner.Min.Y
	this.Height = this.Parent().Inner.Dy()

	if c := this.LeftSide; c != nil {
		c.Top = this.Top
		c.Left = c.Parent().Inner.Min.X
		c.Width = this.Left - c.Left
		c.Height = this.Height
	}

	if c := this.RightSide; c != nil {
		c.Top = this.Top
		c.Left = this.Left + 1
		c.Width = c.Parent().Inner.Max.X - c.Left
		c.Height = this.Height

	}
	return nil
}

func (this *VerticalSplitter) OnDragBegin(event DragBeginEvent) (bool, *DragContext, error) {
	dragContext := NewDragContext(this.Pencere)

	icon, _ := NewPencere(Position(this.Left, this.Top, this.Width, this.Height))
	icon.Bg = this.Bg
	icon.Fg = this.Fg
	icon.Texture = 'X'
	dragContext.DraggingIcon = icon
	return true, dragContext, nil
}
