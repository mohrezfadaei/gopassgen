package main

import (
	"log"

	"github.com/mohrezfadaei/passgen-cli/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatalf("Error executing CLI: %v", err)
	}
}
