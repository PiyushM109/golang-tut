package main

import "fmt"

func main() {
	fmt.Println("Slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Print(fruitList)

	fruitList = append(fruitList, "banana", "Mango")

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 264
	highScores[1] = 265
	highScores[2] = 266

}
