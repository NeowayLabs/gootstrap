package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func getTemplatesToOutput(
	templatesdir string,
	outputdir string,
	project string,
) map[string]string {
	files := getTemplateFiles(templatesdir)
	templateToOutput := map[string]string{}
	outputdirs := map[string]struct{}{}

	for _, templatefile := range files {
		templatedPath, ok := relativePath(templatefile)
		if !ok {
			fmt.Printf("unexpected path[%s], unable to parse it\n", templatefile)
			os.Exit(1)
		}
		relativePath := applyPathTemplate(templatedPath, project)
		outputfile := path.Join(outputdir, relativePath)
		templateToOutput[templatefile] = outputfile
		outputdirs[filepath.Dir(outputfile)] = struct{}{}
	}

	for outdir := range outputdirs {
		err := os.MkdirAll(outdir, 0775)
		fatalerr(err, fmt.Sprintf("creating inner output dir[%s]", outdir))
	}

	return templateToOutput
}

func applyPathTemplate(templatedPath string, project string) string {
	return strings.Replace(templatedPath, "{{project}}", project, -1)
}

func relativePath(path string) (string, bool) {
	parsedPath := strings.Split(path, "/templates/")
	if len(parsedPath) > 1 {
		return parsedPath[len(parsedPath)-1], true
	}

	return "", false
}

func getTemplateFiles(dir string) []string {
	files := []string{}

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == dir {
				return err
			}
			recursiveFiles := getTemplateFiles(path)
			files = append(files, recursiveFiles...)
		} else {
			files = append(files, path)
		}
		return err
	})

	return files
}

func applyTemplates(
	templateToOutput map[string]string,
	project string,
	dockerRegistry string,
) {
	type ProjectInfo struct {
		Name           string
		DockerRegistry string
	}
	projectInfo := ProjectInfo{
		Name:           project,
		DockerRegistry: dockerRegistry,
	}
	for templatefile, outputfile := range templateToOutput {
		out, err := os.Create(outputfile)
		fatalerr(err, fmt.Sprintf("creating output file[%s]", outputfile))
		defer out.Close()

		tmpl, err := template.ParseFiles(templatefile)
		fatalerr(err, fmt.Sprintf("creating template from file[%s]", templatefile))

		err = tmpl.Execute(out, projectInfo)
		fatalerr(err, fmt.Sprintf("executing template[%s]", templatefile))
	}
}

func fatalerr(err error, context string) {
	if err != nil {
		fmt.Printf("error[%s] while[%s]\n", err, context)
		os.Exit(1)
	}
}

func getGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath != "" {
		return gopath
	}
	return path.Join(os.Getenv("HOME"), "go")
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

func requiredOption(v string, msg string) {
	if v == "" {
		fmt.Println(msg + ":\n")
		flag.PrintDefaults()
		fmt.Println()
		os.Exit(1)
	}
}

func main() {
	outputdir := ""
	project := ""
	dockerRegistry := ""

	cwd, err := os.Getwd()
	fatalerr(err, "getting default cwd")
	templatesdir := getTemplatesDir()

	flag.StringVar(&project, "project", "", "project name (required)")
	flag.StringVar(&dockerRegistry, "docker-registry", "", "docker registry base name (required)")
	flag.StringVar(&outputdir, "output", cwd, "output directory where project will be created (optional)")
	flag.Parse()

	requiredOption(project, "'project' is a required option")
	requiredOption(dockerRegistry, "'docker-registry' is a required option")

	fmt.Printf("\ncreating: project[%s] docker-registry[%s] output[%s]\n\n", project, dockerRegistry, outputdir)

	templateToOut := getTemplatesToOutput(templatesdir, outputdir, project)

	applyTemplates(templateToOut, project, dockerRegistry)

	fmt.Println("success")
}
