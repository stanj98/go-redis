package kv

import (
	"encoding/base64"
	"errors"
	"fmt"
	"go-redis/store"
	"slices"
	"strconv"
	"sync"
	"time"
)

var ErrDoesNotExist = errors.New("Key does not exist")
var ErrEmptyKey = errors.New("Key is mandatory")
var ErrStoreFull = errors.New("Store is full")

type Store struct {
	mu      sync.Mutex
	data    map[string]string
	maxSize int
}

func (s *Store) Get(key string) (string, error) {
	if key == "" {
		return "", ErrEmptyKey
	}
	s.mu.Lock()
	defer s.mu.Unlock()
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

	s.mu.Lock()
	defer s.mu.Unlock()

	time.Sleep(time.Second)

	_, exists := s.data[key]

	//check capacity
	if s.maxSize > 0 && len(s.data) >= s.maxSize && !exists {
		return fmt.Errorf("Set(%q): %w", key, ErrStoreFull)
	}

	s.data[key] = value
	return nil
}

func (s *Store) Delete(key string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, key)
}

// keys return every key currently in the store, sorted alphabetically
func (s *Store) Keys() []string {
	s.mu.Lock()
	defer s.mu.Unlock()
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

func (s *Store) Clone() *Store {
	cp := &Store{
		data:    make(map[string]string, len(s.data)),
		maxSize: s.maxSize,
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	for k, v := range s.data {
		cp.data[k] = v
	}
	return cp
}

func (s *Store) Len() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return len(s.data)
}

func (s *Store) SetKeyWithEncryption(key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := s.Set(key, encoded); err != nil {
		return "", err
	}
	return s.Get(key)
}

func (s *Store) Incr(key string) (int, error) {
	if key == "" {
		return 0, store.ErrEmptyKey
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	current := 0
	if raw, ok := s.data[key]; ok {
		v, err := strconv.Atoi(raw)
		if err != nil {
			return 0, fmt.Errorf("INCR %q: %w", key, err)
		}
		current = v
	}
	current++
	time.Sleep(time.Microsecond)

	s.data[key] = strconv.Itoa(current)

	return current, nil
}
