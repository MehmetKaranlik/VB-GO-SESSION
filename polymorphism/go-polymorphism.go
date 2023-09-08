package polymorphism

// In golang polymorphism achieved through registering same method on different types
type Poly interface {
	Conform()
}

type XImplementer struct{}

func (X XImplementer) Conform() {}

type YImplementer struct{}

func (Y YImplementer) Conform() {}

func ProcessPoly() {
	var poly Poly
	poly = XImplementer{}
	poly.Conform()
	poly = YImplementer{}
	poly.Conform()
}
