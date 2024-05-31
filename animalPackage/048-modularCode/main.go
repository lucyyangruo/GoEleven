package main

import (
	"fmt"

	puppycopy "github.com/lucyyangruo/puppyCopy"
)

func main() {
	s1 := puppycopy.Bark()
	s2 := puppycopy.Barks()
	fmt.Print(s1, ", ", s2)
	// Woof!, Woof! Woof! Woof!
	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("/n")

	s3 := puppycopy.BigBark(s1)
	s4 := puppycopy.BigBarks(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	fmt.Printf("\n")

}
