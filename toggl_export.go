package toggl_export

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

type Timesheet struct {
	ProjectEnties []*ProjectEntry `json:"data"`
}

type Config struct {
	ApiToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
	UserName    string `toml:"user_name"`
}

type TogglService interface {
	GetTimeSheet(start, end string) (*Timesheet, error)
}

type TimesheetExporter interface {
	Export(timesheet Timesheet) (string, error)
}
