package export

import (
	"strconv"

	toggl_export "github.com/nylo-andry/toggl-export"
)

type CsvData [][]string

func ProcessTimesheet(timesheet *toggl_export.Timesheet, username string) CsvData {
	data := make([][]string, 0)

	for _, projectEntry := range timesheet.ProjectEnties {
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
