package main

import "fmt"

func main() {
	s := 1
	switch s {
	case 1, 2:
		fmt.Println("it is 1 or 2.")
	case 3:
		fmt.Println("it is 3")
	default:
		fmt.Println("it is nothing")

	}
}
