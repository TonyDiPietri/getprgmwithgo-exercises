//Listing 7.3: Display the bits
package main

import (
	"fmt"
)

func main() {
	var green uint8 = 3
	fmt.Printf("%08b\n", green)
	green++
	fmt.Printf("%08b\n", green)
}
