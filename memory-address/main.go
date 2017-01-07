package main

import "fmt"

func main() {

	var userName string

	fmt.Print("Enter in your name: ")
	fmt.Scan(&userName)

	if userName == "Dave" {
		fmt.Printf("Your name is %s\n", userName)
	} else {
		fmt.Println("Sorry your not Dave, access denied")
	}

}
