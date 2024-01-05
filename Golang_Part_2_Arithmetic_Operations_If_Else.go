

package main
// Importing main 
import (
	"fmt"
)

func main() {
	var num1 float64 = 9
	var num2 int = 4
	answer := (num1) / float64((num2)) //operations happen on two same data types
	//follow BODMASS Brackets->off->division.....
	// we can also compare strings, it compares letter by letter by its ASCII value
	fmt.Printf("%g", answer)
	// conditional operators <, > , <= , >= , == , !=
	x := 5
	//y:= 6
	val := x < 5
	fmt.Printf("%t \n", val)
	age := 12
	if age >= 18 {
		fmt.Println("You can ride alone")
	} else if age >= 14 {
		fmt.Println("You can ride with a parent")
	} else {
		fmt.Println("You cannot ride")
	}

}
