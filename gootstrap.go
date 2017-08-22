package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func getTemplatesToOutput(
	templatesdir string,
	outputdir string,
) map[string]string {
	templateToOut := map[string]string{}
	files := getfiles(templatesdir)
	return nil
}

func getfiles(dir string) []string {
	files := []string{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) {
		if info.IsDir() {
			files = append(files, getfiles(path))
		} else {
			files = append(files, path)
		}
	})

	return files
}

func prepareDirectories(templateToOut map[string]string) {
}

func applyTemplates(
	templateToOut map[string]string,
	project string,
	docker_registry string,
) {
}

func fatalerr(err error, context string) {
	if err != nil {
		panic(fmt.Sprintf("error[%s] while[%s]", err, context))
	}
}

func getGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		return gopath
	}
	return os.Getenv("HOME") + "/go"
}

func getTemplatesDir() string {
	templatesdir := getGoPath() + "/src/github.com/NeowayLabs/gootstrap/templates"
	if _, err := os.Stat(templatesdir); os.IsNotExist(err) {
		errmsg := fmt.Sprintf(
			"templates directory[%s] do not exist, probably "+
				"a bad/corrupted gootstrap installation\n"+
				"are you sure you installed with 'go get' ?",
			templatesdir,
		)
		panic(fmt.Sprintf("%s\nerror[%s]", errmsg, err))
	}
	return templatesdir
}

func main() {
	outputdir := ""
	project := ""
	docker_registry := ""

	cwd, err := os.Getwd()
	fatalerr(err, "getting default cwd")
	templatesdir := getTemplatesDir()

	flag.StringVar(&project, "project", "", "project name")
	flag.StringVar(&docker_registry, "docker-registry", "", "docker registry base name")
	flag.StringVar(&outputdir, "output", cwd, "output directory where all files will be created")
	flag.Parse()

	templateToOut := getTemplatesToOutput(templatesdir, outputdir)
	prepareDirectories(templateToOut)
	applyTemplates(templateToOut, project, docker_registry)
}
