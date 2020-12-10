package main

//arithmetic operators & math

import (
	"fmt"
	//"math" //math package has sqrt, trig functions, etc.
)

func main() {
	var num1 float64 = 8
	var num2 int = 4
	answer := num1 / float64(num2)
	fmt.Printf("%d", answer)

}
