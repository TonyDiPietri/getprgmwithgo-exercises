//Listing 6.5: Floating-point inaccuracies
package main

import (
	"fmt"
)

func main() {
	third := 1.0 / 3
	fmt.Println(third + third + third)

	piggyBank := 0.1
	piggyBank += 0.2
	fmt.Println(piggyBank)
}
