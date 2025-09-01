package main

import "fmt"

func main() {

	a := 1
	b := 2
	b = b - a
	a = b + b
	fmt.Printf("a, b = %v %v\n", a, b)

	c := 3
	d := 4
	c, d = d, c
	fmt.Printf("c, d = %v %v\n", c, d)

	e := 5
	f := 6
	e = e ^ f
	f = e ^ f
	e = e ^ f

	fmt.Printf("e, f = %v %v", e, f)
}
