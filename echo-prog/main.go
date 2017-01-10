package main

import (
	"fmt"
	"os"
)

func main() {

	for indx, arg := range os.Args[0:] {
		fmt.Println(indx, " - ", arg)
	}
}
