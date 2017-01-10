package main

import "fmt"

func baseAll(anArray *[6]int) {
	for i := range anArray {
		anArray[i] = 0
	}
}

func main() {

	myArray := [...]int{12, 34, 67, 34, 66, 100}
	baseAll(&myArray)

	var myOtherArray [6]int = [...]int{1, 2, 3, 4, 5, 6}
	baseAll(&myOtherArray)

	for _, v := range myArray {
		fmt.Println(v)
	}

	fmt.Println("********")

	for _, v := range myOtherArray {
		fmt.Println(v)
	}
}
