package main

import (
	"errors"
	"fmt"
	"slices"
)

var ErrDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("Key is mandatory")
var ErrStoreFull = errors.New("Store is full")

type Store struct {
	data    map[string]string
	maxSize int
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}
	value, ok := s.data[key]
	if !ok {
		return "", ErrDoesNotExist
	}
	return value, nil
}

func (s *Store) Set(key, value string) error {
	if key == "" {
		return ErrEmptyKey
	}

	_, exists := s.data[key]

	//check capacity
	if s.maxSize > 0 && s.Len() >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	}

	s.data[key] = value
	return nil
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

func NewStore(maxSize int) *Store {
	return &Store{
		data:    make(map[string]string),
		maxSize: maxSize,
	}
}

func (s *Store) Len() int {
	return len(s.data)
}
