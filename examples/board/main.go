package main

import (
	"fmt"

	"github.com/ilkeraksu/pencere"
)

func main() {

	err := run()
	if err != nil {
		panic(err)
	}

}

var boardRecs map[string]*pencere.Pencere = make(map[string]*pencere.Pencere)
var letters = "abcdefgh"
var numbers = "12345678"

func run() error {
	pencere.Init()

	board, err := newBoard()
	if err != nil {
		return err
	}

	pencere.Root().AddPencere(board)

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			rec := newRecWidget(x, y)
			boardRecs[string(letters[x])+fmt.Sprintf("%v", 7-y+1)] = rec
			board.AddPencere(rec)
		}
	}

	for x := 1; x < 8; x++ {
		for y := 0; y < 8; y++ {

			edge := newVerticalEdgeWidget(x, y)
			board.AddPencere(edge)

		}
	}

	for x := 0; x < 8; x++ {
		for y := 0; y < 9; y++ {
			hedge := newHorizontalEdgeWidget(x, y)
			board.AddPencere(hedge)

		}
	}

	boardRecs["a1"].SetValue("tas", "castle")
	boardRecs["a1"].SetValue("tascolor", 77)
	boardRecs["b1"].SetValue("tas", "horse")
	boardRecs["c1"].SetValue("tas", "bishop")
	boardRecs["f1"].SetValue("tas", "bishop")
	boardRecs["g1"].SetValue("tas", "horse")
	boardRecs["h1"].SetValue("tas", "castle")

	boardRecs["a2"].SetValue("tas", "pawn")
	boardRecs["b2"].SetValue("tas", "pawn")
	boardRecs["c2"].SetValue("tas", "pawn")
	boardRecs["d2"].SetValue("tas", "pawn")
	boardRecs["e2"].SetValue("tas", "pawn")
	boardRecs["f2"].SetValue("tas", "pawn")
	boardRecs["g2"].SetValue("tas", "pawn")
	boardRecs["h2"].SetValue("tas", "pawn")

	pencere.Loop()

	pencere.Close()
	return nil
}

var fullWidth = 12
var fullheight = 6
var recColor pencere.Color = 4
var boardMarginTop = 4
var boardMarginleft = 6

func newRecWidget(x, y int) *pencere.Pencere {

	p, err := pencere.NewPencere(pencere.Position(boardMarginleft+x*fullWidth, boardMarginTop+y*fullheight, 10, 5))
	p.Bg = recColor
	p.Fg = 3
	_ = err
	p.HasBorder = false
	p.CanFocus = true

	//p.Texture = '░'
	//p.Texture = '▒'
	//p.Texture = '■'
	p.OnMouseLeftClick = func(event pencere.MouseEvent) error {
		//p.SetValue("tas", "castle")
		return nil
	}

	p.Render = func(buf *pencere.Buffer) error {
		var tasFg pencere.Color = pencere.Color(p.GetInt("tascolor", -1))
		//var tasFg pencere.Color = 110 //pencere.Color(p.GetInt("tascolor", 110))

		switch fmt.Sprintf("%v", p.GetValue("tas")) {
		case "pawn":
			drawPawn(buf, tasFg, p.Bg)
		case "bishop":
			drawBishop(buf, tasFg, p.Bg)
		case "castle":
			drawCastle(buf, 2, p.Bg)
		case "horse":
			drawHorse(buf, tasFg, p.Bg)
		}
		return nil
	}

	p.OnFocus = func() error {

		p.Bg = 9
		return nil
	}

	p.OnLostFocus = func() error {
		p.Bg = recColor
		return nil
	}

	p.OnKeyEvent = func(event pencere.KeyEvent) error {
		switch event.Rune {
		case 'c':
			p.SetValue("tas", "castle")
		case 'h':
			p.SetValue("tas", "horse")

		}

		return nil
	}

	return p
}

