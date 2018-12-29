package export

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func ExportCSV(data CsvData) (string, error) {
	t := time.Now()
	suffix := t.Format("20060102-150405")
	fileName := fmt.Sprintf("result-%v.csv", suffix)
	file, err := os.Create(fileName)

	if err != nil {
		return "", err
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		return "", err
	}

	return fileName, err
}
