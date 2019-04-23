//Listing 3.4
package main

import "fmt"

func main() {

	var year = 2020
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	
	fmt.Printf("The year is %v, should you leap?\n", year)

	if leap {
		fmt.Println("Look before you leap!")
	} else {
		fmt.Println("Keep your feet on the ground.")
	}
}