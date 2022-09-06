package main

import (
	"fmt"
)

func main() {
	var r float32 = 10.04

	// fmt.Scanf("%f", &r)
	fmt.Println("R =", r)

	fmt.Printf("Area: %0.2f\n", area(r))
}

func area(r float32) (area float32) {
	const p float32 = 3.14159265359
	var circle = p * r * r
	var square = r * 2 * r * 2
	area = square - circle

	return area
}
