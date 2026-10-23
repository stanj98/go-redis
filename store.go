package main

import (
	"errors"
	"slices"
)

var ErrDoesNotExist = errors.New("Key does not exist")

type Store struct {
	data map[string]string
}

func (s *Store) Get(key string) (string, error) {
	value, ok := s.data[key]
	if !ok {
		return "", ErrDoesNotExist
	}
	return value, nil
}

func (s *Store) Set(key, value string) {
	s.data[key] = value
}

func (s *Store) Delete(key string) {
	delete(s.data, key)
}

// keys return every key currently in the store, sorted alphabetically
func (s *Store) Keys() []string {
	keys := make([]string, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	//sort the keys alphabetically
	slices.Sort(keys)
	return keys
}

func NewStore() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Len() int {
	return len(s.data)
}
