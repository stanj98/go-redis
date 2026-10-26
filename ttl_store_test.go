package main

import (
	"testing"
	"time"
)

func TestTTLStore_GetBeforeExpiry(t *testing.T) {
	s := NewTTLStore(1 * time.Second)
	_ = s.Set("k", "v")

	got, err := s.Get("k")
	if err != nil || got != "v" {
		t.Fatalf("expected v, got %q %v", got, err)
	}
}

func TestTTLStore_GetAfterExpiry(t *testing.T) {
	s := NewTTLStore(50 * time.Millisecond)
	_ = s.Set("k", "v")

	time.Sleep(100 * time.Millisecond)

	_, err := s.Get("k")
	if err == nil {
		t.Fatalf("expected error after TTL, got nil")
	}
}

func TestTTLStore_ExpiredKeyEvicted(t *testing.T) {
	s := NewTTLStore(50 * time.Millisecond)
	_ = s.Set("k", "v")
}

//1:26:04
