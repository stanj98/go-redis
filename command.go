package main

import (
	"log"
	"sync"

	"go-redis/store"
)

type Command struct {
	Op    string
	Key   string
	Value string
}

func dispatch(s store.Storer, c Command) {
	switch c.Op {
	case "SET":
		log.Printf("SET key %s", c.Key)
		s.Set(c.Key, c.Value)
	case "GET":
		s.Get(c.Key)
	}

}

func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	for _, c := range cmds {
		wg.Go(func() {
			dispatch(s, c)
		})
		// go dispatch(s, c)
	}

	wg.Wait()
}
