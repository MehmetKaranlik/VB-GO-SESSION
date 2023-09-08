package variables

import "fmt"

// Constant (immutable) variable declaration
const (
	ConnectionString = "XXXXXXX"
)

// Regular variable declaration
var (
	str     string  = "Hello World"
	number  int     = 10
	decimal float32 = 10.5
	boolean bool    = true
)

// This will give a compile-time error, shorthand variable declaration cannot be used on global scope
// hellow := "Hello World"

func Variables() {

	// Shorthand variable declaration
	str := "Hello World"
	fmt.Println(str)
}
