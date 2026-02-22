package main

import (
	"fmt"
	"log/slog"
	"os"
)

func main() {
	cfg := config{
		addr: ":8888",
		db:   dbConfig{},
	}
	api := application{
		config: cfg,
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if err := api.run(api.mount()); err != nil {
		fmt.Printf("err: %s", err)
	}
}
