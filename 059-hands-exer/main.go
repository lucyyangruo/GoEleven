package main

import (
	"fmt"

	puppycopy "github.com/lucyyangruo/puppyCopy"
)

func main() {
	fmt.Println("hello, 世界☄️")
	s1 := puppycopy.Bark()
	s2 := puppycopy.Barks()
	s3 := puppycopy.BigBark(s1)
	s4 := puppycopy.BigBarks(s2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	puppycopy.ShowEver()
}
