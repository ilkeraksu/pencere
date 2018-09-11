package pencere

func NewBox(options ...Option) (*Pencere, error) {
	p, err := NewPencere(options...)
	if err != nil {
		return nil, err
	}
	p.HasBorder = true
	p.CanFocus = true
	return p, nil
}
