package main

import (
	"fmt"

	"github.com/ilkeraksu/pencere"
)

func main() {
	pencere.Init()

	menubar := pencere.NewMenuBar()
	menubar.Data = "ALOO"
	menubar.Fg = 0
	menubar.Bg = 123

	pencere.Root().SetTopBar(menubar)

	b := pencere.NewBox()
	pencere.Root().Add(b)

	b.Label = "ilker"
	b.LabelFg = 230
	b.LabelBg = 0

	// b.Width, b.Height = 60, 30
	b.BorderFg = 230
	b.BorderBg = 0

	// b.Top = 0
	// b.Left = 0

	statusbar := pencere.NewStatusBar()
	statusbar.LayoutOrder = -10
	statusbar.BorderFg = 230
	statusbar.BorderBg = 0

	b.Add(statusbar)

	c1 := pencere.NewBox()
	b.Add(c1)

	c1.Label = "[*] c1"
	c1.LabelFg = 230
	c1.LabelBg = 0

	c1.Width, c1.Height = 10, 10
	c1.BorderFg = 230
	c1.BorderBg = 0

	c1.Top = 5
	c1.Left = 10

	c2 := pencere.NewBox()
	b.Add(c2)
	c2.Label = "c2"
	c2.LabelFg = 230
	c2.LabelBg = 0

	c2.Width, c2.Height = 10, 10
	c2.BorderFg = 230
	c2.BorderBg = 0

	c2.Top = 15
	c2.Left = 22

	c2menubar := pencere.NewMenuBar()
	c2menubar.Data = "TAB1 TAB2 TAB3"
	c2menubar.Fg = 0
	c2menubar.Bg = 123

	c2.SetTopBar(c2menubar)

	c2status := pencere.NewMenuBar()
	c2status.Data = "Window Status"
	c2status.Fg = 0
	c2status.Bg = 123

	c2.SetButtomBar(c2status)

	txt := pencere.NewTextBox()
	txt.HasBorder = true
	txt.Left = 10
	txt.Top = 10
	txt.Width = 40
	txt.Height = 30
	txt.Scrollable = true
	txt.BorderFg = 200
	c2.Add(txt)

	txt2 := pencere.NewTextBox()
	txt2.HasBorder = false
	txt2.Left = 55
	txt2.Top = 10
	txt2.Width = 40
	txt2.Height = 30
	txt2.Bg = 200
	c2.Add(txt2)

	b2 := pencere.NewBox()
	b2.Left = 15
	b2.Top = 5
	b2.Width = 15
	b2.Height = 15
	b2.BorderFg = 230
	b2.BorderBg = 0
	c2.Add(b2)

	b3 := pencere.NewBox()
	b3.BorderFg = 230
	b3.BorderBg = 0
	b3.Layout = func() error {
		b3.Left = b2.Inner.Min.X
		b3.Top = b2.Inner.Min.Y
		b3.Width = b2.Inner.Dx()
		b3.Height = b2.Inner.Dy()
		return nil
	}

	b2.Add(b3)

	v := pencere.NewVerticalSplitter()
	v.Top = 1
	v.Left = 55
	v.Height = 20
	v.Width = 1
	v.LayoutOrder = -2

	buton1 := pencere.NewButton("tıkla")

	buton1.Left = 5
	buton1.Top = 5

	c2.Add(buton1)

	b.Add(v)

	b.Layout = func() error {
		b.Left = b.Parent().Inner.Min.X
		b.Top = b.Parent().Inner.Min.Y

		b.Width = b.Parent().Inner.Dx()
		b.Height = b.Parent().Inner.Dy() - 2

		return nil
	}

	v.Layout = func() error {
		v.Top = 1
		v.Height = v.Parent().Height - statusbar.Height - 2

		return nil
	}

	c2.Layout = func() error {
		c2.Top = v.Top
		c2.Left = v.Left + 1
		c2.Width = c2.Parent().Width - c2.Left - 1
		c2.Height = v.Height
		return nil
	}

	c1.Layout = func() error {
		c1.Top = v.Top
		c1.Left = 1
		c1.Width = v.Left - 1
		c1.Height = v.Height
		return nil
	}

	pencere.Root().OnMouseLeftClick = func(event pencere.MouseEvent) error {
		//panic("TIK Id:" + event.Target.Id)
		statusbar.Data = fmt.Sprintf("TIK Id:%v x:%v y:%v globalX:%v globlY:%v", event.Target.Id, event.X, event.Y, event.GlobalX, event.GlobalY)

		return nil
	}

	// menubar.OnMouseLeftClick = func(event pencere.MouseEvent) error {
	// 	panic("LAAAN")
	// 	return nil
	// }
	pencere.Loop()

	pencere.Close()
}
