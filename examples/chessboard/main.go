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

	pencere.Root().AddWindow(board)

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
		Pencere:           p,
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
			s := NewSequare(pencere.Position(board.leftMargin+x*board.sequareWidth, board.topMargin+y*board.sequareHeight, board.sequareWidth, board.sequareHeight))
			board.sequares = append(board.sequares, s)

			board.sequareMap[string(letters[x])+fmt.Sprintf("%v", 7-y+1)] = s
			board.AddWindow(s)

			if (x+y)%2 == 0 {
				s.Bg = board.WhiteSequareColor
			} else {
				s.Bg = board.BlackSequareColor
			}
		}
	}

	for x := 0; x < 8; x++ {
		board.AddWindow(board.newBoardLabel(5+x*board.sequareWidth+(board.sequareWidth/2), 2, string(letters[x])))
		board.AddWindow(board.newBoardLabel(5+x*board.sequareWidth+(board.sequareWidth/2), 45, string(letters[x])))

	}
	for y := 0; y < 8; y++ {
		board.AddWindow(board.newBoardLabel(3, 4+y*board.sequareHeight+(board.sequareHeight/2), string(numbers[7-y])))
		board.AddWindow(board.newBoardLabel(89, 4+y*board.sequareHeight+(board.sequareHeight/2), string(numbers[7-y])))

	}

	return board, nil
}

type Board struct {
	*pencere.Pencere
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

func NewSequare(options ...pencere.Option) *Square {
	p, err := pencere.NewPencere(options...)
	p.Bg = recColor
	p.Fg = 3
	_ = err
	p.HasBorder = false
	p.CanFocus = true

	s := &Square{
		Pencere: p,
	}

	p.OnFocus = s.OnFocus

	p.OnLostFocus = s.OnLostFocus

	p.Render = s.Render
	p.OnDragBegin = s.OnDragBegin
	p.OnDragEnd = s.OnDragEnd
	p.OnDragEnter = s.OnDragEnter
	p.OnDragLeave = s.OnDragLeave
	p.OnDrop = s.OnDrop
	p.CanDrop = s.CanDrop

	return s
}

type Square struct {
	*pencere.Pencere
	PieceType  string
	PieceColor string

	Selected    bool
	HighLighted bool
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
		drawPawn(buf, pg, this.Bg)
	case "bishop":
		drawBishop(buf, pg, this.Bg)
	case "castle":
		drawCastle(buf, pg, this.Bg)
	case "horse":
		drawHorse(buf, pg, this.Bg)
	case "king":
		drawKing(buf, pg, this.Bg)
	case "queen":
		drawQueen(buf, pg, this.Bg)
	}
	return nil
}

func (this *Square) CanDrop(event pencere.DropEvent) (bool, error) {
	if event.DragContext.Properties.GetString("pieceColor", "") == this.PieceColor {
		return false, nil
	}

	return true, nil
}

func (this *Square) OnDrop(event pencere.DropEvent) error {
	this.PieceType = event.DragContext.Properties.GetString("pieceType", "")
	this.PieceColor = event.DragContext.Properties.GetString("pieceColor", "")

	return nil
}

func (this *Square) OnDragLeave(event pencere.DragLeaveEvent) error {
	this.Bg = pencere.Color(this.Properties.GetInt("dragBg", int(this.Bg)))
	return nil
}

func (this *Square) OnDragEnter(event pencere.DragEnterEvent) error {
	this.Properties.SetInt("dragBg", int(this.Bg))
	this.Bg = 145
	return nil
}

func (this *Square) OnDragEnd(event pencere.DragEndEvent) error {
	if event.DropedPencere != nil {
		this.PieceType = ""
		this.PieceColor = ""
		pencere.SetFocus(event.DropedPencere)
	}

	return nil
}

func (this *Square) OnDragBegin(event pencere.DragBeginEvent) (bool, *pencere.DragContext, error) {
	if this.PieceType == "" {
		return false, nil, nil
	}
	dragContext := pencere.NewDragContext(this.Pencere)
	dragContext.IconOffsetX = -event.X
	dragContext.IconOffsetY = -event.Y

	draggingSquare := NewSequare(pencere.Position(0, 0, this.Width, this.Height))

	draggingSquare.PieceType = this.PieceType
	draggingSquare.PieceColor = this.PieceColor
	draggingSquare.Fg = this.Fg
	draggingSquare.Bg = this.Bg

	dragContext.DraggingIcon = draggingSquare.Window()

	dragContext.Properties.SetString("pieceType", this.PieceType)
	dragContext.Properties.SetString("pieceColor", this.PieceColor)

	return true, dragContext, nil
}

func (this *Square) OnFocus() error {
	this.Properties.SetValue("Bg", int(this.Bg))
	this.Bg = focusColor
	return nil
}

func (this *Square) OnLostFocus() error {
	this.Bg = pencere.Color(this.Properties.GetInt("Bg", int(this.Bg)))
	return nil
}
