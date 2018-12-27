package http_test

import toggl_export "github.com/nylo-andry/toggl-export"

var timeEntryTitle = &toggl_export.TimeEntryTitle{"test time entry"}
var timeEntry = &toggl_export.TimeEntry{timeEntryTitle, 1.0}
var projectTitle = &toggl_export.ProjectTitle{"test time entry"}
var validProjectEntry = &toggl_export.ProjectEntry{projectTitle, []*toggl_export.TimeEntry{timeEntry}}
var validTimesheet = &toggl_export.Timesheet{
	[]*toggl_export.ProjectEntry{validProjectEntry},
}
