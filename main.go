package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"infohouse/fullhous"
	"os"
)

func main() {
	fullhous.House()
	urlExample := "postgres://postgres:postgres@localhost:5432/test_db"
	_, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("ready")
	}
}
