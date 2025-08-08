package main

/* Go By Example */

// Comments in go are in C-Style
// and C++ style

// Import modules
import (
	/* Example 1 - Hello World! */
	"gobyexample/hello"

	/* Example 2 - Values */
	"gobyexample/values"

	/* Example 3 - Variables */
	"gobyexample/variables"

	/* Example 4 - Constants */
	"gobyexample/constants"
)

// Only one main function per project
func main() {

	/* Example 1 - Hello World! */
	hello.Hello()

	/* Example 2 - Values */
	values.Values()

	/* Example 3 - Variables */
	variables.Variables()

	/* Example 4 - Constants */
	constants.Constants()
}
