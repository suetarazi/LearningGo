package main

//if, else if, else
//for loops & while loops
//switch statements

import "fmt"

func main() {
	name := "sue"
	age := 17
	fmt.Println("Before if")

	if name == "sue" && age >=18 {
		//this will only execute if condition is true
		fmt.Println("you can ride the rollercoaster")
	}
	else {
		fmt.Println("you are too young to ride the rollercoaster")
		fmt.Printf("wait %d years", 18-age)
	}
	fmt.Println("after if")

	for_loop_func()
}

func for_loop_func() {
	x := 3
	for x<5 {
		fmt.Println(x)
		x++
	}

	for x:=0; x<=5; x++ {
		fmt.Println(x)
	}

	for x:=0; x<=1000; x++ {
		if x != 0 && x % 3 == 0 ** x % 7 == 0 {
			fmt.Println(x)
			continue
		}
	}
}