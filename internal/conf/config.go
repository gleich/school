package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Author  *string
	Classes *[]Class
}

type Class struct {
	Name    *string
	Teacher *string
	Code    string // optional
}

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("%v failed to get user's home directory", err)
	}

	path := filepath.Join(home, ".config", "school", "config.toml")

	var config Config
	_, err = toml.DecodeFile(path, &config)
	if err != nil {
		return Config{}, fmt.Errorf("%v failed to read config from %s", err, path)
	}
	return config, nil
}
