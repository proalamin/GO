package main

import "fmt"

var (
	a = 10
)

func main(){
	age := 30  //when a variable declared in a nested scope has the same name as a variable declared in an outer scope

	if age > 18{
		a:= 47
		fmt.Println(a)
	}
	fmt.Println(a)

}
