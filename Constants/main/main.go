package main

import "fmt"

func main() {

	var myName = "Dave"
	myAge := 25
	var myDob int = 1991
	const myRealName string = "DaveMcMahon"

	if myAge == 25{
		fmt.Println("You are 25!!")
	}

	fmt.Println(myName)
	fmt.Println(myAge)
	fmt.Println(myDob)
	fmt.Println(myRealName)

}
