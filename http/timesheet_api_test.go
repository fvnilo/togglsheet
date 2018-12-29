package http_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	toggl_export "github.com/nylo-andry/toggl-export"
	api "github.com/nylo-andry/toggl-export/http"
	"github.com/nylo-andry/toggl-export/test"
)

var config = &toggl_export.Config{
	ApiToken:    "api_token",
	WorkspaceID: "workspace_id",
	UserName:    "user_name",
}

func TestGetTimeSheet_ValidTimeSheet(t *testing.T) {
	server := newTestServer(test.ValidTimesheet)

	timesheetApi := api.NewTimesheetAPI(server.URL, config, server.Client())
	timesheet, err := timesheetApi.GetTimeSheet("", "")

	if err != nil {
		t.Errorf("Got an error while getting the timesheet: %v", err)
	}

	if len(timesheet.ProjectEnties) == 0 {
		t.Error("Didn't get any timesheet project entries")
	}

	timeEntryTime := timesheet.ProjectEnties[0].TimeEntries[0].Time

	if timesheet.ProjectEnties[0].TimeEntries[0].Time != 3600000 {
		t.Errorf("Expected a time entry of %v but got %v", 3600000, timeEntryTime)
	}
}

func TestGetTimeSheet_EmptyTimeSheet(t *testing.T) {
	server := newTestServer(&toggl_export.Timesheet{})

	timesheetApi := api.NewTimesheetAPI(server.URL, config, server.Client())
	timesheet, err := timesheetApi.GetTimeSheet("", "")

	if err != nil {
		t.Errorf("Got an error while getting the timesheet: %v", err)
	}

	if len(timesheet.ProjectEnties) != 0 {
		t.Error("Didn't get an empty timesheet")
	}
}

func newTestServer(data *toggl_export.Timesheet) *httptest.Server {
	b, _ := json.Marshal(data)

	fmt.Println(data)

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(b)
	}))
}
