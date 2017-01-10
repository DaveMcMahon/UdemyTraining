package main

import "fmt"

func main() {

	var myArray [5]int
	myArray = [5]int{1, 2, 3, 4, 5}

	for _, val := range myArray {
		fmt.Println(val)
	}

	fmt.Printf("%T\n", myArray)

	fmt.Println("*******************************")

	myNewArray := [...]int{1, 2, 3, 4, 5, 6}
	for _, val := range myNewArray {
		fmt.Println(val)
	}

	fmt.Printf("%T\n", myNewArray)

	fmt.Println(myArray[0])
	fmt.Println(myArray[len(myArray)-1])

	fmt.Println("*******************************")

	var myList [3]int = [3]int{1, 2, 3}
	myList2 := [...]int{1, 2, 3, 4}

	for _, value := range myList {
		fmt.Println(value)
	}

	fmt.Println("*******************************")

	for _, value := range myList2 {
		fmt.Println(value)
	}

}
