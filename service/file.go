package service

import (
	"encoding/csv"
	"fmt"
	"os"
)

func ReadCsvFile(fileName string) ([][]string, error) {
	result := [][]string{}
	file, err := os.Open(fileName)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Cannot open the file %s | Reason: %s", fileName, err)
		return result, err
	}

	csvReader := csv.NewReader(file)
	for {
		line, err := csvReader.Read()
		if err != nil {
			break
		}
		result = append(result, line)
	}
	return result, nil
}
