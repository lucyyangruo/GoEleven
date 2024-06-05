package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 打印随机数
	x := rand.Intn(250)
	fmt.Printf("the value of x is %v \n", x)
	println("this is a random X-int: ", x)

	// if statement
	if 0 <= x && x < 100 {
		println("print between 0 and 100")
	} else if 101 <= x && x < 200 {
		println("print between 101 and 200")
	} else {
		println("print between 201 and 250")
	}

	// switch
	switch {
	case x <= 100:
		println("print between 0 and 100")
	case x >= 101 && x <= 200:
		println("print between 101 and 200")
	case x >= 201 && x < 250:
		println("print between 201 and 250")
	default:
		println("never show")
	}

	// if ：not us
	// 在 Go 语言中，select 语句用于处理多个通道操作，而不是简单的条件判断。
	// y := rand.Intn(350)
	// select {
	// // select case must be send or receive (possibly with assignment)
	// // 在case内复制
	// case y <= 100:
	// 	println("print between 0 and 100")
	// case y >= 101 && y <= 200:
	// 	println("print between 0 and 100")
	// case y >= 201 && y < 250:
	// 	println("print between 0 and 100")
	// }
}
