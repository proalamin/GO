/*
Arithmetic Operator-

Operators				Example
+ (Addition)			a + b
- (Subtraction)			a - b
*  (Multiplication)		a * b
/ (Division)			a / b
% (Modulo Division)		a % b


Relational Operators-

Operator						Example		Descriptions
== (equal to)					a == b		returns true if a and b are equal
!= (not equal to)				a != b		returns true if a and b are not equal
> (greater than)				a > b		returns true if a is greater than b
< (less than)					a < b		returns true if a is less than b
>= (greater than or equal to)	a >= b		returns true if a is either greater than or equal to b
<= (less than or equal to)		a <= b		returns true is a is either less than or equal to b


Logical Operators-

Operator				Description			Example
&& (Logical AND)		exp1 && exp2		returns true if both expressions exp1 and exp2 are true
|| (Logical OR)			exp1 || exp2		returns true if any one of the expressions is true.
! (Logical NOT)			!exp				returns true if exp is false and returns false if exp is true.


Assignment Operators-

Operator						Example			Same as
+= (addition assignment)		a += b			a = a + b
-= (subtraction assignment)		a -= b			a = a - b
*= (multiplication assignment)	a *= b			a = a * b
/= (division assignment)		a /= b			a = a / b
%= (modulo assignment)			a %= b			a = a % b

*/

package main
import "fmt"

func main() {
	num1 := 6
	num2 := 2

	sum := num1 + num2
	fmt.Println(sum)

	difference := num1 - num2
	fmt.Println(difference)

	product := num1 * num2
	fmt.Println(product)

	quotient := num1 / num2
	fmt.Println(quotient)

	remainder := num1 % num2
  	fmt.Println(remainder)

}