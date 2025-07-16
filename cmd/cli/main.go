package main

import (
	"GoAI/internal/cli"
	"os"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
