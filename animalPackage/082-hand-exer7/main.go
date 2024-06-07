package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// switch 另外一种写法 [0,5)
	// run the code 42 times
	for i := 0; i < 42; i++ {
		fmt.Println("run the code for:", i+1, "9999ok")
		x := rand.Intn(5)
		fmt.Printf("the value of x is %v.\n", x)

		switch x {
		case 0:
			fmt.Println("x value is 0.")
		case 1:
			fmt.Println("x value is 1.")
		case 2:
			fmt.Println("x value is 2.")
		case 3:
			fmt.Println("x value is 3.")
		default:
			fmt.Println("x value is 4.")
		}
	}
}
