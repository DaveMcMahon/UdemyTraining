package main

import (
	"fmt"

	"github.com/DaveMcMahon/UdemyTraining/stringutils"
)

var x = 4

func main() {
	fmt.Printf("The value of x is: %v\n", x)
	foo()

	stringutils.PrintIt("123")

}

func foo() {
	fmt.Printf("The value of x in the foo function is: %v\n", x)
}
