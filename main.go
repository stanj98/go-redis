package main

import (
	"errors"
	"fmt"
)

func main() {

	s := NewStore(2)
	_ = s.Set("a", "42")
	_ = s.Set("b", "72")

	if err := s.Set("c", "2"); err != nil {
		fmt.Println(err)
		return
	}

	val, err := s.Get("a")
	if err != nil {
		if errors.Is(err, ErrDoesNotExist) {
			fmt.Println(err)
		}
	}
	fmt.Println(val)

}
