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

	board.SetPiece("a1", "castle", "white")
	board.SetPiece("b1", "horse", "white")
	board.SetPiece("c1", "bishop", "white")
	board.SetPiece("d1", "king", "white")
	board.SetPiece("e1", "queen", "white")
	board.SetPiece("f1", "bishop", "white")
	board.SetPiece("g1", "horse", "white")
	board.SetPiece("h1", "castle", "white")

	board.SetPiece("a2", "pawn", "white")
	board.SetPiece("b2", "pawn", "white")
	board.SetPiece("c2", "pawn", "white")
	board.SetPiece("d2", "pawn", "white")
	board.SetPiece("e2", "pawn", "white")
	board.SetPiece("f2", "pawn", "white")
	board.SetPiece("g2", "pawn", "white")
	board.SetPiece("h2", "pawn", "white")

	board.SetPiece("a8", "castle", "black")
	board.SetPiece("b8", "horse", "black")
	board.SetPiece("c8", "bishop", "black")
	board.SetPiece("d8", "king", "black")
	board.SetPiece("e8", "queen", "black")
	board.SetPiece("f8", "bishop", "black")
	board.SetPiece("g8", "horse", "black")
	board.SetPiece("h8", "castle", "black")

	board.SetPiece("a7", "pawn", "black")
	board.SetPiece("b7", "pawn", "black")
	board.SetPiece("c7", "pawn", "black")
	board.SetPiece("d7", "pawn", "black")
	board.SetPiece("e7", "pawn", "black")
	board.SetPiece("f7", "pawn", "black")
	board.SetPiece("g7", "pawn", "black")
	board.SetPiece("h7", "pawn", "black")

	pencere.Loop()

	pencere.Close()
	return nil
}

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

	//draggingSquare    *Square
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

func (this *Board) SetPiece(name string, pieceType string, pieceColor string) {
	if s, ok := this.sequareMap[name]; ok {
		s.PieceType = pieceType
		s.PieceColor = pieceColor
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
		p.Properties.SetValue("Bg", int(p.Bg))
		p.Bg = focusColor
		return nil
	}

	p.OnLostFocus = func() error {
		p.Bg = pencere.Color(p.Properties.GetInt("Bg", int(p.Bg)))
		return nil
	}

	p.Render = s.Render

	p.OnDragBegin = func(event pencere.DragBeginEvent) (bool, *pencere.DragContext, error) {

		if s.PieceType == "" {
			return false, nil, nil
		}
		dragContext := pencere.NewDragContext(p)
		dragContext.IconOffsetX = -event.X
		dragContext.IconOffsetY = -event.Y

		draggingSquare := NewSequare(board, pencere.Position(0, 0, board.sequareWidth, board.sequareHeight))

		draggingSquare.PieceType = s.PieceType
		draggingSquare.PieceColor = s.PieceColor
		draggingSquare.Pencere().Fg = p.Fg
		draggingSquare.Pencere().Bg = p.Bg

		dragContext.DraggingIcon = draggingSquare.Pencere()

		dragContext.Properties.SetString("pieceType", s.PieceType)
		dragContext.Properties.SetString("pieceColor", s.PieceColor)

		return true, dragContext, nil
	}

	p.OnDragEnd = func(event pencere.DragEndEvent) error {

		if event.DropedPencere != nil {
			s.PieceType = ""
			s.PieceColor = ""
			pencere.SetFocus(event.DropedPencere)
		}

		return nil
	}

	p.OnDragging = func(event pencere.DraggingEvent) error {

		// dy := p.Properties.GetInt("dragY", 0)
		// dx := p.Properties.GetInt("dragX", 0)
		// x, y := p.Parent().TranslateToXY(event.GlobalX, event.GlobalY)

		// board.draggingSquare.Pencere().Left, board.draggingSquare.Pencere().Top = x-dx, y-dy

		return nil
	}

	p.OnDragEnter = func(event pencere.DragEnterEvent) error {

		p.Properties.SetInt("dragBg", int(p.Bg))
		p.Bg = 145
		return nil
	}

	p.OnDragLeave = func(event pencere.DragLeaveEvent) error {

		p.Bg = pencere.Color(p.Properties.GetInt("dragBg", int(p.Bg)))
		return nil
	}

	p.OnDrop = func(event pencere.DropEvent) error {

		s.PieceType = event.DragContext.Properties.GetString("pieceType", "")
		s.PieceColor = event.DragContext.Properties.GetString("pieceColor", "")

		return nil
	}

	p.CanDrop = func(event pencere.DropEvent) (bool, error) {
		if event.DragContext.Properties.GetString("pieceColor", "") == s.PieceColor {
			return false, nil
		}

		return true, nil
	}

	return s
}

type Square struct {
	PieceType   string
	PieceColor  string
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
	var pg pencere.Color
	if this.PieceColor == "white" {
		pg = 255
	} else {
		pg = 0

	}
	switch this.PieceType {
	case "pawn":
		drawPawn(buf, pg, this.p.Bg)
	case "bishop":
		drawBishop(buf, pg, this.p.Bg)
	case "castle":
		drawCastle(buf, pg, this.p.Bg)
	case "horse":
		drawHorse(buf, pg, this.p.Bg)
	case "king":
		drawKing(buf, pg, this.p.Bg)
	case "queen":
		drawQueen(buf, pg, this.p.Bg)
	}
	return nil
}
