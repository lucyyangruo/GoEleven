package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 是运行100次，而不是输出100次
	// x := rand.Intn(10)
	x := 1
	y := rand.Intn(10)
	fmt.Printf("the value of x is %v, the value of y is %v.\n", x, y)
	/**
		// if
		if x < 4 && y < 4 {
			fmt.Println("x and y are both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("x and y are both greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		} else if y != 5 {
			fmt.Println("y is not 5")
		} else {
			fmt.Println("none of the previous cases were met")
		}

	**/
	// for run 100
	for i := 0; i < 100; i++ {
		x := rand.Intn(10)
		fmt.Printf("this is %v times.\n", i)
		switch {
		case x < 4 && y < 4:
			fmt.Println("x and y are both less than 4")
		case x > 6 && y > 6:
			fmt.Println("x and y are both greater than 6")
		case x >= 4 && x <= 6:
			fmt.Println("x is greater than or equal to 4 and less than or equal to 6")
		case y != 5:
			fmt.Println("y is not 5")
		default:
			fmt.Println("none of the previous cases were met")
		}
	}
	// switch

}
