package main

import (
	"fmt"
	"net/http"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

type SpyStore struct {
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResposeWriter, r *http.Request) {
		store.Cancel()
		fmt.Fprint(w, store.Fetch())
	}
}
