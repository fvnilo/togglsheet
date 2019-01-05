package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/nylo-andry/togglsheet/config"
	"github.com/nylo-andry/togglsheet/export"
	"github.com/nylo-andry/togglsheet/httpclient"
)

func main() {
	startDate := flag.String("start", "", "The first day to start the report from")
	endDate := flag.String("end", "", "The last day to export from")

	flag.Parse()

	config, err := config.ReadConfig("config.toml")

	if err != nil {
		log.Fatalf("Could not read config: %v", err)
	}

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
