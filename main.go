package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/stanj98/go-redis/store"
	"github.com/stanj98/go-redis/store/kv"
	"github.com/stanj98/go-redis/store/ttl"
)

func main() {

	s := kv.NewStore(2)

	encrypted, err := SetKeyWithEncryption(s, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encrypted)
	ttlStore := ttl.NewTTLStore(time.Second * 2)

	encryptedTTLValue, err := SetKeyWithEncryption(ttlStore, "a", "hey you can see me")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(encryptedTTLValue)
}

func SetKeyWithEncryption(store store.Storer, key, val string) (string, error) {
	encoded := base64.StdEncoding.EncodeToString([]byte(val))
	if err := store.Set(key, encoded); err != nil {
		return "", err
	}
	return store.Get(key)

}
