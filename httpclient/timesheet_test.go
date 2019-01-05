package httpclient_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nylo-andry/togglsheet"
	"github.com/nylo-andry/togglsheet/httpclient"
	"github.com/nylo-andry/togglsheet/internal/pkg/test"
)

var config = &togglsheet.Config{
	APIToken:    "api_token",
	WorkspaceID: "workspace_id",
	UserName:    "user_name",
}

func TestGetTimeSheet_ValidTimeSheet(t *testing.T) {
	server := newTestServer(test.ValidTimesheet)

	timesheetApi := httpclient.NewTimesheetAPI(server.URL, config, server.Client())
	timesheet, err := timesheetApi.GetTimeSheet("", "")

	if err != nil {
		t.Errorf("Got an error while getting the timesheet: %v", err)
	}

	if len(timesheet.ProjectEntries) == 0 {
		t.Error("Didn't get any timesheet project entries")
	}

	timeEntryTime := timesheet.ProjectEntries[0].TimeEntries[0].Time

	if timesheet.ProjectEntries[0].TimeEntries[0].Time != 3600000 {
		t.Errorf("Expected a time entry of %v but got %v", 3600000, timeEntryTime)
	}
}

func TestGetTimeSheet_EmptyTimeSheet(t *testing.T) {
	server := newTestServer(&togglsheet.Timesheet{})

	timesheetApi := httpclient.NewTimesheetAPI(server.URL, config, server.Client())
	timesheet, err := timesheetApi.GetTimeSheet("", "")

	if err != nil {
		t.Errorf("Got an error while getting the timesheet: %v", err)
	}

	if len(timesheet.ProjectEntries) != 0 {
		t.Error("Didn't get an empty timesheet")
	}
}

func newTestServer(data *togglsheet.Timesheet) *httptest.Server {
	b, _ := json.Marshal(data)

	fmt.Println(data)

	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(b)
	}))
}
