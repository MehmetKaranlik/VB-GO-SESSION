package control_flow

import "fmt"

func ControlFlow() {
	boolean := true

	// If condition check
	if boolean {
		// doSomething
	} else if !boolean {
		// doSomething else
	} else {
		// doSomething else
	}

	array := []int{1, 2, 3, 4, 5}

	// For loop
	for index, value := range array {
		fmt.Println(index, value)
	}

	// Switch case
	switch boolean {
	case true:
	// doSomething
	case false:
	// doSomething
	default:
		// doSomething else
	}

	// Infinite loop, blocks thread(execution)
	for {
		fmt.Println("Hello World")
	}
	ch := make(chan int)

	// Select statement blocks thread(execution), but only works if there is an input on what its listening
	select {
	case x := <-ch:
		fmt.Println(x)
	}
}
