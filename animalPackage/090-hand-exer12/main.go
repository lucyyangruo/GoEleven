package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// statement statement
	for i := 0; i < 100; i++ {
		fmt.Println("this is iteration: ", i)
		x := rand.Intn(5)
		if x == 3 {
			fmt.Printf("x is 3.\n")
		}

	}
	// 这个idiom是用在if中，可以将变量的作用域缩短到当前if

	// video
	c := 1
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			// 3 出现的次数
			fmt.Printf("iteration %v \t total count %v \t x is %v\n", i, c, x)
			c++
		}
	}
}
