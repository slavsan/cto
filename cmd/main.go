package main

import (
	"os"

	"github.com/slavsan/cto"
)

func main() {
	err := cto.Colorize(os.Stdin, os.Stdout)
	if err != nil {
		os.Exit(1)
	}
}
