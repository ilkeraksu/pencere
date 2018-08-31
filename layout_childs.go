package pencere

type PencereOrder struct {
	Order   int
	Pencere *Pencere
}

func NextTo(p *Pencere, n *Pencere) func() error {

	return func() error {

		return nil
	}
}
