package pencere

var root *Pencere
var body *Pencere

func init() {
	root = newRoot()
}

func newRoot() *Pencere {
	p := NewPencere()

	//p.Render = ColumnRenderer(p)
	return p
}

func Root() *Pencere {
	return root
}
