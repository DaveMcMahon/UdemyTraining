package main

import "fmt"

var x = 4

func main() {
	fmt.Printf("The value of x is: %v\n", x)
	foo()
}

func foo() {
	fmt.Printf("The value of x in the foo function is: %v\n", x)
}