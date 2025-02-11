package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/theo-krutiy/minimal-go/internal/db"
	"github.com/theo-krutiy/minimal-go/internal/server"
)

const HOST = "localhost"
const PORT = "8080"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	pg, err := db.NewPostgres(context.Background(), "postgres://postgres:password@localhost:5432/postgres")
	if err != nil {
		return err
	}

	s := server.New()
	s.Db = pg

	addr := fmt.Sprintf("%s:%s", HOST, PORT)
	err = http.ListenAndServe(addr, s)
	return err
}
