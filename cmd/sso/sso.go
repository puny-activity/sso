package main

import (
	"fmt"
	"sso/config"
	"sso/internal/app"
	"sso/pkg/lggr"
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
}
