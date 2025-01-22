package main

import "fmt"

var a = 10
var b = 30
/*
	block -> { }
*/

func main() {
	x := 23

	if x >= 21 {
		p := 4000
		fmt.Println("You are eligible for driving test")
		fmt.Println("and you need", p, "tk for test fee")
	}
}
