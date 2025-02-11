// iqs - interview question

package main

import "fmt"

/*
Higher order function or first class function
	any one of the following 3
		1. parameter -> function
		2. function return
		3. 1, 2 both
*/

// higher order function (parameter -> function)
func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}
func add(x, y int) {
	z := x + y
	fmt.Println(z)
}


// higher order function (parameter -> function return)
func call() func(x int, y int) {
	return sum
}
func sum(x, y int) {
	z := x + y
	fmt.Println(z)
}

func main() {
	processOperation(2, 3, add)
	ans := call()
	ans(4, 4)
}

/*
	1. parameter vs argument
	2. First order function
		i. standard function or named function
		ii. anonymous function
		iii. IIFE
		iv. function expression
	3. Higher order function or first class function
		{
			any one of the following 3
			1. parameter -> function
			2. function return
			3. 1, 2 both
		}
	4. callback function -> if any higher order function pass any function it's call 	callback function
	5. first class citizen -> (variable assign data)
	6. 


	functional paradigm -> haskel, racket,

	math -> logic (discrete mathematics)
	1. first order logic
	2. higher order logic

	### logic ###
	1. Object (people, car , animal etc)
	2. Property (color, student, teacher)
	3. Relation




*/
