package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nylo-andry/togglsheet/config"
	"github.com/nylo-andry/togglsheet/httpclient"
)

func main() {
	config, err := config.ReadConfig("config.toml")

	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	client := &http.Client{}
	workspaceAPI := httpclient.NewWorkspaceAPI("https://toggl.com", config, client)
	w, err := workspaceAPI.GetWorkspaces()

	if err != nil {
		log.Fatalf("Could not get workspace data: %v", err)
	}

	fmt.Println("Here are the known workspaces for your account:")

	for _, w := range w.Workspaces {
		fmt.Printf("* ID: %v, Name: %v\n", w.ID, w.Name)
	}

	fmt.Println("")
	fmt.Println("Identify the one you want to export and put its ID in the config file.")
}
