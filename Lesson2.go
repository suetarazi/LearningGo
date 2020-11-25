package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	//the var keyword is used to define a variable
	var name string = "Sue Tarazi"
	var pet string
	pet = "Pinot Cat"
	fmt.Println(name)
	fmt.Println(pet)

	//explicit vs implicit variable declaration
	//var number uint16 = 260 //this is explicitly defined because we gave it the type uint16
	var number2 = 260 //this is implicitly defined because we are letting Go decide what type it needs to be under the hood
	fmt.Printf("%T", number2)
}
