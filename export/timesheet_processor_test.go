package export_test

import "testing"
import "github.com/nylo-andry/togglsheet/export"
import "github.com/nylo-andry/togglsheet/internal/pkg/test"

func TestProcessTimesheet(t *testing.T) {
	data := export.ProcessTimesheet(test.ValidTimesheet, "test")

	if len(data) != 1 {
		t.Error("Received no data while processing.")
	}

	if data[0][0] != "test" {
		t.Errorf("Username column does match. Got %v instead of %v", data[0][0], "test")
	}

	if data[0][1] != "project time entry" {
		t.Errorf("Project name column does match. Got %v instead of %v", data[0][1], "project time entry")
	}

	if data[0][2] != "Production" {
		t.Errorf("Entry type column does match. Got %v instead of %v", data[0][2], "Production")
	}

	if data[0][3] != "test time entry" {
		t.Errorf("Entry title column does match. Got %v instead of %v", data[0][3], "test time entry")
	}

	if data[0][4] != "1.00" {
		t.Errorf("Time column does match. Got %v instead of %v", data[0][4], "1.00")
	}
}
