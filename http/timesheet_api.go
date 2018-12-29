package http

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	toggl_export "github.com/nylo-andry/toggl-export"
)

const timeSheetURL = "reports/api/v2/summary"
const userAgent = "toggl-export"

// TimesheetAPI is the client to the summary reports endpoint of Toggl.
type TimesheetAPI struct {
	baseURL    string
	config     *toggl_export.Config
	httpClient *http.Client
}

// NewTimesheetAPI returns a new instance of the TimesheetAPI.
func NewTimesheetAPI(baseURL string, config *toggl_export.Config, httpClient *http.Client) *TimesheetAPI {
	return &TimesheetAPI{baseURL, config, httpClient}
}

// GetTimeSheet returns the timesheet data from Toggl.
func (t *TimesheetAPI) GetTimeSheet(start, end string) (*toggl_export.Timesheet, error) {
	req, err := t.buildRequest(start, end)

	if err != nil {
		return nil, err
	}

	res, err := t.httpClient.Do(req)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	timesheet := &toggl_export.Timesheet{}
	err = json.Unmarshal(body, timesheet)
	if err != nil {
		return nil, err
	}

	return timesheet, nil
}

func (t *TimesheetAPI) buildRequest(start, end string) (*http.Request, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%v/%v", t.baseURL, timeSheetURL), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(t.config.APIToken, "api_token")
	q := req.URL.Query()

	q.Add("workspace_id", t.config.WorkspaceID)
	q.Add("since", start)
	q.Add("until", end)
	q.Add("user_agent", userAgent)

	req.URL.RawQuery = q.Encode()

	return req, nil
}
