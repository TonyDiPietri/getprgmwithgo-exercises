//Listing 7.2: Integers wrap around
package main

import (
	"fmt"
)

func main() {
	var red uint8 = 255
	red++
	fmt.Println(red)
	
	var number int8 = 127
	number++
	fmt.Println(number)
}
