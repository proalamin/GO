package main

import "fmt"

func main(){
	yourNumber :=88

	if yourNumber > 79{
		fmt.Println("You got A")
	}else if yourNumber < 79 && yourNumber > 32{
		fmt.Println("You are pass")
	}else{
		fmt.Println("You are fail")
	}

}