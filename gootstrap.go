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
	templateDesc := newTemplateDescriptor(templatesdir)
	return nil
}

type templateDescriptor struct {
	dirs  []string
	files []string
}

func newTemplateDescriptor(dir string) templateDescriptor {
	desc := templateDescriptor{}
	desc.dirs = append(desc.dirs, dir)

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == dir {
				return err
			}
			recursiveDesc := newTemplateDescriptor(path)
			desc.dirs = append(desc.dirs, recursiveDesc.dirs...)
			desc.files = append(desc.files, recursiveDesc.files...)
		} else {
			desc.files = append(desc.files, path)
		}
		return err
	})

	return desc
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
