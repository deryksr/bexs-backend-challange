package service

import (
	"encoding/csv"
	"fmt"
	"os"
)

var csvFileName string

func GetCsvFileName() string {
	return csvFileName
}

func ReadCsvFile(fileName string) ([][]string, error) {
	result := [][]string{}
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR, 0775)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Cannot open the file %s | Reason: %s", fileName, err)
		return result, err
	}

	csvReader := csv.NewReader(file)
	csvWriter := csv.NewWriter(file)
	csvFileName = fileName

	for {
		line, err := csvReader.Read()
		if err != nil {
			break
		}
		result = append(result, line)
	}
	csvWriter.Write([]string{""}) // adds a new endline into csv file
	csvWriter.Flush()
	return result, nil
}

func WriteCsvFile(fileName string, text []string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0775)
	defer file.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: Cannot open the file %s | Reason: %s", fileName, err)
		return err
	}

	csvWriter := csv.NewWriter(file)
	csvWriter.Write(text)
	csvWriter.Flush()
	return nil
}
