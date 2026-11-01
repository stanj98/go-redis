package main

import (
	"encoding/base64"
	"fmt"

	"github.com/stanj98/go-redis/store"
	"github.com/stanj98/go-redis/store/kv"
)

func main() {

	kvStore := kv.NewStore(20)
	PopulateDefaults(kvStore)

	newStore := kvStore.Clone()
	fmt.Println(kvStore)
	fmt.Println(newStore)
}

func PopulateDefaults(s store.Storer) error {
	defaults := map[string]string{
		"env":     "development",
		"version": "0.0.1",
		"debug":   "true",
	}
	for k, v := range defaults {
		if err := s.Set(k, v); err != nil {
			return fmt.Errorf("PopulateDefaults: %w", err)
		}
	}
	return nil
}

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)

}
