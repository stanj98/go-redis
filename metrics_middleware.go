package main

import "go-redis/store"

type MetricsMiddleware struct {
	inner       store.Storer
	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	val, err := m.inner.Get(key)
	m.getCalls++
	return val, err
}

func (m *MetricsMiddleware) Set(key string, value string) error {
	err := m.inner.Set(key, value)
	m.setCalls++
	return err
}
