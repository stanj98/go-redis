package main

import (
	"fmt"
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
	case "INCR":
		_, err := s.Incr(c.Key)
		return err
	default:
		return fmt.Errorf("unknown op %q", c.Op)
	}

}

func RestoreOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	errs := make(chan error, len(cmds))

	for _, c := range cmds {
		wg.Go(func() {
			errs <- dispatch(s, c)
		})
		// go dispatch(s, c)
	}

	go func() {
		wg.Wait()
		close(errs)
	}()

	wg.Wait()

	for err := range errs {
		if err != nil {
			fmt.Printf("Error : %v\n", err)
		}
	}
}
