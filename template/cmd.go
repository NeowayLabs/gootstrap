package template

const Cmd = `package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	version := false
	flag.BoolVar(&version, "version", false, "Show version")
	flag.Parse()

	if version {
		fmt.Printf("version: %s\n", Version)
	}
	return nil
}
`
