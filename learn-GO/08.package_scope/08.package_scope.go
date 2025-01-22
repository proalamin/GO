package main

import (
	"fmt"
	"example.com/mathlb"
)

/*
---- package step create step -----

1. create a package folder in same directory
2. create a mod file
	command -> go mod init example.com

*/


var a = 10
var b = 20

func main() {
	fmt.Println("showing custom package")
	mathlb.Add(a,b)
	mathlb.Sum()
}