/*CAPSTONE: TICKET TO MARS
Welcome to the first challenge. It’s time to take everything covered in unit 1 and write
a program on your own. Your challenge is to write a ticket generator in the Go Playground
that makes use of variables, constants, switch, if, and for. It should also draw
on the fmt and math/rand packages to display and align text and to generate random
numbers.
Start by building a prototype that generates 10 random
tickets and displays them in a tabular format
with a nice header, as follows:
(TABLE REMOVED)
The table should have four columns:
 The spaceline company providing the service
 The duration in days for the trip to Mars (one-way)
 Whether the price covers a return trip
 The price in millions of dollars 
For each ticket, randomly select one of the following spacelines: Space Adventures,
SpaceX, or Virgin Galactic.
Use October 13, 2020 as the departure date for all tickets. Mars will be 62,100,000 km
away from Earth at the time.
Randomly choose the speed the ship will travel, from 16 to 30 km/s. This will determine
the duration for the trip to Mars and also the ticket price. Make faster ships more expensive,
ranging in price from $36 million to $50 million. Double the price for round trips.
*/
package main

import (
	"fmt"
	"math/rand"
)

const distance = 62100000

func main() {

	var company, days, returnyn, price int

	fmt.Printf("%-16v %v %v %v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Printf("=====================================\n")

	for count := 0; count < 10; count++ {

		var compname string

		//randomly choose a company
		company = rand.Intn(3)

		//Set string 'compname' based on outcome of random int 'company'
		if company == 0 {
			compname = "Space Adventures"
		} else if company == 1 {
			compname = "SpaceX"
		} else if company == 2 {
			compname = "Virgin Galactic"
		}

		//Determine days for each ticket based on random speed and 'distance'
		speed := (rand.Intn(14) + 16)
		days = (distance / (speed * 86400)) //86400 converts 16-30 km/s into km/day
		price = 36 + (speed - 16) //starts price at 36 and increases as speed rises

		fmt.Printf("%-16v %4v %9v $%2v\n", compname, days, returnyn, price)

	}

}
