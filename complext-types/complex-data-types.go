package complext_types

// Struct's are  only complex data types in golang
// In Golang structs can only hold data, they cannot have methods
// But we can register methods on it
type Person struct {
	name string
}

// This is a method on Person struct
func (p *Person) GetName() string {
	return p.name
}
