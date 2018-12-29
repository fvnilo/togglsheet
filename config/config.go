package config

import (
	"github.com/BurntSushi/toml"
	toggl_export "github.com/nylo-andry/toggl-export"
)

// ReadConfig reads the given config file and returns the necessary
// configuration to the application.
func ReadConfig(file string) (*toggl_export.Config, error) {
	var config toggl_export.Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
