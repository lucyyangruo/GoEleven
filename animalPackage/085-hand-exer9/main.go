package main

import "fmt"

func main() {
	// 使用 余数 % + continue
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("the odd is :", i)
	}
}
