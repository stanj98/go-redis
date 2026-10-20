package main

import "fmt"

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, bool) {
	value, ok := s.data[key]
	return value, ok
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func main() {

	s := NewStore()
	s.Set("a", "42")
	s.Set("b", "43")
	s.Set("c", "44")
	val, exists := s.Get("d")
	fmt.Printf("Val : %s, Exists: %t\n", val, exists)

	fmt.Println("GoKV - Go Key Value project")
}
