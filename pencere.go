package pencere

import (
	"sort"
	"su/pkg/e"
)

type Event struct {
}

type Pencere struct {
	parent  *Pencere
	Label   string
	LabelFg Color
	LabelBg Color

	Fg Color
	Bg Color

	ZIndex        int
	childs        []*Pencere
	Left, Top     int
	Width, Height int
	CanFocus      bool
	HasBorder     bool

	BorderFg Color
	BorderBg Color

	Render  func(buffer *Buffer) error
	Enabled bool
	Visible bool

	LayoutOrder int
	Layout      func() error

	Id   string
	Kind string
	Tag  string

	MouseLeftClick func(event MouseEvent) error
	Data           interface{}
}

func NewPencere() *Pencere {
	p := &Pencere{}

	return p
}

func (this *Pencere) Parent() *Pencere {
	return this.parent
}

func (this *Pencere) Childs() []*Pencere {
	return this.childs
}

func (this *Pencere) Add(p *Pencere) {
	p.parent = this
	this.childs = append(this.childs, p)
}

func (this *Pencere) render() (*Buffer, error) {

	buf := NewBuffer()
	buf.SetAreaXY(this.Width, this.Height)
	buf.Fill(Cell{' ', ColorDefault, this.Bg})

	if this.Render == nil {
		return buf, nil
	}

	err := this.Render(buf)
	if err != nil {
		return nil, e.WrapfEx(err, "", "could not Render")
	}
	return buf, nil

}

func (this *Pencere) layout() error {

	if this.Layout != nil {
		err := this.Layout()
		if err != nil {
			return e.WrapfEx(err, "", "could not layout")
		}
	}

	if len(this.childs) == 0 {
		return nil
	}

	ordered := make([]PencereOrder, len(this.childs))
	for i, p := range this.childs {
		ordered[i].Order = p.LayoutOrder
		ordered[i].Pencere = p
	}

	sort.Slice(ordered, func(i, j int) bool {
		return ordered[i].Order < ordered[j].Order
	})

	for _, o := range ordered {
		err := o.Pencere.layout()
		if err != nil {
			return e.WrapfEx(err, "jds3jf", "could not layout")
		}

	}
	return nil

}

func (this *Pencere) dispatchMouseLeftClick(event MouseEvent) {
	for _, c := range this.childs {
		if c.Left <= event.X && event.X < c.Left+c.Width &&
			c.Top <= event.Y && event.Y < c.Top+c.Height {

			ce := event
			ce.Target = c
			ce.X = event.X - c.Left
			ce.Y = event.Y - c.Top
			c.dispatchMouseLeftClick(ce)
			return
		}

	}

	this.bubbleMouseEvent(event)

}

func (this *Pencere) bubbleMouseEvent(event MouseEvent) {
	if this.MouseLeftClick != nil {
		this.MouseLeftClick(event)
	}

	if this.parent == nil {
		return
	}

	this.parent.bubbleMouseEvent(event)
}
