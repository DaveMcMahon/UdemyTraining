package main

import "fmt"

func main() {

	a := 50

	fmt.Println("a has a value of ", a)
	fmt.Println("the value of a is stored in memory location: ", &a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

	*b = 60
	fmt.Println(a)

}
