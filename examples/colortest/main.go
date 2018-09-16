package main

import (
	"github.com/ilkeraksu/pencere"
)

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

func run() error {
	pencere.Init()

	box, err := pencere.NewBox(pencere.Position(2, 2, 20, 10))
	if err != nil {
		return err
	}

	box.Render = func(buf *pencere.Buffer) error {
		fg := pencere.Color(8)
		bg := pencere.Color(3)
		buf.SetString(0, 0, "    ▁▂▂▟ ", fg, bg)
		buf.SetString(0, 1, "  ▟█████▙  ", fg, bg)
		buf.SetString(0, 2, "  ███▘ ▜█▙ ", fg, bg)
		buf.SetString(0, 3, " ▟████▙      ", fg, bg)
		buf.SetString(0, 4, "DEEME", fg, bg)

		return nil
	}

	pencere.Root().Add(box)

	pencere.Loop()

	pencere.Close()
	return nil
}
