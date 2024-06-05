package main

import (
	"math/rand"
)

func init() {
	x := rand.Intn(250)
	// fmt.Printf("the value of x is %v \n", x)
	// println("this is a random X-int: ", x)

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
}
func main() {
	println("hello 世界☄️")
}
