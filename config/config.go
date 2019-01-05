package config

import (
	"bufio"
	"log"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/nylo-andry/togglsheet"
)

const defaultConfigFileName = ".togglsheet"
const apiTokenKey = "api_token"
const workspaceIDKey = "workspace_id"
const usernameKey = "user_name"

// LoadConfig reads the config file and returns the necessary
// configuration to the application.
func LoadConfig() (*togglsheet.Config, error) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configFile := path.Join(usr.HomeDir, defaultConfigFileName)
	f, err := os.Open(configFile)

	if err != nil {
		return nil, err
	}

	configValues := make(map[string]string)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		configValue := strings.Split(scanner.Text(), "=")
		configValues[configValue[0]] = configValue[1]
	}

	return &togglsheet.Config{
		APIToken:    configValues[apiTokenKey],
		WorkspaceID: configValues[workspaceIDKey],
		UserName:    configValues[usernameKey],
	}, nil
}
