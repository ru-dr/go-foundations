package main

import (
	"errors"
	"fmt"
)

type Store struct {
	data     map[string]string
	readOnly bool
}

type ValidationError struct {
	Field string
	Value any
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("invalid %s : %v", e.Field, e.Value)
}

var (
	ErrNotFound = errors.New("key not found")
	ErrReadOnly = errors.New("store is read-only")
)

func (s *Store) Get(key string) (string, error) {
	val, ok := s.data[key]
	if !ok {
		return "", fmt.Errorf("get %q: %w", key, ErrNotFound)
	}
	return val, nil
}

func (s *Store) Set(key string, value string) error {
	if key == "" {
		return &ValidationError{Field: "key", Value: key}
	}
	if s.readOnly {
		return fmt.Errorf("set %q: %w", key, ErrReadOnly)
	}
	s.data[key] = value
	return nil
}

func (s *Store) GetRequired(key string) (string, error) {
	val, err := s.Get(key)
	if err != nil {
		return "", fmt.Errorf("required config - %q: %w", key, err)
	}
	return val, nil
}

func handle(err error) {
	if errors.Is(err, ErrNotFound) {
		fmt.Println("missing key")
		return
	}
	if errors.Is(err, ErrReadOnly) {
		fmt.Println("read only")
		return
	}
	var vErr *ValidationError
	if errors.As(err, &vErr) {
		fmt.Println("bad field:", vErr.Field)
		return
	}
	fmt.Println("unknown error:", err)
}

func main() {
	s := &Store{data: map[string]string{"host": "localhost"}}

	_, err := s.GetRequired("port")
	handle(err)

	handle(s.Set("", "x"))

	ro := &Store{data: map[string]string{}, readOnly: true}
	handle(ro.Set("k", "v"))
}
