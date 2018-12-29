package test

import toggl_export "github.com/nylo-andry/toggl-export"

var timeEntryTitle = &toggl_export.TimeEntryTitle{
	Name: "test time entry",
}

var timeEntry = &toggl_export.TimeEntry{
	Title: timeEntryTitle,
	Time:  3600000.0,
}

var projectTitle = &toggl_export.ProjectTitle{
	Project: "project time entry",
}

var validProjectEntry = &toggl_export.ProjectEntry{
	Title:       projectTitle,
	TimeEntries: []*toggl_export.TimeEntry{timeEntry},
}

// ValidTimesheet represents a minimal valid timesheet entry
var ValidTimesheet = &toggl_export.Timesheet{
	ProjectEntries: []*toggl_export.ProjectEntry{validProjectEntry},
}
