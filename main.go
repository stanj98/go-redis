package main

import (
	"errors"
	"fmt"
)

func main() {

	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "43")
	s.Set("c", "44")
	s.Delete("a")

	val, err := s.Get("a")
	if err != nil {
		if errors.Is(err, ErrDoesNotExist) {
			fmt.Println(err)
		}
	}
	fmt.Println(val)

}
