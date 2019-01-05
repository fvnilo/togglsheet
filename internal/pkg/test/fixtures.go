package test

import "github.com/nylo-andry/togglsheet"

var timeEntryTitle = &togglsheet.TimeEntryTitle{
	Name: "test time entry",
}

var timeEntry = &togglsheet.TimeEntry{
	Title: timeEntryTitle,
	Time:  3600000.0,
}

var projectTitle = &togglsheet.ProjectTitle{
	Project: "project time entry",
}

var validProjectEntry = &togglsheet.ProjectEntry{
	Title:       projectTitle,
	TimeEntries: []*togglsheet.TimeEntry{timeEntry},
}

// ValidTimesheet represents a minimal valid timesheet entry
var ValidTimesheet = &togglsheet.Timesheet{
	ProjectEntries: []*togglsheet.ProjectEntry{validProjectEntry},
}
