package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	// define a variable
	var myName string = "John Doe"

	// define a constant
	const mySecondName string = "Jane Doe"
	// shorthand
	myThirdName := "Bob Doe"

	fmt.Println(myName)
	fmt.Println(mySecondName)
	fmt.Println(myThirdName)
}
