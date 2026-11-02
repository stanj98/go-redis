package main

import (
	"encoding/base64"
	"fmt"

	"github.com/stanj98/go-redis/store"
	"github.com/stanj98/go-redis/store/kv"
)

func main() {

	cmds := []Command{
		{Op: "SET", Key: "env", Value: "production"},
		{Op: "SET", Key: "version", Value: "0.0.1"},
		{Op: "SET", Key: "debug", Value: "true"},
		{Op: "SET", Key: "env"},
		{Op: "SET", Key: "region", Value: "eu-west-1"},
		{Op: "SET", Key: "version", Value: "version"},
	}

	s := kv.NewStore(0)
	RestoreOnBoot(s, cmds)

	fmt.Println("restored, store size: ", s.Len())
	fmt.Println("keys: ", s.Keys())
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
