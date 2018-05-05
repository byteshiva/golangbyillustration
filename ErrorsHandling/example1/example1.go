package main

import (
	"fmt"
)

// http://golang.org/pkg/builtin#error
type error interface {
	Error() string
}

// http://golang.org/pkg/errors/error.go
type errorString struct {
	s string
}

// http://golang.org/pkg/errors/errors.go
func (e *errorString) Error() string {
	return e.s
}

// New returns an error that formats as the given text. http://golang.org/src/pkg/errors/errors.go
func New(text string) error {
	return &errorString{text}
}

func main() {
	if err := randomAPICall(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Life is good... cool")
}

func randomAPICall() error {
	return New("Bad Request")
}
