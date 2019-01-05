package config

import (
	"github.com/BurntSushi/toml"

	"github.com/nylo-andry/togglsheet"
)

// ReadConfig reads the given config file and returns the necessary
// configuration to the application.
func ReadConfig(file string) (*togglsheet.Config, error) {
	var config togglsheet.Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