func drawPawn(buf *pencere.Buffer, fg, bg pencere.Color) {
	fg = 192
	bg = 78
	buf.SetString(0, 1, "  ilker   ", fg, bg)
	buf.SetString(0, 2, "        ", fg, bg)
	buf.SetString(0, 3, "   ▟▙  ", fg, bg)
	buf.SetString(0, 4, "   ██  ", fg, bg)
}

func drawBishop(buf *pencere.Buffer, fg, bg pencere.Color) {
	buf.SetString(0, 1, "   ▟▙  ", fg, bg)
	buf.SetString(0, 2, "   ▜▛  ", fg, bg)
	buf.SetString(0, 3, "   ▟▙  ", fg, bg)
	buf.SetString(0, 4, "   ██  ", fg, bg)
}

func drawHorse(buf *pencere.Buffer, fg, bg pencere.Color) {

	buf.SetString(0, 0, "   ▁▂▂▟ ", fg, bg)
	buf.SetString(0, 1, " ▟█████▙  ", fg, bg)
	buf.SetString(0, 2, " ███▘ ▜█▙ ", fg, bg)
	buf.SetString(0, 3, " ████▙      ", fg, bg)
	buf.SetString(0, 4, "▟███████▙   ", fg, bg)
}

func drawCastle(buf *pencere.Buffer, fg pencere.Color, bg pencere.Color) {

	buf.SetString(0, 0, "  ▄ ▄ ▄ ", fg, bg)
	buf.SetString(0, 1, " ▜█████▛  ", fg, bg)
	buf.SetString(0, 2, "  █████ ", fg, bg)
	buf.SetString(0, 3, "  █████    ", fg, bg)
	buf.SetString(0, 4, " ▟█████▙   ", fg, bg)

}

func drawKing(buf *pencere.Buffer, fg, bg pencere.Color) {
	buf.SetString(0, 0, "    ╋  ", fg, bg)
	buf.SetString(0, 1, "   ▟▙  ", fg, bg)
	buf.SetString(0, 2, "  █████ ", fg, bg)
	buf.SetString(0, 3, "  █████      ", fg, bg)
	buf.SetString(0, 4, " ▟█████▙   ", fg, bg)
}

func newVerticalEdgeWidget(x, y int) *pencere.Pencere {

	p, err := pencere.NewPencere(pencere.Position(boardMarginleft+x*fullWidth-2, boardMarginTop+y*fullheight, 2, fullheight-1))
	//p.Bg = 183
	_ = err
	p.HasBorder = false

	p.OnMouseLeftClick = func(event pencere.MouseEvent) error {
		p.Bg = 255
		return nil
	}
	return p
}

func newHorizontalEdgeWidget(x, y int) *pencere.Pencere {

	p, err := pencere.NewPencere(pencere.Position(boardMarginleft+x*fullWidth, boardMarginTop+y*fullheight-1, fullWidth-2, 1))
	//p.Bg = 183
	//p.ZIndex = 10
	_ = err
	p.HasBorder = false

	p.OnMouseLeftClick = func(event pencere.MouseEvent) error {
		p.Bg = 255
		return nil
	}
	return p
}

func newBoard() (*pencere.Pencere, error) {
	board, err := pencere.NewBox(pencere.Position(0, 0, 122, 64))
	if err != nil {
		return nil, err
	}

	board.Render = func(buf *pencere.Buffer) error {
		for x := 0; x < 8; x++ {
			buf.SetString(4+x*fullWidth+(fullWidth/2), 2, string(letters[x]), board.Fg, board.Bg)
			buf.SetString(4+x*fullWidth+(fullWidth/2), 60, string(letters[x]), board.Fg, board.Bg)
		}
		for y := 0; y < 8; y++ {
			buf.SetString(3, 4+y*fullheight+(fullheight/2), string(numbers[7-y]), board.Fg, board.Bg)
			buf.SetString(118, 4+y*fullheight+(fullheight/2), string(numbers[7-y]), board.Fg, board.Bg)
		}
		return nil
	}
	return board, nil
}
