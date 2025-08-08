package variables

/** Go By Example
 *  Exercise 3 - Variables
 */

import "fmt"

func Variables() {

	// Declare variable with var
	var a = "initial"
	fmt.Println(a)

	// Declare multiple variables at a time
	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	// := declares and initializes the variable at once
	f := "apple"
	fmt.Println(f)

	num := 52
	fmt.Println(num)

	var num2 = 53
	fmt.Println(num2)
}
