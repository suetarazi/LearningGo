package main

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)

func main() {
	fmt.Printf("%T", 10)
	var x string = fmt.Sprintf("%T", 10)
	fmt.Printf(x)

	//taking in user input
	//1. make a scanner object
	scanner := bufio.NewScanner(os.Stdin)
	//2. take in input
	scanner.Scan()
	input := scanner.Text()

	fmt.Printf("You typed: %q", input)

}
