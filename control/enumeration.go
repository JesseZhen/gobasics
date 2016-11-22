package main

import (
	"fmt"
)

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
