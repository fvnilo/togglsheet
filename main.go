package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

const userAgent = "toggl-export"

type ProjectTitle struct {
	Project string `json:"project"`
}

type TimeEntryTitle struct {
	Name string `json:"time_entry"`
}

type TimeEntry struct {
	Title *TimeEntryTitle `json:"title"`
	Time  int             `json:"time"`
}

type ProjectEntry struct {
	Title       *ProjectTitle `json:"title"`
	TimeEntries []*TimeEntry  `json:"items"`
}

type Response struct {
	ProjectEnties []*ProjectEntry `json:"data"`
	// Total float64 `json:"total_grand"`
}

type Config struct {
	ApiToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
	UserName    string `toml:"user_name"`
}

func main() {
	startDate := flag.String("start", "", "The first day to start the report from")
	endDate := flag.String("end", "", "The last day to export from")

	flag.Parse()

	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://toggl.com/reports/api/v2/summary", nil)
	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(config.ApiToken, "api_token")
	q := req.URL.Query()

	q.Add("workspace_id", config.WorkspaceID)
	q.Add("since", *startDate)
	q.Add("until", *endDate)
	q.Add("user_agent", userAgent)

	req.URL.RawQuery = q.Encode()
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	responseBody := Response{}
	jsonErr := json.Unmarshal(body, &responseBody)
	if jsonErr != nil {
		panic(jsonErr)
	}

	data := make([][]string, 0)
	var total float64

	for _, projectEntry := range responseBody.ProjectEnties {
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
