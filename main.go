package main

import (
	"fmt"
)

func main() {

	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "43")
	s.Set("c", "44")
	val, exists := s.Get("d")
	fmt.Printf("Val : %s, Exists: %t\n", val, exists)

	fmt.Println("GoKV - Go Key Value project")

}
