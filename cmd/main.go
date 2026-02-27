package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jack/ecom/internal/env"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: could not load .env file: %v", err)
	}

	ctx := context.Background()

	cfg := config{
		addr: ":8888",
		db: dbConfig{
			dsn: env.GetString("GOOSE_DBSTRING", "host=172.19.0.2 user=postgres password=postgres dbname=ecom sslmode=disable"),
		},
	}

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
