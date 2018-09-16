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

var letters = "abcdefgh"
var numbers = "12345678"

var recColor pencere.Color = 47
var recWhiteColor pencere.Color = 45
var focusColor pencere.Color = 1

var boardMarginTop = 4
var boardMarginleft = 6

func run() error {
	pencere.Init()

	board, err := newBoard()
	if err != nil {
		return err
	}

	pencere.Root().AddControl(board)

	board.SetPiece("a1", "castle", board.WhitePieceColor)
	board.SetPiece("b1", "horse", board.WhitePieceColor)
	board.SetPiece("c1", "bishop", board.WhitePieceColor)
	board.SetPiece("d1", "king", board.WhitePieceColor)
	board.SetPiece("e1", "queen", board.WhitePieceColor)
	board.SetPiece("f1", "bishop", board.WhitePieceColor)
	board.SetPiece("g1", "horse", board.WhitePieceColor)
	board.SetPiece("h1", "castle", board.WhitePieceColor)

	board.SetPiece("a2", "pawn", board.WhitePieceColor)
	board.SetPiece("b2", "pawn", board.WhitePieceColor)
	board.SetPiece("c2", "pawn", board.WhitePieceColor)
	board.SetPiece("d2", "pawn", board.WhitePieceColor)
	board.SetPiece("e2", "pawn", board.WhitePieceColor)
	board.SetPiece("f2", "pawn", board.WhitePieceColor)
	board.SetPiece("g2", "pawn", board.WhitePieceColor)
	board.SetPiece("h2", "pawn", board.WhitePieceColor)

	board.SetPiece("a8", "castle", board.BlackPieceColor)
	board.SetPiece("b8", "horse", board.BlackPieceColor)
	board.SetPiece("c8", "bishop", board.BlackPieceColor)
	board.SetPiece("d8", "king", board.BlackPieceColor)
	board.SetPiece("e8", "queen", board.BlackPieceColor)
	board.SetPiece("f8", "bishop", board.BlackPieceColor)
	board.SetPiece("g8", "horse", board.BlackPieceColor)
	board.SetPiece("h8", "castle", board.BlackPieceColor)

	board.SetPiece("a7", "pawn", board.BlackPieceColor)
	board.SetPiece("b7", "pawn", board.BlackPieceColor)
	board.SetPiece("c7", "pawn", board.BlackPieceColor)
	board.SetPiece("d7", "pawn", board.BlackPieceColor)
	board.SetPiece("e7", "pawn", board.BlackPieceColor)
	board.SetPiece("f7", "pawn", board.BlackPieceColor)
	board.SetPiece("g7", "pawn", board.BlackPieceColor)
	board.SetPiece("h7", "pawn", board.BlackPieceColor)

	pencere.Loop()

	pencere.Close()
	return nil
}

// func newRecWidget(x, y int) *pencere.Pencere {

// 	p, err := pencere.NewPencere(pencere.Position(boardMarginleft+x*fullWidth, boardMarginTop+y*fullheight, fullWidth, fullheight))
// 	p.Bg = recColor
// 	p.Fg = 3
// 	_ = err
// 	p.HasBorder = false
// 	p.CanFocus = true

// 	if (x+y)%2 == 0 {
// 		p.Bg = recWhiteColor
// 	}

// 	//p.Texture = '░'
// 	//p.Texture = '▒'
// 	//p.Texture = '■'
// 	p.OnMouseLeftClick = func(event pencere.MouseEvent) error {
// 		//p.SetValue("tas", "castle")
// 		return nil
// 	}

// 	// p.Render = func(buf *pencere.Buffer) error {

// 	// 	var tasFg pencere.Color = pencere.Color(p.GetInt("tascolor", -1))
// 	// 	//var tasFg pencere.Color = 110 //pencere.Color(p.GetInt("tascolor", 110))

// 	// 	switch fmt.Sprintf("%v", p.GetValue("tas")) {
// 	// 	case "pawn":
// 	// 		drawPawn(buf, tasFg, p.Bg)
// 	// 	case "bishop":
// 	// 		drawBishop(buf, tasFg, p.Bg)
// 	// 	case "castle":
// 	// 		drawCastle(buf, tasFg, p.Bg)
// 	// 	case "horse":
// 	// 		drawHorse(buf, tasFg, p.Bg)
// 	// 	case "king":
// 	// 		drawKing(buf, tasFg, p.Bg)
// 	// 	case "queen":
// 	// 		drawQueen(buf, tasFg, p.Bg)
// 	// 	}
// 	// 	return nil
// 	// }

