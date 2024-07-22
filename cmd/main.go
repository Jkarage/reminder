package main

import (
	"log/slog"
	"os"
)

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
		os.Exit(-1)
	}
}

func run() error {
	//todo: Implement the startup logic

	return nil
}
