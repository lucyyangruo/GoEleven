package main

import (
	"fmt"
)

func main() {
	// if
	x := "James Bond"
	if x == "Janms Bond" {
		fmt.Println(x)
	}
	// if-else if
	y := "Moneypenny"
	if y == "Moneypenny" {
		fmt.Println(y)
	} else if y == "Janms Bond" {
		fmt.Println("BONDDONBONDONBOND", y)
	} else {
		fmt.Println("neither")
	}
	// switch- 默认为true
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
	// switch var
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to the hawaii!")
	case "wingsuit flying":
		fmt.Println("Living dangerously!")

	}
}
