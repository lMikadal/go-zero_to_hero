package main

import "fmt"

// global case 1
var global bool

// global case 2
var global_2 = true

func main() {
	// local case 1
	var x int
	x = 1

	// local case 2
	var y = 2

	// local case 3 (short declaration)
	z := 3

	// zero value
	var i int
	var s string
	var b bool

	// not use
	_ = global
	_ = global_2

	fmt.Printf("x: %v, y: %v, z: %v\n", x, y, z)
	fmt.Printf("global => 1: %v, 2: %v\n", global, global_2)
	fmt.Printf("zero value => int: %#v, string: %#v, bool: %#v\n", i, s, b)
}
