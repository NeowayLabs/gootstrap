package main

import (
	"flag"
	"fmt"
)

var (
	// version is set at build time
	Version = "No version provided at build time"
)

func main() {

	version := false
	flag.BoolVar(&version, "version", false, "Show version")
	flag.Parse()

	if version {
		fmt.Printf("build tag: %s\n", Version)
		return
	}
}
