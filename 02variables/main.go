package main

import "fmt"

func main() {
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variables is a of type: %T \n", isLoggedIn)

	var smallFloat float32 = 255.455445411254461885
	fmt.Println(smallFloat)
	fmt.Printf("Variables is a of type: %T \n", smallFloat)

	var bigFloat float64 = 255.455445411254461885
	fmt.Println(bigFloat)
	fmt.Printf("Variables is a of type: %T \n", bigFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variables is a of type: %T \n", anotherVariable)

	// iomplicit typ

	var website = "learncodeonline.in"
	fmt.Println(website)

	// no var style

	numberOfUser := 3000.0
	fmt.Println(numberOfUser)

}
