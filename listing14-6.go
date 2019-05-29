//listing 14.6: Sensor calibration
package main

import "fmt"

type kelvin float64

//sensor function type; constrains 'sensor' to being a function that returns 'kelvin' type
type sensor func() kelvin

func realSensor() kelvin {
	return 0
}

func calibrate(s sensor, offset kelvin) sensor {
	//sensor expects a kelvin type return, so we create an anonymous function to cast the result as kelvin
	return func() kelvin {
		return s() + offset
	}
}

func main() {
	//assign sensor variable to call 'calibrate' function, passing in the results of the 'realSensor' function
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor())
}
