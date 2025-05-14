package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time to study Golang")

	prensetTime := time.Now()
	// fmt.Print(prensetTime)
	fmt.Println(prensetTime.Format("01-02-2006 Monday 15:04:05"))

	createDate := time.Date(2002, time.September, 10, 11, 55, 13, 0, time.UTC)
	fmt.Println(createDate)
}
