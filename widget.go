package pencere

type Control interface {
	Window() *Pencere
	AddWindow(c Control)
}

type Widget interface {
	Control
}
