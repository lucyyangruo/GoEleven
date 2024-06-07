package main

import "fmt"

func main() {
	// for range loop-clice
	xi := []int{42, 43, 44, 45, 46, 47}
	for key, value := range xi {
		fmt.Printf("xi is key is %v, value is %v.\n", key, value)
	}

	// for range loop-map
	// map key 是string，value是int 用sring来存储int值
	m := map[string]int{
		"James": 42,
		// 最后一个元素也有逗号
		"Moneypenny": 32,
	}

	for key, value := range m {
		fmt.Printf("m has key is %v, has value is %v.\n", key, value)
	}

	// 通过_来省略其中的元素
	for _, value := range m {
		fmt.Printf("m has key is ...., has value is %v.\n", value)
	}

	// add code and print
	age := m["James"]
	fmt.Println(age)

	// lookup of "Q" and print that value
	for key, _ := range m {
		if key == "Q" {
			fmt.Printf("this is Q")
			break
		}
	}

	// "comma ok" idiom习语
	// statement;statement，表示检查m中是否有Q值，有的话ok为T，且value是Q值，没有的话ok为F
	if value, ok := m["Q"]; !ok {
		fmt.Printf("there is no Q %v", value)
	}

	// video
	age1 := m["James"]
	fmt.Println("The age of Bond", age1)
	if v, ok := m["James"]; ok {
		fmt.Println("There is a BOND lookup entry, and Bond's age is", v)
	}

	age2 := m["Q"]
	fmt.Println("That age of Q", age2)

	if v, ok := m["Q"]; !ok {
		fmt.Println("There is no Q, and here is the zero value of an int", v)
	}

}
