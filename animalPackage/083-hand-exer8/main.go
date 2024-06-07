package main

import "fmt"

func main() {
	// using only a condition

	x := 0
	for x < 5 {
		fmt.Println("x is: ", x)
		x++
	}

	// infinite loop
	y := 0
	for {
		fmt.Println("y is:", y)
		// 如果没有break就是无限循环，会一直到y is: 522449。。。,但是在go playground就会提示build failed。
		if y == 5 {
			break
		}
		// 如果没有变量的改变条件，那会一直输出上面的内容👆，like 下面
		// y is: 0
		// y is: 0
		y++
	}

	// video
	i := 20
	for i >= 0 {
		fmt.Println(i)
		i--
	}
}
