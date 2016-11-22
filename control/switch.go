package main

import (
	"fmt"
)

func main() {
	a := 1

	switch a {
	case 0:
		fmt.Println("a=0")
	case 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")

	}

	switch {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough //even the condition is mathched, the program will continue to be executed
	case a == 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}

//output
/*
-----------------
a=1
a>=0
a=1
-----------------
*/
