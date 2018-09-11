package main

import (
	"fmt"

	"github.com/ilkeraksu/pencere"
)

var output chan string

func init() {
	output = make(chan string)
}

func Out(s string) {

	output <- s
}

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

func run() error {
	pencere.Init()

	menubar, err := pencere.NewMenuBar()
	if err != nil {
		return err
	}
	menubar.Data = "ALOO"

	pencere.Root().SetTopBar(menubar)

	output, err := pencere.NewOutput(output)
	output.Width = 50
	output.Label = "output"

	pencere.Root().SetRightBar(output)

	b, err := pencere.NewBox(pencere.LayoutFill())
	if err != nil {
		return err
	}

	pencere.Root().Add(b)

	b.Label = "ilker"

	b.HasBorder = false

	statusbar, err := pencere.NewStatusBar()
	if err != nil {
		return err
	}

	statusbar.Height = 3

	b.SetButtomBar(statusbar)

	c1, err := pencere.NewBox(pencere.Position(10, 5, 10, 10))
	if err != nil {
		return err
	}

	b.Add(c1)

	c1.Label = "[x] c1"

	c2, err := pencere.NewBox(pencere.Position(15, 22, 10, 10))
	if err != nil {
		return err
	}

	b.Add(c2)
	c2.Label = "c2"

	c2.Scrollable = true
	c2.ContentY = 300

	c2menubar, err := pencere.NewMenuBar()
	if err != nil {
		return err
	}
	c2menubar.Data = "TAB1 TAB2 TAB3"

	c2.SetTopBar(c2menubar)

	c2status, err := pencere.NewMenuBar()
	if err != nil {
		return err
	}
	c2status.Data = "Window Status"

	c2.SetButtomBar(c2status)

	txt, err := pencere.NewTextBox(pencere.Position(10, 10, 40, 30))
	if err != nil {
		return err
	}
	txt.HasBorder = true

	txt.Scrollable = true

	c2.Add(txt)

	txt2, err := pencere.NewTextBox(pencere.Position(55, 10, 40, 30))
	if err != nil {
		return err
	}

	txt2.HasBorder = false

	c2.Add(txt2)

	b2, err := pencere.NewBox(pencere.LayoutFill())
	if err != nil {
		return err
	}

	c2.Add(b2)

	b3, err := pencere.NewBox(pencere.LayoutFill())
	if err != nil {
		return err
	}

	b2.Add(b3)

	v, err := pencere.NewVerticalSplitter()
	if err != nil {
		return err
	}

	v.Left = 55

	buton1, err := pencere.NewButton("tıkla", pencere.Position(5, 5, 10, 3))
	if err != nil {
		return err
	}

	c2.Add(buton1)

	b.Add(v)

	c2.Layout = func() error {
		c2.Top = v.Top
		c2.Left = v.Left + 1
		c2.Width = c2.Parent().Inner.Max.X - c2.Left
		c2.Height = v.Height
		return nil
	}

	c1.Layout = func() error {
		c1.Top = v.Top
		c1.Left = c1.Parent().Inner.Min.X
		c1.Width = v.Left - c1.Left
		c1.Height = v.Height
		return nil
	}

	pencere.Root().OnMouseLeftClick = func(event pencere.MouseEvent) error {
		//panic("TIK Id:" + event.Target.Id)
		statusbar.Data = fmt.Sprintf("TIK Id:%v x:%v y:%v globalX:%v globlY:%v", event.Target.Id, event.X, event.Y, event.GlobalX, event.GlobalY)

		Out(fmt.Sprintf("TIK Id:%v x:%v y:%v globalX:%v globlY:%v", event.Target.Id, event.X, event.Y, event.GlobalX, event.GlobalY))
		return nil
	}

	pencere.Loop()

	pencere.Close()
	return nil
}
