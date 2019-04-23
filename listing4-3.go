//Listing 4.3: Short declaration in a for loop
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}
}
