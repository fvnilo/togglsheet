package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/nylo-andry/toggl-export/config"
	api "github.com/nylo-andry/toggl-export/http"
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
	timesheetAPI := api.NewTimesheetAPI("https://toggl.com", config, client)

	timesheet, err := timesheetAPI.GetTimeSheet(*startDate, *endDate)

	if err != nil {
		log.Fatalf("Could not get timesheet data: %v", err)
	}

	data := make([][]string, 0)
	var total float64

	for _, projectEntry := range timesheet.ProjectEnties {
		for _, timeEntry := range projectEntry.TimeEntries {
			hours := float64(timeEntry.Time) / 3600000
			roundedHours := math.Round(hours*4) / 4

			if roundedHours == 0 {
				roundedHours = 0.25
			}

			total += roundedHours
			data = append(data, []string{config.UserName, projectEntry.Title.Project, "Production", timeEntry.Title.Name, strconv.FormatFloat(roundedHours, 'f', 2, 64)})
		}
	}

	file, err := os.Create("result.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := csv.NewWriter(file)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}

	fmt.Printf("Total time: %vh\n", total)
}
