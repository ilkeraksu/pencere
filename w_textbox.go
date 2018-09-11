package pencere

import (
	tb "github.com/nsf/termbox-go"
)

func NewTextBox(options ...Option) (*Pencere, error) {

	p, err := NewPencere(options...)
	if err != nil {
		return nil, err
	}
	p.HasBorder = false
	p.CanFocus = true

	runeBuffer := &RuneBuffer{
		wordwrap: true,
	}
	p.SetValue("text", runeBuffer)
	p.SetValue("offset", 0)

	p.Render = func(buf *Buffer) error {
		//buf.SetString(1, 1, "DENEME 2", -1, -1)

		text := p.GetValue("text").(*RuneBuffer)
		text.SetMaxWidth(p.Width - 2)

		lines := text.SplitByLine()
		for i, line := range lines {
			//p.FillRect(0, i, s.X, 1)
			//p.DrawText(0, i, line)
			buf.SetString(1, i+1, line, p.Fg, p.Bg)
		}

		if p.HasFocus {
			pos := text.CursorPos()
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

	p.OnKeyEvent = func(ev KeyEvent) error {
		if !p.HasFocus {
			return nil
		}

		screenWidth := p.Width

		text := p.GetValue("text").(*RuneBuffer)
		text.SetMaxWidth(screenWidth)
		offset := p.GetValue("offset").(int)

		if ev.Rune == 0 {
			switch ev.Key {
			case tb.KeyEnter:
				text.WriteRune('\n')
			case tb.KeyBackspace:
				fallthrough
			case tb.KeyBackspace2:
				text.Backspace()

				isTextRemaining := text.Width()-offset > p.Width
				if offset > 0 && !isTextRemaining {
					offset--
					p.SetValue("offset", offset)
				}
				p.FireEvent("text_changed", nil)
			case tb.KeyDelete, tb.KeyCtrlD:
				text.Delete()
				p.FireEvent("text_changed", nil)
			//case KeyLeft, tb.KeyCtrlB:
			case tb.KeyArrowLeft, tb.KeyCtrlB:
				text.MoveBackward()
				if offset > 0 {
					offset--
					p.SetValue("offset", offset)
				}
			//case tb.KeyRight, tb.KeyCtrlF:
			case tb.KeyArrowRight, tb.KeyCtrlF:
				text.MoveForward()

				isCursorTooFar := text.CursorPos().X >= screenWidth
				isTextLeft := (text.Width() - offset) > (screenWidth - 1)

				if isCursorTooFar && isTextLeft {
					offset++
					p.SetValue("offset", offset)
				}
			case tb.KeyHome, tb.KeyCtrlA:
				text.MoveToLineStart()
				offset = 0
				p.SetValue("offset", offset)
			case tb.KeyEnd, tb.KeyCtrlE:
				text.MoveToLineEnd()
				left := text.Width() - (screenWidth - 1)
				if left >= 0 {
					offset = left
					p.SetValue("offset", offset)
				}
			case tb.KeyCtrlK:
				text.Kill()
			case tb.KeySpace:
				ev.Rune = ' '
				goto yaz
			}
			p.ContentY = text.heightForWidth(p.Width)
			return nil
		}

	yaz:
		text.WriteRune(ev.Rune)
		if text.CursorPos().X >= screenWidth {
			offset++
			p.SetValue("offset", offset)
		}
		p.FireEvent("text_changed", nil)
		p.ContentY = text.heightForWidth(p.Width)
		return nil
	}

	return p, nil
}
