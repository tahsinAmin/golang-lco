package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is:", fruitList)

	var vegList = [5]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", len(vegList))
}
