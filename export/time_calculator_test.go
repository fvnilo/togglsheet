package export_test

import (
	"strconv"
	"testing"

	"github.com/nylo-andry/togglsheet/export"
)

var timeValuesTest = []struct {
	in  int
	out float64
}{
	{540000, 0.25},
	{3600000, 1},
	{4140000, 1.25},
	{1188000, 0.25},
	{1620000, 0.5},
}

func TestCalculateTimeEntry(t *testing.T) {
	for _, tt := range timeValuesTest {
		t.Run(strconv.Itoa(tt.in), func(t *testing.T) {
			v := export.CalculateTimeEntry(tt.in)
			if v != tt.out {
				t.Errorf("got %v, want %v", v, tt.out)
			}
		})
	}
}
