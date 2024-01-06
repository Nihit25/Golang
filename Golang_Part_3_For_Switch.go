package main

import (
	"fmt"
)

func main() {
	x := 3

  // for loop in Golang 
	for x < 5 {
		fmt.Println(x)
		x++
	}
	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}
	// break and continue works as it is
	ans := -1
	//switch statements
	switch ans {
	case 1, -1: // means when ans = 1, but here -1
		fmt.Println("case one")
	case 2:
		fmt.Println("case two")
	default:
		fmt.Println("default case")
	}
}
