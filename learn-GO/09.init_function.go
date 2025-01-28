package main

import "fmt"


func main() {
	fmt.Println("init")
}
// The init function in Go is a special, automatically invoked function used for initializing packages or variables before the main function runs.
func init() {
	fmt.Println("I am the first function that is executed first")
}
