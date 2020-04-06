package main

import "fmt"

func main() {
	var x float64
	fmt.Printf("Enter a floating point number: ")
	fmt.Scan(&x)
	fmt.Printf("Truncated value: %d", int(x))
}
