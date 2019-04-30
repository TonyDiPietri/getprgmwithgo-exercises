//Listing 6.4: Zero padding
package main

import (
	"fmt"
)

func main() {
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%05.2f\n", third)
}
