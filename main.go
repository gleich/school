package main

import (
	"pkg.mattglei.ch/school/internal/conf"
	"pkg.mattglei.ch/timber"
)

func main() {
	config, err := conf.Read()
	if err != nil {
		timber.Fatal(err, "failed to read config")
	}

	timber.Debug(*config.Author)
}
