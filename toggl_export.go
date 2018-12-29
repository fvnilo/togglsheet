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
	ProjectEntries []*ProjectEntry `json:"data"`
}

type Workspace struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Config struct {
	ApiToken    string `toml:"api_token"`
	WorkspaceID string `toml:"workspace_id"`
	UserName    string `toml:"user_name"`
}
