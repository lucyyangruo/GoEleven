package main

import (
	"fmt"
)

var a string = "42"

const b = 42

func main() {
	// 使用printf可以输出%v %t这些的表示内容，对于字符串可以使用 ""+ a + 的形式来输出，但是对于int这些类型就不能使用+ 的格式。
	// 需要使用 先定义号要输出的内容规范，在通过 , a,b 来输出值。
	fmt.Printf("this is test from video 52. number one_a is : %v, number_b is : %v\n", a, b)
	fmt.Printf("this is test from video 52. number one_a  of type is : %T, number_b of type is : %T\n", a, b)
}
