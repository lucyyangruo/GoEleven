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

	// 解答问题，测试rand的包含性和排他性
	// include zero in its output? yes
	// include 250 in its output? no is [0,250)
	// show this in code using the numbers 0 - 3 rand.Intn(3)
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
	println(rand.Intn(3))
}
