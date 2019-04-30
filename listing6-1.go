//Listing 6.1: 64-bit vs. 32-bit floating point
package main

import (
	"fmt"
	"math"
)

func main() {
	var pi64 = math.Pi
	var pi32 float32 = math.Pi

	fmt.Println(pi64)
	fmt.Println(pi32)
}
