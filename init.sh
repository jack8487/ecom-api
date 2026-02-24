#!/bin/bash
export GOOSE_DBSTRING="host=localhost user=postgres password=postgres dbname=ecom sslmode=disable"
export GOOSE_DRIVER=postgres
export GOOSE_MIGRATIONS_DIR="./internal/adapters/postgresql/migrations/"
goose up