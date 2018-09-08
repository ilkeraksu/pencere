package pencere

import (
	"image"
	"sort"
	"sync"
)

type Event struct {
	Type string
	Data interface{}
}

type Pencere struct {
	parent  *Pencere
	Label   string
	LabelFg Color
	LabelBg Color

	Text string

	Inner image.Rectangle

	Fg Color
	Bg Color

	ZIndex        int
	childs        []*Pencere
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

	LayoutOrder int
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

	propertiesOnce sync.Once
	properties     map[string]interface{}

	topBar    *Pencere
	buttomBar *Pencere
	leftBar   *Pencere
	rightBar  *Pencere
}

func NewPencere() *Pencere {
	p := &Pencere{
		handle:     getHandleId(),
		properties: make(map[string]interface{}),
	}
	return p
}

func (this *Pencere) GetValue(name string) interface{} {
	return this.properties[name]
}

func (this *Pencere) SetValue(name string, value interface{}) {
	this.propertiesOnce.Do(func() {
		this.properties = make(map[string]interface{})
	})
	this.properties[name] = value
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

// func (this *Pencere) render() (*Buffer, error) {

// 	buf := NewBuffer()
// 	buf.SetAreaXY(this.Width, this.Height)
// 	buf.Fill(Cell{' ', ColorDefault, this.Bg})
// 	Decorate(this, buf)

// 	if this.Render == nil {
// 		return buf, nil
// 	}

// 	err := this.Render(buf)
// 	if err != nil {
// 		return nil, e.WrapfEx(err, "", "could not Render")
// 	}
// 	return buf, nil

// }

// func (this *Pencere) layout() error {

// 	if this.Layout != nil {
// 		err := this.Layout()
// 		if err != nil {
// 			return e.WrapfEx(err, "", "could not layout")
// 		}
// 	}

// 	if len(this.childs) == 0 {
// 		return nil
// 	}

// 	ordered := make([]PencereOrder, len(this.childs))
// 	for i, p := range this.childs {
// 		ordered[i].Order = p.LayoutOrder
// 		ordered[i].Pencere = p
// 	}

// 	sort.Slice(ordered, func(i, j int) bool {
// 		return ordered[i].Order < ordered[j].Order
// 	})

// 	for _, o := range ordered {
// 		err := o.Pencere.layout()
// 		if err != nil {
// 			return e.WrapfEx(err, "jds3jf", "could not layout")
// 		}

// 	}
// 	return nil

// }

// func (this *Pencere) dispatchMouseLeftClick(event MouseEvent) {
// 	for _, c := range this.childs {
// 		if c.Left <= event.X && event.X < c.Left+c.Width &&
// 			c.Top <= event.Y && event.Y < c.Top+c.Height {

// 			ce := event
// 			ce.Target = c
// 			ce.X = event.X - c.Left
// 			ce.Y = event.Y - c.Top
// 			c.dispatchMouseLeftClick(ce)
// 			return
// 		}

// 	}

// 	this.bubbleMouseEvent(event)

// }

func SetFocus(p *Pencere) {
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

func getPencereAt(globalX, globalY int) (*Pencere, int, int) {
	p, x, y := getPencereOrChild(root, globalX, globalY)
	return p, x, y
}

func getPencereOrChild(p *Pencere, x, y int) (*Pencere, int, int) {
	ordered := make([]PencereOrder, len(p.childs))
	for i, p := range p.childs {
		ordered[i].Order = p.ZIndex
		ordered[i].Pencere = p
	}

	if p.topBar != nil {
		ordered = append(ordered, PencereOrder{Order: p.topBar.ZIndex, Pencere: p.topBar})
	}

	if p.buttomBar != nil {
		ordered = append(ordered, PencereOrder{Order: p.buttomBar.ZIndex, Pencere: p.buttomBar})
	}

	if p.leftBar != nil {
		ordered = append(ordered, PencereOrder{Order: p.leftBar.ZIndex, Pencere: p.leftBar})
	}

	if p.rightBar != nil {
		ordered = append(ordered, PencereOrder{Order: p.rightBar.ZIndex, Pencere: p.rightBar})
	}

	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].Order == ordered[j].Order {
			return ordered[i].Pencere.handle < ordered[j].Pencere.handle
		}
		return ordered[i].Order < ordered[j].Order
	})

	for _, o := range ordered {
		c := o.Pencere
		if c.Left <= x && x < c.Left+c.Width &&
			c.Top <= y && y < c.Top+c.Height {

			return getPencereOrChild(c, x-c.Left, y-c.Top)
		}
	}
	return p, x, y
}
