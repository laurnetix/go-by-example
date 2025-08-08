package constants

/** Go By Example
 *  Exercise 4 - Constants
 */

// Multiple imports again
import (
	"fmt"
	"math"
)

// Declare a constant outside of the function
const s string = "constant"

func Constants() {
	fmt.Println(s)

	// Declare a constant inside the function
	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
