package pencere

type Control interface {
	Pencere() *Pencere
	AddPencere(*Pencere)
	AddControl(c Control)
}

type Widget interface {
	Control
}
