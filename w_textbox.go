package pencere

import (
	tb "github.com/nsf/termbox-go"
)

func NewMultiLineTextBox(options ...Option) (*MultiLineTextBox, error) {

	p, err := NewPencere(options...)
	if err != nil {
		return nil, err
	}
	p.HasBorder = false
	p.CanFocus = true

	textboxStyle := p.Theme.Style("textbox")
	p.Bg = textboxStyle.Bg
	p.Fg = textboxStyle.Fg

	textBox := &MultiLineTextBox{
		Pencere: p,
		textBuffer: &RuneBuffer{
			wordwrap: true,
		},
	}

	p.Render = textBox.Render
	p.OnKeyEvent = textBox.OnKeyEvent

	return textBox, nil
}

type MultiLineTextBox struct {
	*Pencere
	textBuffer *RuneBuffer
	offset     int
}

func (this *MultiLineTextBox) Render(buf *Buffer) error {
	this.textBuffer.SetMaxWidth(this.Width - 2)

	lines := this.textBuffer.SplitByLine()
	for i, line := range lines {
		//p.FillRect(0, i, s.X, 1)
		//p.DrawText(0, i, line)
		buf.SetString(1, i+1, line, this.Fg, this.Bg)
	}

	if this.HasFocus {
		pos := this.textBuffer.CursorPos()
		c := buf.At(pos.X+1, pos.Y+1)

		c.Fg = 0
		c.Bg = 123
		//c.Ch = 'X'

		buf.SetCell(pos.X+1, pos.Y+1, c)
		//buf.SetCell(pos.X+1, pos.Y+1, Cell{Ch: '|', Fg: p.Fg, Bg: p.Bg})
		//p.DrawCursor(pos.X, pos.Y)
	}

	return nil
}

func (this *MultiLineTextBox) OnKeyEvent(ev KeyEvent) error {
	if !this.HasFocus {
		return nil
	}

	screenWidth := this.Width

	this.textBuffer.SetMaxWidth(screenWidth)

	if ev.Rune == 0 {
		switch ev.Key {
		case tb.KeyEnter:
			this.textBuffer.WriteRune('\n')
		case tb.KeyBackspace:
			fallthrough
		case tb.KeyBackspace2:
			this.textBuffer.Backspace()

			isTextRemaining := this.textBuffer.Width()-this.offset > this.Width
			if this.offset > 0 && !isTextRemaining {
				this.offset--

			}
			this.FireEvent("text_changed", nil)
		case tb.KeyDelete, tb.KeyCtrlD:
			this.textBuffer.Delete()
			this.FireEvent("text_changed", nil)
		//case KeyLeft, tb.KeyCtrlB:
		case tb.KeyArrowLeft, tb.KeyCtrlB:
			this.textBuffer.MoveBackward()
			if this.offset > 0 {
				this.offset--

			}
		//case tb.KeyRight, tb.KeyCtrlF:
		case tb.KeyArrowRight, tb.KeyCtrlF:
			this.textBuffer.MoveForward()

			isCursorTooFar := this.textBuffer.CursorPos().X >= screenWidth
			isTextLeft := (this.textBuffer.Width() - this.offset) > (screenWidth - 1)

			if isCursorTooFar && isTextLeft {
				this.offset++

			}
		case tb.KeyHome, tb.KeyCtrlA:
			this.textBuffer.MoveToLineStart()
			this.offset = 0

		case tb.KeyEnd, tb.KeyCtrlE:
			this.textBuffer.MoveToLineEnd()
			left := this.textBuffer.Width() - (screenWidth - 1)
			if left >= 0 {
				this.offset = left

			}
		case tb.KeyCtrlK:
			this.textBuffer.Kill()
		case tb.KeySpace:
			ev.Rune = ' '
			goto yaz
		}
		this.ContentY = this.textBuffer.heightForWidth(this.Width)
		return nil
	}

yaz:
	this.textBuffer.WriteRune(ev.Rune)
	if this.textBuffer.CursorPos().X >= screenWidth {
		this.offset++

	}
	this.FireEvent("text_changed", nil)
	this.ContentY = this.textBuffer.heightForWidth(this.Width)
	return nil
}
