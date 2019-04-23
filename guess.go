package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//generate random number as correct answer (answer)
	var answer int
	answer = 98
	fmt.Println(answer)

	//generate random number (guess)
	var guess int
	guess = rand.Intn(99)
	fmt.Println("My first guess is: ", guess)

	//create loop with 3 possibilities:
	for guess != answer {
		switch {
		//case 1: number is too big, print "%v is too high", modify guess
		case guess > answer:
			fmt.Println(guess, "is too high.")
			guess--
		//case 2: number is too small, print "%v is too low", modify guess
		case guess < answer:
			fmt.Println(guess, " is too low.")
			guess++
		//case 3: number is correct, print "I LOVE IT" and exit
		//case guess == answer:
			//fmt.Println("I LOVE IT")
		}
	}
	fmt.Println("I LOVE IT. Thanks, ", guess, "!")
}
