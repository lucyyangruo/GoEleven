package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)

	// video
	fmt.Printf("true && true \t\t%v\n", (true && true))
	fmt.Printf("true && false \t\t%v\n", (true && false))
	fmt.Printf("true || true \t\t%v\n", (true || true))
	fmt.Printf("true || false \t\t%v\n", (true || false))
	fmt.Printf("!true \t\t\t%v\n", !true)
}
