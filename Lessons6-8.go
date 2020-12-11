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

	conditions_and_booleans()
	chained_conditionals()
}

//conditions & boolean expressions
func conditions_and_booleans() {
	x := "sue"
	y := "Sue"
	val := x == y
	fmt.Printf("%t", val) //will compare the two strings and return 'false' because sue != Sue

}

func chained_conditionals() {
	//! = not
	//|| = or
	//&& = and

	val := (true || false) && true
	fmt.Printf("%t", val)
}