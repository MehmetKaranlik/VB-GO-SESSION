package person

import "net/http"

type PersonInterface[X http.Handler] interface {
	GetName() X
}

type Person struct {
	name string
}

func (p Person) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (p Person) SetName(name string) {
	//TODO implement me
	panic("implement me")
}
