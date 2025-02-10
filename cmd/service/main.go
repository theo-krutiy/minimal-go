package main

import (
	"fmt"
	"net/http"
	"os"

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
	s := server.New()
	addr := fmt.Sprintf("%s:%s", HOST, PORT)
	err := http.ListenAndServe(addr, s)

	return err
}
