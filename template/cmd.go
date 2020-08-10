package template

const Cmd = `package main

import (
	"flag"
	"fmt"
)

var version = "dev" // this will be set on build time

func main() {
	showVersion := false
	flag.BoolVar(&showVersion, "version", false, "Show version")
	flag.Parse()

	if showVersion {
		fmt.Printf("version: %s\n", version)
	}
}
`
