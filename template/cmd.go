package template

const Cmd = `package main

import (
	"flag"
	"fmt"
	"{{.Module}}/pkg/version"
)

func main() {
	showVersion := flag.Bool("version", false, "Show version")
	flag.Parse()

	if *showVersion {
		fmt.Print(version.Version())
	}
}
`
