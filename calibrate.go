//Experiment for Lesson 14: calibrate.go
package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

//sensor function type; constrains 'sensor' to being a function that returns 'kelvin' type
type sensor func() kelvin

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func calibrate(s sensor, offset kelvin) sensor {
	//sensor expects a kelvin type return, so we create an anonymous function to cast the result as kelvin
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	//assign sensor variable to call 'calibrate' function, passing in the results of the 'realSensor' function
	var offset kelvin = 5
	sensor := calibrate(fakeSensor, offset)

	for i := 0; i <= 10; i++ {
		fmt.Println(sensor())
	}
}
