package main

import "fmt"

func main() {
	// 嵌套循环
	for i := 0; i < 5; i++ {
		if i == 0 {
			for j := 0; j < 5; j++ {
				fmt.Println("the first loop is: ", i, " time. the second loop is:", j)
			}
		}
		fmt.Println("the first loop is: ", i)
	}
}
