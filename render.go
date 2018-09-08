package pencere

import (
	"sort"
	"su/pkg/e"
	"sync"

	tb "github.com/nsf/termbox-go"
)

// Bufferer should be implemented by all renderable components.
type Bufferer interface {
	Buffer() *Buffer
	GetXOffset() int
	GetYOffset() int
}

// Render renders all Bufferers in the given order to termbox, then asks termbox to print the screen.
func RenderOLD(bs ...Bufferer) {
	var wg sync.WaitGroup
	for _, b := range bs {
		wg.Add(1)
		go func(b Bufferer) {
			defer wg.Done()
			buf := b.Buffer()
			// set cells in buf
			for p, c := range buf.CellMap {
				if p.In(buf.Area) {
					tb.SetCell(p.X+b.GetXOffset(), p.Y+b.GetYOffset(), c.Ch, tb.Attribute(c.Fg)+1, tb.Attribute(c.Bg)+1)
				}
			}
		}(b)
	}

	wg.Wait()
	tb.Flush()
}

func Render() error {

	select {
	case eventStream.shouldRender <- true:
	default:
		isRenderDirty = true
	}

	return nil
}

var lock sync.Mutex

// Render renders all Bufferers in the given order to termbox, then asks termbox to print the screen.
func render() error {

	layoutPencere(root)
	isRenderDirty = false
	//buf, err := root.render()
	buf, err := renderPencere(root)

	if err != nil {
		return e.WrapfEx(err, "4kd", "could not Render")
	}

	for p, c := range buf.CellMap {
		if p.In(buf.Area) {

			tb.SetCell(p.X, p.Y, c.Ch, tb.Attribute(c.Fg)+1, tb.Attribute(c.Bg)+1)

		}
	}

	tb.Flush()
	return nil
}

func renderPencere(p *Pencere) (*Buffer, error) {

	buf := NewBuffer()
	buf.SetAreaXY(p.Width, p.Height)
	buf.Fill(Cell{' ', ColorDefault, p.Bg})

	if p.Render != nil {
		err := p.Render(buf)
		if err != nil {
			return nil, e.WrapfEx(err, "", "could not Render")
		}
	}

	ordered := make([]PencereOrder, len(p.childs))
	for i, p := range p.childs {
		ordered[i].Order = p.ZIndex
		ordered[i].Pencere = p
	}

	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].Order == ordered[j].Order {
			return ordered[i].Pencere.handle > ordered[j].Pencere.handle
		}
		return ordered[i].Order < ordered[j].Order
	})

	for _, o := range ordered {
		c := o.Pencere
		childBuffer, err := renderPencere(c)
		if err != nil {
			return nil, e.WrapfEx(err, "jdjf", "could not render")
		}

		//buf.MergeChildArea(childBuffer, c.Left+c.Inner.Min.X, c.Top+c.Inner.Min.Y, c.Inner.Dx(), c.Inner.Dy())
		buf.MergeChildArea(childBuffer, c.Left, c.Top, c.Width, c.Height)

	}
	if p.topBar != nil {
		c := p.topBar
		childBuffer, err := renderPencere(p.topBar)
		if err != nil {
			return nil, e.WrapfEx(err, "jdjf", "could not render")
		}

		buf.MergeChildArea(childBuffer, c.Left, c.Top, c.Width, c.Height)
	}

	if p.buttomBar != nil {
		c := p.buttomBar
		childBuffer, err := renderPencere(p.buttomBar)
		if err != nil {
			return nil, e.WrapfEx(err, "jdjf", "could not render")
		}

		buf.MergeChildArea(childBuffer, c.Left, c.Top, c.Width, c.Height)
	}

	if p.leftBar != nil {
		c := p.leftBar
		childBuffer, err := renderPencere(p.leftBar)
		if err != nil {
			return nil, e.WrapfEx(err, "jdjf", "could not render")
		}

		buf.MergeChildArea(childBuffer, c.Left, c.Top, c.Width, c.Height)
	}

	if p.rightBar != nil {
		c := p.rightBar
		childBuffer, err := renderPencere(p.rightBar)
		if err != nil {
			return nil, e.WrapfEx(err, "jdjf", "could not render")
		}

		buf.MergeChildArea(childBuffer, c.Left, c.Top, c.Width, c.Height)
	}

	Decorate(p, buf)
	return buf, nil

}

// Clear clears the screen with the default Bg color.
func Clear() {
	tb.Clear(tb.ColorDefault+1, tb.Attribute(Theme.Bg)+1)
}

// func Layout() error {
// 	return layoutPencere(root)
// }

func layoutPencere(p *Pencere) error {

	if p.HasBorder {
		p.Inner.Min.X = 1
		p.Inner.Min.Y = 1
		p.Inner.Max.X = p.Width - 1
		p.Inner.Max.Y = p.Height - 1

	} else {
		p.Inner.Min.X = 0
		p.Inner.Min.Y = 0
		p.Inner.Max.X = p.Width + 1
		p.Inner.Max.Y = p.Height + 1
	}

	if p.topBar != nil {

		ts := p.topBar
		ts.Left = p.Inner.Min.X
		ts.Width = p.Inner.Dx()
		ts.Top = p.Inner.Min.Y

		p.Inner.Min.Y = p.Inner.Min.Y + ts.Height
	}

	if p.buttomBar != nil {

		ts := p.buttomBar
		ts.Left = p.Inner.Min.X
		ts.Width = p.Inner.Dx()
		ts.Top = p.Inner.Max.Y - ts.Height

		p.Inner.Max.Y = p.Inner.Max.Y - ts.Height
	}

	if p.Layout != nil {
		err := p.Layout()
		if err != nil {
			return e.WrapfEx(err, "", "could not layout")
		}
	}

	if len(p.childs) == 0 {
		return nil
	}

	ordered := make([]PencereOrder, len(p.childs))
	for i, p := range p.childs {
		ordered[i].Order = p.LayoutOrder
		ordered[i].Pencere = p
	}

	sort.Slice(ordered, func(i, j int) bool {
		if ordered[i].Order == ordered[j].Order {
			return ordered[i].Pencere.handle < ordered[j].Pencere.handle
		}
		return ordered[i].Order < ordered[j].Order
	})

	for _, o := range ordered {
		err := layoutPencere(o.Pencere)
		if err != nil {
			return e.WrapfEx(err, "jds3jf", "could not layout")
		}

	}
	return nil
}
