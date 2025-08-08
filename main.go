package main

/* Go By Example */

// Comments in go are in C-Style
// and C++ style

// Import modules
import (
	"gobyexample/hello"
	"gobyexample/values"
	"gobyexample/variables"
)

// Only one main function per project
func main() {
	hello.Hello()
	values.Values()
	variables.Variables()
}
