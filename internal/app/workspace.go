package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nylo-andry/togglsheet"
	"github.com/nylo-andry/togglsheet/httpclient"

	"github.com/spf13/cobra"
)

func NewWorkspaceCommand(config *togglsheet.Config) command {
	return func(cmd *cobra.Command, args []string) {
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
}
