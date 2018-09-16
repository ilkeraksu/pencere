package pencere

// Decoration represents a bold/underline/etc. state
type Decoration int

// Decoration modes: Inherit from parent widget, explicitly on, or explicitly off.
const (
	DecorationInherit Decoration = iota
	DecorationOn
	DecorationOff
)

type Style struct {
	Fg Color
	Bg Color

	Texture   rune
	Brush     rune
	Reverse   Decoration
	Bold      Decoration
	Underline Decoration
}

type Theme struct {
	styles map[string]Style
}

// DefaultTheme is a theme with reasonable defaults.
var DefaultTheme = &Theme{
	styles: map[string]Style{
		"default":             {Fg: 100, Bg: 230},
		"border":              {Fg: 230, Bg: 0},
		"label":               {Fg: 230, Bg: 0},
		"menubar":             {Fg: 12, Bg: 123},
		"scrollbar":           {Fg: 12, Bg: 123, Texture: '░'},
		"textbox":             {Fg: 12, Bg: 50, Texture: '░'},
		"list.item.selected":  {Reverse: DecorationOn},
		"table.cell.selected": {Reverse: DecorationOn},
		"button.focused":      {Reverse: DecorationOn},
	},
}

// NewTheme return an empty theme.
func NewTheme() *Theme {
	return &Theme{
		styles: make(map[string]Style),
	}
}

// SetStyle sets a style for a given identifier.
func (p *Theme) SetStyle(n string, i Style) {
	p.styles[n] = i
}

// Style returns the style associated with an identifier.
// If there is no Style associated with the name, it returns a default Style.
func (p *Theme) Style(name string) Style {
	return p.styles[name]
}

// HasStyle returns whether an identifier is associated with an identifier.
func (p *Theme) HasStyle(name string) bool {
	_, ok := p.styles[name]
	return ok
}
