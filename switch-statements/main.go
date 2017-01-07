package main

import "fmt"

func main() {

	myName := "Dave"

	switch myName {
	case "Dave":
		fmt.Println("Welcome, Dave")
	case "Michelle":
		fmt.Println("Welcome, Lorna")
	case "Derek":
		fmt.Println("Welcome, Derek")
	default:
		fmt.Println("Welcome Casper!!")
	}

}
