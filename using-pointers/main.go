package main

import "fmt"

func changeValue(x *int) {

	fmt.Println(&x)

	*x = 0
}

func main() {

	myNum := 10

	fmt.Println(&myNum)

	changeValue(&myNum)

	fmt.Println(&myNum)
}
