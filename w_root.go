package pencere

var root *Pencere = newRoot()

func newRoot() *Pencere {
	p := NewPencere()

	p.Render = ColumnRenderer(p)
	return p
}

func Root() *Pencere {
	return root
}
