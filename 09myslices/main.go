package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("The typowe of fruitlist is: %T \n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println("fruitList =", fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println("fruitList =", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("fruitList =", fruitList)

	highScore := make([]int, 4)
	highScore[0] = 923
	highScore[1] = 223
	highScore[2] = 323
	highScore[3] = 423
	fmt.Println("Is highScore sorted:", sort.IntsAreSorted(highScore))
	fmt.Println("highScore:", highScore)

	highScore = append(highScore, 1, 2, 3)
	fmt.Println("highScore:", highScore)

	sort.Ints(highScore)
	fmt.Println("Sorted highScore =", highScore)

	fmt.Println("Is highScore sorted:", sort.IntsAreSorted(highScore))

}
