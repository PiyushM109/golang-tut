package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcom to UserInput tutorial"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name :")

	input, _ := reader.ReadString('\n')
	fmt.Println(input)

}
