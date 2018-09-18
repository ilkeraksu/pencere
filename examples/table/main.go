package main

import (
	"github.com/ilkeraksu/pencere"
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

	pencere.Loop()

	pencere.Close()
	return nil
}
