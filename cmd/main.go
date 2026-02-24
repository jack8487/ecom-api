package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/jack/ecom/internal/env"
	"github.com/jackc/pgx/v5"
)

func main() {

	ctx := context.Background()

	cfg := config{
		addr: ":8888",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// DataBase
	conn, err := pgx.Connect(ctx, cfg.db.dsn)
	if err != nil {
		panic(err)
	}

	defer conn.Close(ctx)

	logger.Info("Connect to database", "dsn", cfg.db.dsn)

	api := application{
		config: cfg,
		db:     conn,
	}

	if err := api.run(api.mount()); err != nil {
		fmt.Printf("err: %s", err)
		os.Exit(1)
	}
}
