package service

import (
	"encoding/csv"
	"fmt"
	"os"
)

var CsvFileName string

func ReadCsvFile(fileName string) ([][]string, error) {
	result := [][]string{}
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Cannot open the file %s | Reason: %s", fileName, err)
		return result, err
	}

	csvReader := csv.NewReader(file)
	CsvFileName = fileName

	for {
		line, err := csvReader.Read()
		if err != nil {
			break
		}
		result = append(result, line)
	}
	return result, nil
}

func WriteCsvFile(fileName string, text []string) error {
	file, err := os.OpenFile(CsvFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0775)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Cannot open the file %s | Reason: %s", CsvFileName, err)
		return err
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Write(text)
	csvWriter.Flush()
	return nil
}
