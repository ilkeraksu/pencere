package pencere

func NewBox() *Pencere {
	p := NewPencere()
	p.HasBorder = true
	p.Render = ColumnRenderer(p)
	return p
}
