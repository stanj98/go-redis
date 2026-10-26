package main

import (
	"encoding/base64"
	"fmt"
)

func main() {

	s := NewStore(2)

	encrypted, err := SetKeyWithEncryption(*s, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encrypted)
	ttlStore := NewTTLStore()

	encryptedTTLValue, err := SetKeyWithEncryption(*ttlStore, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encryptedTTLValue)
}

func SetKeyWithEncryption(store Store, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return s.Get(key)
}
