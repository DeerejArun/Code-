package main

import (
	"fmt"
	"math"
)

func main() {
	var number int
	var result float64
	result = math.Cbrt(float64(number))
	fmt.Print("Enter a number : ")
	fmt.Scanln(&number)
	fmt.Printf("The cuberoot of a number is %f", result)
}
