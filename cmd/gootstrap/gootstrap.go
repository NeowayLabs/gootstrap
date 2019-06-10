package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/NeowayLabs/gootstrap"
)

const GoDigest = "sha256:d17a1d8f0c20d108d1177d560f4afb9de10104c46df756d885cfa4282bbaac65" // golang:1.12.5-stretch
const CILintVersion = "1.13.2"
const AlpineDigest = "sha256:769fddc7cc2f0a1c35abb2f91432e8beecf83916c421420e6a6da9f8975464b6" // alpine:3.9

func main() {

	outputdir := ""
	module := ""
	dockerimg := ""

	flag.StringVar(
		&dockerimg,
		"image",
		"",
		"docker image of the project",
	)
	flag.StringVar(
		&module,
		"module",
		"",
		"The module name of the project, like: 'github.com/NeowayLabs/gootstrap'",
	)
	flag.StringVar(
		&outputdir,
		"output-dir",
		getcwd(),
		"directory where the generated files are going to be saved",
	)

	flag.Parse()

	if module == "" {
		fmt.Println("-module is an obligatory parameter")
		os.Exit(1)
	}

	if dockerimg == "" {
		fmt.Println("-docker-image is an obligatory parameter")
		os.Exit(1)
	}

	fmt.Printf("creating project module[%s] docker-image[%s] files at dir[%s]\n",
		module, dockerimg, outputdir)

	project, err := parseNameFromModule(module)
	abortonerr(err)

	cfg := gootstrap.Config{
		Project:       project,
		Module:        module,
		DockerImg:     dockerimg,
		GoDigest:      GoDigest,
		CILintVersion: CILintVersion,
		AlpineDigest:  AlpineDigest,
	}
	gootstrap.CreateProject(cfg, outputdir)
}

func getcwd() string {
	wd, err := os.Getwd()
	abortonerr(err)
	return wd
}

func abortonerr(err error) {
	if err != nil {
		panic(err)
	}
}

func parseNameFromModule(module string) (string, error) {
	// Go modules are like this: github.com/NeowayLabs/gootstrap
	// Lets assume that the last component of the path is the project name
	parsed := strings.Split(module, "/")
	if len(parsed) == 1 {
		return "", fmt.Errorf("invalid module[%s] cant extract project name from it", module)
	}

	return parsed[len(parsed)-1], nil
}
