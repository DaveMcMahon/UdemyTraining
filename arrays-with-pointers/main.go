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

	for _, v := range myArray {
		fmt.Println(v)
	}

}
