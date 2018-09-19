package main

import (
	"github.com/ilkeraksu/pencere"
	"github.com/pkg/errors"
)

//var output chan string

func init() {
	//	output = make(chan string)
}

// func Out(s string) {

// 	output <- s
// }

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

func run() error {
	pencere.Init()

	table := pencere.NewTable()
	table.Width = 50
	table.Height = 30
	table.HasBorder = true
	table.CursorColor = 123
	table.Cursor = true

	table.Header = []string{"Title"}
	table.ColWidths = []int{10}
	table.Rows = [][]string{{"Apple"}, {"orange"}}
	pencere.Root().AddWindow(table.Pencere)

	box, err := pencere.NewBox()
	if err != nil {
		return errors.Wrapf(err, "could not create Box")
	}
	pencere.Root().AddWindow(box)

	splitter, err := pencere.NewVerticalSplitter()
	splitter.Left = 100
	pencere.Root().AddWindow(splitter.Pencere)

	if err != nil {
		return errors.Wrapf(err, "could not create verticalsplitter")
	}

	splitter.LeftSide = table.Window()
	splitter.RightSide = box

	pencere.Loop()

	pencere.Close()
	return nil
}
