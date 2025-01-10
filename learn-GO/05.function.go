package main
import "fmt"

func add(num1 int, num2 int) int {
	sum:= num1+ num2;

	return sum;
}

func printMust(){
	fmt.Println("education must be free")
}

func sayHello(name string){
	fmt.Println("welcome to the real world,",name)
}

func main(){
	a:= 130;
	b:= 20;

	ans := add(a, b)
	fmt.Println(ans)

	printMust()

	sayHello("al amin")
}