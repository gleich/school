package main

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
	"pkg.mattglei.ch/school/internal/conf"
	"pkg.mattglei.ch/school/internal/prompt"
	"pkg.mattglei.ch/school/internal/template"
	"pkg.mattglei.ch/timber"
)

func main() {
	timber.SetTimezone(time.Local)
	timber.SetTimeFormat("15:04:05")

	config, err := conf.Read()
	if err != nil {
		timber.Fatal(err, "failed to read config")
	}

	err = conf.CreateClassFolders(*config.Classes)
	if err != nil {
		timber.Fatal(err, "failed to create class folders")
	}

	answers, err := prompt.Ask(*config.Classes)
	if err != nil {
		timber.Fatal(err, "failed to prompt user")
	}

	templ, err := template.Read(answers.TemplatePath)
	if err != nil {
		timber.Fatal(err, "failed to read template")
	}

	now := time.Now()
	err = template.Execute(template.Template{
		Name:   answers.DocumentName,
		Author: *config.Author,
		Class:  answers.Class,
		Date: fmt.Sprintf(
			"%s %s, %s",
			now.Format("Monday, January"),
			humanize.Ordinal(now.Day()),
			now.Format("2006"),
		),
	}, templ)
	if err != nil {
		timber.Fatal(err, "failed to execute template")
	}
}
