package template

import (
	"bytes"
	"fmt"
	"text/template"
)

func apply(name string, templ string, cfg interface{}) (string, error) {
	tmpl, err := template.New(name).Parse(templ)
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
