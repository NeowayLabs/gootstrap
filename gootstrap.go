package gootstrap

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	texttemplate "text/template"

	"github.com/NeowayLabs/gootstrap/template"
)

// Config used by template functions
type Config struct {
	Project       string
	Module        string
	DockerImg     string
	GoVersion     string
	CILintVersion string
	AlpineVersion string
}

// CreateProject creates a project
func CreateProject(cfg Config, rootdir string) error {
	files := map[string]string{
		"go.sum":                           "",
		"go.mod":                           template.GoMod,
		"README.md":                        template.Readme,
		"Makefile":                         template.Makefile,
		"Dockerfile":                       template.Dockerfile,
		".gitignore":                       template.GitIgnore,
		".dockerignore":                    template.DockerIgnore,
		".gitlab-ci.yml":                   template.GitlabCI,
		"cmd/{{.Project}}/{{.Project}}.go": template.Cmd,
	}

	for pathtmpl, contenttmpl := range files {
		path, err := execTemplate(pathtmpl, pathtmpl, cfg)
		if err != nil {
			return err
		}

		contents, err := execTemplate(path, contenttmpl, cfg)
		if err != nil {
			return err
		}

		if err := writeFile(filepath.Join(rootdir, path), contents); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(path string, contents string) error {
	dir := filepath.Dir(path)
	if err := ensureDir(dir); err != nil {
		return err
	}

	if _, err := os.Stat(path); err == nil {
		fmt.Printf("file[%s] already exists, it will not be touched, nothing to do\n", path)
		return nil
	}
	return ioutil.WriteFile(path, []byte(contents), 0644)
}

func execTemplate(name string, templ string, cfg Config) (string, error) {
	tmpl, err := texttemplate.New(name).Parse(templ)
	if err != nil {
		return "", fmt.Errorf("error creating %s template: %s", name, err)
	}

	buf := &bytes.Buffer{}

	err = tmpl.Execute(buf, cfg)
	if err != nil {
		return "", fmt.Errorf("error executing %s template: %s", name, err)
	}

	return buf.String(), nil
}

func ensureDir(dir string) error {
	return os.MkdirAll(dir, 0755)
}
