package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nylo-andry/togglsheet"
	"github.com/nylo-andry/togglsheet/export"
	"github.com/nylo-andry/togglsheet/httpclient"

	"github.com/spf13/cobra"
)

func NewExportCommand(config *togglsheet.Config, startDate *string, endDate *string) command {
	return func(cmd *cobra.Command, args []string) {
		client := &http.Client{}
		timesheetAPI := httpclient.NewTimesheetAPI("https://toggl.com", config, client)

		timesheet, err := timesheetAPI.GetTimeSheet(*startDate, *endDate)

		if err != nil {
			log.Fatalf("Could not get timesheet data: %v", err)
		}

		csvData := export.ProcessTimesheet(timesheet, config.UserName)
		fileName, err := export.CSV(csvData)

		if err != nil {
			log.Fatalf("Could not export file: %v", err)
		}

		fmt.Printf("Exported timesheet at: %v\n", fileName)
	}
}
