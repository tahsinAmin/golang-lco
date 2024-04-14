package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on ponters")

	// var ptr *int
	// fmt.Println("Value of pointert is", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is", ptr)  // Holds the reference of the pointer of myNumber
	fmt.Println("Value of actual pointer is", *ptr) // I want to see what's inside the pointer

	*ptr += 2
	fmt.Println("New Value is:", myNumber) // This shows that it is being updated in the values and not on the copies of the value

}
