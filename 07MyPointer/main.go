package main

import "fmt"

func main() {
	fmt.Println("Pointers in go")

	// var one int = 2

	var ptr *int

	fmt.Println("value of pointer is ", &ptr)
}
