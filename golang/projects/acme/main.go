package main

import (
	"context"

	"github.com/enesanbar/workspace/golang/projects/acme/internal/config"
	"github.com/enesanbar/workspace/golang/projects/acme/internal/modules/exchange"
	"github.com/enesanbar/workspace/golang/projects/acme/internal/modules/get"
	"github.com/enesanbar/workspace/golang/projects/acme/internal/modules/list"
	"github.com/enesanbar/workspace/golang/projects/acme/internal/modules/register"
	"github.com/enesanbar/workspace/golang/projects/acme/internal/rest"
)

func main() {
	// bind stop channel to context
	ctx := context.Background()

	// build the exchanger
	exchanger := exchange.NewConverter(config.App)

	// build model layer
	getModel := get.NewGetter(config.App)
	listModel := list.NewLister(config.App)
	registerModel := register.NewRegisterer(config.App, exchanger)

	// start REST server
	server := rest.New(config.App, getModel, listModel, registerModel)
	server.Listen(ctx.Done())
}
