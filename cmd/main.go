package main

import (
	_ "api/app/ports-in"
	"context"

	"api/pkg/log"
	"api/pkg/ports/adapters"
)

func main() {
	logger := log.New(context.Background())
	logger.Info("Starting app...")

	adapters.FiberListen()
}
