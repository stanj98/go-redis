package main

import (
	"fmt"
	"go-redis/store"
	"time"
)

type MetricsMiddleware struct {
	inner           store.Storer
	getCalls        int
	setCalls        int
	deleteCalls     int
	lenCalls        int
	getMisses       int
	totalGetLatency time.Duration
	totalSetLatency time.Duration
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	start := time.Now()
	val, err := m.inner.Get(key)
	if err != nil {
		m.getMisses++
	}
	m.getCalls++
	duration := time.Since(start)
	m.totalSetLatency += duration
	return val, err
}

func (m *MetricsMiddleware) Set(key string, value string) error {
	start := time.Now()

	err := m.inner.Set(key, value)
	m.setCalls++

	duration := time.Since(start)
	m.totalSetLatency += duration
	return err
}

func (m *MetricsMiddleware) Delete(key string) {
	m.inner.Delete(key)
	m.deleteCalls++
}

func (m *MetricsMiddleware) Len() int {
	got := m.inner.Len()
	m.lenCalls++
	return got
}

func (m *MetricsMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricsMiddleware) Report() {
	fmt.Printf("GET calls: %d (missed %d)\n", m.getCalls, m.getMisses)
	fmt.Printf("SET calls: %d\n", m.setCalls)
	fmt.Printf("DEL calls: %d\n", m.deleteCalls)
	fmt.Printf("LEN calls: %d\n", m.lenCalls)

	if m.getCalls > 0 {
		avg := m.totalGetLatency / time.Duration(m.getCalls)
		fmt.Printf("Avg GET latency: %d\n", avg)
	}
	if m.setCalls > 0 {
		avg := m.totalSetLatency / time.Duration(m.setCalls)
		fmt.Printf("Avg SET latency: %d\n", avg)
	}
}
