package main

import (
	"log"

	"github.com/nylo-andry/togglsheet"
	"github.com/nylo-andry/togglsheet/config"
	"github.com/nylo-andry/togglsheet/internal/app"
	"github.com/spf13/cobra"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

	rootCmd := &cobra.Command{Use: "togglsheet"}
	rootCmd.AddCommand(newExportCmd(config), newWorkspaceCmd(config))
	rootCmd.Execute()
}

func newExportCmd(config *togglsheet.Config) *cobra.Command {
	var startDate, endDate string

	var exportCmd = &cobra.Command{
		Use:   "export [string to print]",
		Short: "Exports a Toggl timesheet",
		Long:  `export is used to produce a timesheet for a given period of time.`,
		Run:   app.NewExportCommand(config, &startDate, &endDate),
	}

	exportCmd.Flags().StringVarP(&startDate, "start", "s", "", "start date of the range to export [YYYY-MM-DD]")
	exportCmd.Flags().StringVarP(&endDate, "end", "e", "", "end date of the range to export [YYYY-MM-DD]")

	return exportCmd
}

func newWorkspaceCmd(config *togglsheet.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "workspace [string to print]",
		Short: "Gets the workspaces associated to the user",
		Long:  `workspace is used to identify the workspace to export timesheets from.`,
		Run:   app.NewWorkspaceCommand(config),
	}
}
