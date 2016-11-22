package main

import (
	"fmt"
)

func main() {

	var a *int
	fmt.Println(a)

	b := 100
	var p *int = &b
	fmt.Println(p)
	fmt.Println(*p)
}

//output
/*

<nil>
0xc0420441d8
100

*/

