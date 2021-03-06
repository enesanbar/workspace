package main

import (
	"context"
	"time"

	"github.com/enesanbar/workspace/golang/parallelism/channels"
)

func main() {
	ch := make(chan string)
	done := make(chan bool)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go channels.Printer(ctx, ch)
	go channels.Sender(ch, done)

	time.Sleep(2 * time.Second)
	done <- true
	cancel()
	time.Sleep(3 * time.Second)
}
