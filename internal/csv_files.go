package internal

import (
	"encoding/csv"
	"os"
)

func ReadCSVFile(inFile string) ([][]string, error) {
	f, err := os.Open(inFile)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}
	return records, nil
}

func WriteCSVFile(outFile string, outCSV [][]string) error {
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(outCSV)

	if err := w.Error(); err != nil {
		return err
	}

	return nil
}
