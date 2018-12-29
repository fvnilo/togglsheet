package test

import toggl_export "github.com/nylo-andry/toggl-export"

var TimeEntryTitle = &toggl_export.TimeEntryTitle{"test time entry"}
var TimeEntry = &toggl_export.TimeEntry{TimeEntryTitle, 3600000.0}
var ProjectTitle = &toggl_export.ProjectTitle{"project time entry"}
var ValidProjectEntry = &toggl_export.ProjectEntry{ProjectTitle, []*toggl_export.TimeEntry{TimeEntry}}
var ValidTimesheet = &toggl_export.Timesheet{
	[]*toggl_export.ProjectEntry{ValidProjectEntry},
}
