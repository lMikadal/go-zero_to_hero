package main

import "fmt"

func main() {
	point := 0

	if point == 0 {
		fmt.Println("Zero")
	} else if point%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}
