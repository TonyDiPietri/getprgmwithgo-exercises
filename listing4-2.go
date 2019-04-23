//Listing 4.2: Condensed countdown
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var count = 0

	for count = 10; count > 0; count-- {
		fmt.Println(count)
	}
	fmt.Println(count)
}