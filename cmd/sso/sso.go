package main

import (
	"fmt"
	"os"
	"os/signal"
	"sso/config"
	"sso/internal/app"
	"sso/pkg/lggr"
	"syscall"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		panic(fmt.Errorf("failed to parse config: %w", err))
	}

	log := lggr.New(cfg.Logger())

	application := app.New(cfg.App(), log)
	err = application.Start()
	if err != nil {
		panic(fmt.Errorf("failed to start application: %w", err))
	}
	defer application.Stop()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	<-interrupt
}
