package config

import (
	"github.com/BurntSushi/toml"
	toggl_export "github.com/nylo-andry/toggl-export"
)

func ReadConfig(file string) (*toggl_export.Config, error) {
	var config toggl_export.Config
	if _, err := toml.DecodeFile(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
