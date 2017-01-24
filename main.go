package main

import (
	"errors"
	"strings"
)

func main() {
}

// Services are modeled as interfaces with corresponding implementations.
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

var ErrEmpty = errors.New("Empty string")

// Every method in the iterface is modeled as a RPC. We define request & response structs for each method,
//	capturing all input/output parameters respectively.
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}