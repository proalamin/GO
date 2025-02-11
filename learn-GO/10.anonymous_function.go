package main

import "fmt"

// standard or named function
func add(a, b int) {
	fmt.Println(a + b)
}

func main() {
	// we invoke the add function here
	add(2, 3)

	var a = 10 // expression
	fmt.Println(a)

	// anonymous function
	// Immediately Invoked(call) Function Expression
	// IIFE
	func(a, b int) {
		fmt.Println(a + b)
	}(5, 7) //immediately invoked(call) this function

	// function expression
	sum := func(x, y int) {
		fmt.Println((x + y))
	}

	sum(2, 3)

}

func init() {
	fmt.Println("learn go....")
}
