package main

import (
	"fmt"
)

//implemented the numeration of computer storage units by iota
const (
	B int64 = 1 << (iota * 10)
	KB
	MB
	GB
	TB
	PB
)

func main() {                          
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
}

//output:
/* 

1
1024
1048576
1073741824
1099511627776
1125899906842624 

*/
