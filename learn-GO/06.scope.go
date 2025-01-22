package main

import "fmt"

var a =10;
var b=30

func add(x int, y int){
	z:= x+y
	fmt.Println(z)
	// return z
}

func main(){
	p:= 30
	q:= 40

	add(p, q) // 70
	add(a, b) // 40

	add(a, q) // 50
}