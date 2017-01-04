package main

import "fmt"

func main() {

	num1 := 10
	num2 := 100
	name := "Dave"
	decNum := 4.56

	var myNum int = 10
	var myName string = "Dave"

	fmt.Printf("%T \n", myNum)
	fmt.Printf("%T \n", myName)

	fmt.Printf("%v \n", num1)
	fmt.Printf("%v \n", num2)
	fmt.Printf("%v \n", name)
	fmt.Printf("%v \n", decNum)
}
