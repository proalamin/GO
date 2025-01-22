package main

import (
	"fmt"
	"example.com/mathlb"
)

var a = 10
var b = 20

func main() {
	fmt.Println("showing custom package")
	mathlb.Add(a,b)
	mathlb.Sum()
}