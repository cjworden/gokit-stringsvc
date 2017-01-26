package main

import (
	"errors"
	"strings"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
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

// Every method in the interface is modeled as a RPC. We define request & response structs for each method,
//	capturing all input/output parameters respectively.
type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}

// Go kit provides it's functionality through endpoints: type Endpoint func(ctx context.Context, request interface{}) (response interface{}, err error)
// An endpoint represents a single RPC, or single method of our service interface. We write adapters to convert our
// service methods into endpoints.
func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func (ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, " "}, nil
	}
}