// 	// p.OnFocus = func() error {
// 	// 	p.SetValue("Bg", int(p.Bg))
// 	// 	p.Bg = focusColor
// 	// 	return nil
// 	// }

// 	// p.OnLostFocus = func() error {
// 	// 	p.Bg = pencere.Color(p.GetInt("Bg", int(p.Bg)))
// 	// 	return nil
// // 	// }

// 	p.OnKeyEvent = func(event pencere.KeyEvent) error {
// 		switch event.Rune {
// 		case 'c':
// 			p.SetValue("tas", "castle")
// 		case 'h':
// 			p.SetValue("tas", "horse")

// 		}

// 		return nil
// 	}

// 	return p
// }

func drawPawn(buf *pencere.Buffer, fg, bg pencere.Color) {

	buf.SetString(0, 1, "        ", fg, bg)

	buf.SetString(0, 2, "    ▟▙  ", fg, bg)
	buf.SetString(0, 3, "    ██  ", fg, bg)
}

func drawBishop(buf *pencere.Buffer, fg, bg pencere.Color) {
	buf.SetString(0, 1, "    ▟▙  ", fg, bg)
	buf.SetString(0, 2, "    ▜▛  ", fg, bg)
	buf.SetString(0, 3, "    ▟▙  ", fg, bg)

}

func drawHorse(buf *pencere.Buffer, fg, bg pencere.Color) {

	buf.SetString(0, 0, "    ▁▂▂▟ ", fg, bg)
	buf.SetString(0, 1, "  ▟█████▙  ", fg, bg)
	buf.SetString(0, 2, "  ███▘ ▜█ ", fg, bg)
	buf.SetString(0, 3, " ▟████▙      ", fg, bg)

}

func drawCastle(buf *pencere.Buffer, fg pencere.Color, bg pencere.Color) {

	buf.SetString(0, 0, "  ▄ ▄ ▄ ", fg, bg)
	buf.SetString(0, 1, " ▜█████▛  ", fg, bg)
	buf.SetString(0, 2, "  █████ ", fg, bg)
	buf.SetString(0, 3, " ▟█████▙    ", fg, bg)

}

func drawKing(buf *pencere.Buffer, fg, bg pencere.Color) {
	buf.SetString(0, 0, "    ╬   ", fg, bg)
	buf.SetString(0, 1, " ▟█▙ ▟█▙  ", fg, bg)
	buf.SetString(0, 2, " ▜█████▛ ", fg, bg)
	buf.SetString(0, 3, "  ▜███▛      ", fg, bg)

}

func drawQueen(buf *pencere.Buffer, fg, bg pencere.Color) {
	buf.SetString(0, 0, " ▗     ▖", fg, bg)
	buf.SetString(0, 1, "▚ ▚ █ ▞ ▞", fg, bg)
	buf.SetString(0, 2, " ▜█████▛ ", fg, bg)
	buf.SetString(0, 3, "  ▜███▛      ", fg, bg)

}

func newBoard() (*Board, error) {
	p, err := pencere.NewBox(pencere.Position(0, 0, 94, 48))
	if err != nil {
		return nil, err
	}

	p.BorderFg = 255
	board := &Board{
		p:                 p,
		sequareMap:        make(map[string]*Square),
		sequareWidth:      10,
		sequareHeight:     5,
		leftMargin:        6,
		topMargin:         4,
		BlackSequareColor: 140,
		WhiteSequareColor: 23,
		WhitePieceColor:   255,
		BlackPieceColor:   0,
	}

	board.draggingSquare = NewSequare(board, pencere.Position(0, 0, board.sequareWidth, board.sequareHeight))
	board.AddControl(board.draggingSquare)

	dp := board.draggingSquare.Pencere()
	dp.Visible = false
	dp.ZIndex = 100

	for x := 0; x < 8; x++ {
		for y := 0; y < 8; y++ {
			s := NewSequare(board, pencere.Position(board.leftMargin+x*board.sequareWidth, board.topMargin+y*board.sequareHeight, board.sequareWidth, board.sequareHeight))
			board.sequares = append(board.sequares, s)

			board.sequareMap[string(letters[x])+fmt.Sprintf("%v", 7-y+1)] = s
			board.AddControl(s)

			if (x+y)%2 == 0 {
				s.Pencere().Bg = board.WhiteSequareColor
			} else {
				s.Pencere().Bg = board.BlackSequareColor
			}
		}
	}

	for x := 0; x < 8; x++ {
		board.AddPencere(board.newBoardLabel(5+x*board.sequareWidth+(board.sequareWidth/2), 2, string(letters[x])))
		board.AddPencere(board.newBoardLabel(5+x*board.sequareWidth+(board.sequareWidth/2), 45, string(letters[x])))

	}
	for y := 0; y < 8; y++ {
		board.AddPencere(board.newBoardLabel(3, 4+y*board.sequareHeight+(board.sequareHeight/2), string(numbers[7-y])))
		board.AddPencere(board.newBoardLabel(89, 4+y*board.sequareHeight+(board.sequareHeight/2), string(numbers[7-y])))

	}

	return board, nil
}

