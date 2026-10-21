package main

import (
	"reflect"
	"testing"
)

func TestKeys_ReturnsAllKeysSorted(t *testing.T) {
	store := NewStore()
	store.Set("charlie", "3")
	store.Set("alpha", "1")
	store.Set("beta", "2")

	got := store.Keys()
	want := []string{"alpha", "beta", "charlie"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Keys() = %v, want: %v", got, want)
	}
}
