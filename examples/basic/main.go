package main

import (
	"github.com/ilkeraksu/pencere"
)

func main() {
	pencere.Init()

	b := pencere.NewBox()
	pencere.Root().Add(b)

	b.Label = "ilker"
	b.LabelFg = 230
	b.LabelBg = 0

	b.Width, b.Height = 60, 30
	b.BorderFg = 230
	b.BorderBg = 0

	b.Top = 0
	b.Left = 0

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

	v := pencere.NewVerticalSplitter()
	v.Top = 1
	v.Left = 55
	v.Height = 20
	v.Width = 1
	v.LayoutOrder = -2

	buton1 := pencere.NewButton("tıkla")
	buton1.Id = "DENEMEID"
	buton1.Left = 5
	buton1.Top = 5

	c2.Add(buton1)

	b.Add(v)

	b.Layout = func() error {

		b.Width = b.Parent().Width - 2
		b.Height = b.Parent().Height - 3

		return nil
	}

	v.Layout = func() error {
		v.Top = 1
		v.Height = v.Parent().Height - statusbar.Height - 2 - 2

		return nil
	}

	c2.Layout = func() error {
		c2.Top = v.Top
		c2.Left = v.Left + 1
		c2.Width = c2.Parent().Width - c2.Left - 2
		c2.Height = v.Height
		return nil
	}

	c1.Layout = func() error {
		c1.Top = v.Top
		c1.Left = 1
		c1.Width = v.Left - 3
		c1.Height = v.Height
		return nil
	}

	b.MouseLeftClick = func(event pencere.MouseEvent) error {
		//panic("TIK Id:" + event.Target.Id)
		statusbar.Data = "TIK Id:" + event.Target.Id
		pencere.Render()
		return nil
	}

	pencere.Loop()

	pencere.Close()
}
