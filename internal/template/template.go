package template

import (
	"fmt"
	"os"
	"text/template"

	"pkg.mattglei.ch/school/internal/conf"
)

type Template struct {
	Name   string
	Author string
	Class  conf.Class
	Date   string
}

func Read(path string) (*template.Template, error) {
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return &template.Template{}, fmt.Errorf("%v failed to read from file", err)
	}

	return tmpl, nil
}

func Execute(data Template, template *template.Template) error {
	err := template.Execute(os.Stdout, data)
	if err != nil {
		return fmt.Errorf("%v failed to execute template", err)
	}
	return nil
}
