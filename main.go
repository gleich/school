package main

import (
	"pkg.mattglei.ch/school/internal/conf"
	"pkg.mattglei.ch/school/internal/prompt"
	"pkg.mattglei.ch/timber"
)

func main() {
	config, err := conf.Read()
	if err != nil {
		timber.Fatal(err, "failed to read config")
	}

	answers, err := prompt.Ask(*config.Classes)
	if err != nil {
		timber.Fatal(err, "failed to prompt user")
	}

	timber.Debug(answers.Class.Code)
}
