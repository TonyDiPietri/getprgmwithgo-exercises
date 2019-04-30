//Listing 7.4: The bits when integers wrap
package main

import (
	"fmt"
)

func main() {
	var blue uint8 = 255
	fmt.Printf("%08b\n", blue)
	blue++
	fmt.Printf("%08b\n", blue)
}