type Board struct {
	p             *pencere.Pencere
	sequareWidth  int
	sequareHeight int
	leftMargin    int
	topMargin     int

	draggingSquare    *Square
	sequares          []*Square
	sequareMap        map[string]*Square
	BlackSequareColor pencere.Color
	WhiteSequareColor pencere.Color

	BlackPieceColor pencere.Color
	WhitePieceColor pencere.Color

	labels []*pencere.Pencere
}

func (this *Board) Pencere() *pencere.Pencere {
	return this.p
}

func (this *Board) AddPencere(p *pencere.Pencere) {
	this.p.AddPencere(p)
}

func (this *Board) AddControl(c pencere.Control) {
	this.p.AddControl(c)
}

func (this *Board) SetPiece(name string, pieceType string, color pencere.Color) {
	if s, ok := this.sequareMap[name]; ok {
		s.PieceType = pieceType
		s.PieceColor = color
	}
}

func (this *Board) newBoardLabel(x, y int, text string) *pencere.Pencere {
	p, err := pencere.NewPencere(pencere.Position(x, y, 1, 1))
	p.Text = text
	if err != nil {
		panic(err)
	}

	p.Render = func(buf *pencere.Buffer) error {
		buf.SetString(0, 0, p.Text, p.Fg, p.Bg)
		return nil
	}

	return p
}

func NewSequare(board *Board, options ...pencere.Option) *Square {
	p, err := pencere.NewPencere(options...)
	p.Bg = recColor
	p.Fg = 3
	_ = err
	p.HasBorder = false
	p.CanFocus = true

	s := &Square{
		p: p,
	}

	p.OnFocus = func() error {
		p.SetValue("Bg", int(p.Bg))
		p.Bg = focusColor
		return nil
	}

	p.OnLostFocus = func() error {
		p.Bg = pencere.Color(p.GetInt("Bg", int(p.Bg)))
		return nil
	}

	p.Render = s.Render

	p.OnDragBegin = func(event pencere.MouseEvent) (bool, error) {
		p.SetValue("dragX", event.X)
		p.SetValue("dragY", event.Y)

		board.draggingSquare.PieceType = s.PieceType
		board.draggingSquare.PieceColor = s.PieceColor
		board.draggingSquare.Pencere().Fg = p.Fg
		board.draggingSquare.Pencere().Bg = p.Bg
		board.draggingSquare.Pencere().Left, board.draggingSquare.Pencere().Top = p.Left, p.Top
		board.draggingSquare.Pencere().Visible = true
		return true, nil
	}

	p.OnDragEnd = func(event pencere.MouseEvent) error {
		board.draggingSquare.Pencere().Visible = false
		return nil
	}

	p.OnDragging = func(event pencere.MouseEvent) error {

		dy := p.GetInt("dragY", 0)
		dx := p.GetInt("dragX", 0)
		x, y := p.Parent().TranslateToXY(event.GlobalX, event.GlobalY)

		board.draggingSquare.Pencere().Left, board.draggingSquare.Pencere().Top = x-dx, y-dy

		return nil
	}

	return s
}

type Square struct {
	PieceType   string
	PieceColor  pencere.Color
	p           *pencere.Pencere
	Selected    bool
	HighLighted bool
}

func (this *Square) Pencere() *pencere.Pencere {
	return this.p
}

func (this *Square) AddPencere(p *pencere.Pencere) {
	this.p.AddPencere(p)
}

func (this *Square) AddControl(c pencere.Control) {
	this.p.AddControl(c)
}

func (this *Square) Render(buf *pencere.Buffer) error {

	switch this.PieceType {
	case "pawn":
		drawPawn(buf, this.PieceColor, this.p.Bg)
	case "bishop":
		drawBishop(buf, this.PieceColor, this.p.Bg)
	case "castle":
		drawCastle(buf, this.PieceColor, this.p.Bg)
	case "horse":
		drawHorse(buf, this.PieceColor, this.p.Bg)
	case "king":
		drawKing(buf, this.PieceColor, this.p.Bg)
	case "queen":
		drawQueen(buf, this.PieceColor, this.p.Bg)
	}
	return nil
}
