package conf

import (
	"fmt"

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
}

func Read() (Config, error) {
	var config Config
	_, err := toml.DecodeFile(filename, &config)
	if err != nil {
		return Config{}, fmt.Errorf("%v failed to read config from %s", err, filename)
	}
	return config, nil
}
