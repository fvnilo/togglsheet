package toggl_export

// ProjectTitle represents a project title
type ProjectTitle struct {
	Project string `json:"project"`
}

// TimeEntryTitle represents the title of a time entry.
type TimeEntryTitle struct {
	Name string `json:"time_entry"`
}

// TimeEntry represents the details of a given task in toggl
type TimeEntry struct {
	Title *TimeEntryTitle `json:"title"`
	Time  int             `json:"time"`
}

// ProjectEntry repesents the time entries for one project
type ProjectEntry struct {
	Title       *ProjectTitle `json:"title"`
	TimeEntries []*TimeEntry  `json:"items"`
}

// Timesheet represents the data for a toggl workspace
type Timesheet struct {
	ProjectEntries []*ProjectEntry `json:"data"`
}

// Workspace represents a Toggl workspace
type Workspace struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Config represents the application configuration
type Config struct {
	APIToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
	UserName    string `toml:"user_name"`
}
