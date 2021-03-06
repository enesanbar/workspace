package main

import (
	"context"
	"fmt"

	"github.com/enesanbar/workspace/golang/apis/rpc/twirp/rpc/greeter"
)

// Greeter implements the interface
// generated by protoc
type Greeter struct {
	Exclaim bool
}

// Greet implements twirp Greet
func (g *Greeter) Greet(ctx context.Context, r *greeter.GreetRequest) (*greeter.GreetResponse, error) {
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())
	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}
	return &greeter.GreetResponse{Response: msg}, nil
}
