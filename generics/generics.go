package generics

type Interface interface {
}

// Generics without type constraint
type Generic[T any] struct {
	value T
}

// Generics with type constraint
type GenericInterface[T Interface] struct {
	value T
}
