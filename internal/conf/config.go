package conf

import (
	"fmt"
	"strings"

	"github.com/BurntSushi/toml"
)

const filename = "config.toml"

type Config struct {
	Author  *string
	Classes *[]Class
}

type Class struct {
	Name    *string
	Teacher *string
	Code    *string
	Folder  string // optional
}

func Read() (Config, error) {
	var config Config
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return Config{}, fmt.Errorf("%v failed to read config from %s", err, filename)
	}
	for i := range *config.Classes {
		if (*config.Classes)[i].Folder == "" {
			(*config.Classes)[i].Folder = strings.ReplaceAll(
				strings.ToLower(*(*config.Classes)[i].Code),
				" ",
				"-",
			)
		}
	}
	return config, nil
}
