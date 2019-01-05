package export

import (
	"strconv"

	"github.com/nylo-andry/togglsheet"
)

// CSVData represents the data of the timesheet in CSV.
type CSVData [][]string

// ProcessTimesheet converts timesheet data (from Toggl) to a CSV format
func ProcessTimesheet(timesheet *togglsheet.Timesheet, username string) CSVData {
	data := make([][]string, 0)

	for _, projectEntry := range timesheet.ProjectEntries {
		for _, timeEntry := range projectEntry.TimeEntries {
			time := CalculateTimeEntry(timeEntry.Time)
			data = append(data, []string{
				username,
				projectEntry.Title.Project,
				"Production",
				timeEntry.Title.Name,
				strconv.FormatFloat(time, 'f', 2, 64)})
		}
	}

	return data
}
