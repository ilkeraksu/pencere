package pencere

import (
	"image"
	"sort"

	"github.com/pkg/errors"
)

type Pencere struct {
	parent  *Pencere
	Label   string
	LabelFg Color
	LabelBg Color

	Text string

	Inner image.Rectangle

	Fg      Color
	Bg      Color
	Texture rune
	zindex  int
	childs  []*Pencere

	layoutOrderCache []PencereOrder
	zindexOrderCache []PencereOrder

	Left, Top     int
	Width, Height int
	CanFocus      bool
	HasFocus      bool
	HasBorder     bool

	Scrollable bool
	ContentX   int
	ContentY   int

	BorderFg Color
	BorderBg Color

	Render  func(buffer *Buffer) error
	Enabled bool
	Visible bool

	layoutOrder int
	Layout      func() error

	handle uint64
	Id     string
	Kind   string
	Tag    string

	OnMouseLeftClick func(event MouseEvent) error
	OnFocus          func() error
	OnLostFocus      func() error

	OnKeyEvent func(event KeyEvent) error
	Data       interface{}
	Controller interface{}

	OnDragBegin func(event DragBeginEvent) (bool, *DragContext, error)
	OnDragging  func(event DraggingEvent) error
	OnDragEnd   func(event DragEndEvent) error
	OnDragEnter func(event DragEnterEvent) error
	OnDragLeave func(event DragLeaveEvent) error
	OnDragOver  func(event DragOverEvent) error
	OnDrop      func(event DropEvent) error
	CanDrop     func(event DropEvent) (bool, error)

	Properties Properties

	topBar    *Pencere
	buttomBar *Pencere
	leftBar   *Pencere
	rightBar  *Pencere
	Theme     *Theme
}

func NewPencere(options ...Option) (*Pencere, error) {
	p := &Pencere{
		handle:     getHandleId(),
		Properties: NewProperties(),
		Theme:      DefaultTheme,
		Enabled:    true,
		Visible:    true,
	}

	style := p.Theme.Style("default")
	p.Fg = style.Fg
	p.Bg = style.Bg

	labelStyle := p.Theme.Style("label")
	p.Fg = labelStyle.Fg
	p.Bg = labelStyle.Bg

	borderStyle := p.Theme.Style("border")
	p.Fg = borderStyle.Fg
	p.Bg = borderStyle.Bg

	err := p.Apply(options...)
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (this *Pencere) Apply(options ...Option) error {
	for _, option := range options {
		err := option(this)
		if err != nil {
			return errors.Wrapf(err, "can not set option")
		}
	}
	return nil
}

func (this *Pencere) Parent() *Pencere {
	return this.parent
}

func (this *Pencere) Childs() []*Pencere {
	return this.childs
}

func (this *Pencere) AddPencere(p *Pencere) {

	p.parent = this
	this.childs = append(this.childs, p)

}

func (this *Pencere) AddControl(c Control) {
	p := c.Pencere()
	p.parent = this
	this.childs = append(this.childs, p)
}

func (this *Pencere) SetTopBar(p *Pencere) {
	p.parent = this
	this.topBar = p
}
func (this *Pencere) SetButtomBar(p *Pencere) {
	p.parent = this
	this.buttomBar = p
}
func (this *Pencere) SetLeftBar(p *Pencere) {
	p.parent = this
	this.leftBar = p
}
func (this *Pencere) SetRightBar(p *Pencere) {
	p.parent = this
	this.rightBar = p
}

func (this *Pencere) FireEvent(eventType string, data interface{}) error {
	return nil
}

func (this *Pencere) GetZIndex() int {
	return this.zindex
}

func (this *Pencere) SetZIndex(zindex int) {
	this.zindex = zindex
	parent := this.Parent()
	if parent != nil {
		parent.resetZIndexCache()
	}
}

func (this *Pencere) resetZIndexCache() {
	this.zindexOrderCache = nil
}

func (this *Pencere) GetLayoutOrder() int {
	return this.layoutOrder
}

func (this *Pencere) SetLayoutOrder(layoutOrder int) {
	this.layoutOrder = layoutOrder
	parent := this.Parent()
	if parent != nil {
		parent.resetLayoutOrderCache()
	}
}

func (this *Pencere) resetLayoutOrderCache() {
	this.layoutOrderCache = nil
}

func SetFocus(p *Pencere) {
	if p == nil {
		return
	}

	if !p.CanFocus {
		return
	}

	if focusedPencere != nil {
		focusedPencere.HasFocus = false
		if focusedPencere.OnLostFocus != nil {
			focusedPencere.OnLostFocus()
		}
	}

	if p.OnFocus != nil {
		p.OnFocus()
	}

	focusedPencere = p
	p.HasFocus = true
}

func (this *Pencere) bubbleMouseEvent(event MouseEvent) {
	if this.OnMouseLeftClick != nil {
		this.OnMouseLeftClick(event)
	}

	if this.parent == nil {
		return
	}

	this.parent.bubbleMouseEvent(event)
}

func (this *Pencere) TranslateToXY(globalX int, globalY int) (int, int) {
	x, y := this.GlobalPosition()

	return globalX - x, globalY - y
}

func (this *Pencere) GlobalPosition() (int, int) {
	var x, y int

	p := this
	for p != nil {
		x += p.Left
		y += p.Left

		p = p.Parent()
	}

	return x, y
}

func getPencereAt(globalX, globalY int) (*Pencere, int, int) {
	p, x, y := getPencereOrChild(root, globalX, globalY)
	return p, x, y
}

func getPencereOrChild(p *Pencere, x, y int) (*Pencere, int, int) {
	if p.zindexOrderCache == nil {

		ordered := make([]PencereOrder, len(p.childs))
		for i, p := range p.childs {
			ordered[i].Order = p.zindex
			ordered[i].Pencere = p
		}

		if p.topBar != nil {
			ordered = append(ordered, PencereOrder{Order: p.topBar.zindex, Pencere: p.topBar})
		}

		if p.buttomBar != nil {
			ordered = append(ordered, PencereOrder{Order: p.buttomBar.zindex, Pencere: p.buttomBar})
		}

		if p.leftBar != nil {
			ordered = append(ordered, PencereOrder{Order: p.leftBar.zindex, Pencere: p.leftBar})
		}

		if p.rightBar != nil {
			ordered = append(ordered, PencereOrder{Order: p.rightBar.zindex, Pencere: p.rightBar})
		}

		sort.Slice(ordered, func(i, j int) bool {
			if ordered[i].Order == ordered[j].Order {
				return ordered[i].Pencere.handle < ordered[j].Pencere.handle
			}
			return ordered[i].Order < ordered[j].Order
		})

		p.zindexOrderCache = ordered
	}
	for _, o := range p.zindexOrderCache {
		c := o.Pencere
		if c.Visible && c.Enabled && c.Left <= x && x < c.Left+c.Width &&
			c.Top <= y && y < c.Top+c.Height {

			return getPencereOrChild(c, x-c.Left, y-c.Top)
		}
	}
	return p, x, y
}
