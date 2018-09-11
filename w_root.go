package pencere

var root *Pencere
var body *Pencere

func init() {
	root = newRoot()
}

func newRoot() *Pencere {
	p, err := NewPencere()
	if err != nil {
		panic(err)
	}

	//p.Render = ColumnRenderer(p)
	return p
}

func Root() *Pencere {
	return root
}
