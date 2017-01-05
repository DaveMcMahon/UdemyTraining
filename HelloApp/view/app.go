package view

import "fmt"

var x = 10
var myName =  "Jason"

func DisplayVar(){
	fmt.Printf("The value from the view package is: %v\n", x)
}

func DisplayName(){
	println(myName)
}