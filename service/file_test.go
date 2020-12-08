package service

import (
	"encoding/csv"
	"errors"
	"os"
	"reflect"
	"testing"
)

func createTempCsvFile(text [][]string) {
	file, _ := os.OpenFile("temp.csv", os.O_RDWR|os.O_CREATE, 0775)
	defer file.Close()

	csvWriter := csv.NewWriter(file)
	for _, line := range text {
		csvWriter.Write(line)
	}
	csvWriter.Flush()
}

func deleteTempCsvFile() {
	os.Remove("temp.csv")
}

func TestReadCsvFile(test *testing.T) {
	testGroup := []struct {
		name         string
		args         string
		resultWanted [][]string
		errorWanted  error
	}{
		{
			"Read input.sv file with success",
			"temp.csv",
			[][]string{
				{"GRU", "BRC", "10"},
				{"BRC", "SCL", "5"},
				{"GRU", "CDG", "75"},
				{"GRU", "SCL", "20"},
				{"GRU", "ORL", "56"},
				{"ORL", "CDG", "5"},
				{"SCL", "ORL", "20"},
			},
			nil,
		},
		{
			"Try to read a nonexistent file",
			"nonexistent-file.csv",
			[][]string{},
			errors.New("open nonexistent-file.csv: no such file or directory"),
		},
	}

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			createTempCsvFile(testCase.resultWanted)

			resultGot, err := ReadCsvFile(testCase.args)
			if err != nil && err.Error() != testCase.errorWanted.Error() {
				test.Errorf(
					"ReadCsvFile() got an unexpected error: - want: <%v> but got: <%v>",
					testCase.errorWanted,
					err,
				)
			}
			if !reflect.DeepEqual(resultGot, testCase.resultWanted) {
				test.Errorf(
					"ReadCsvFile() got an unexpected result - want: <%v> but got: <%v>",
					testCase.resultWanted,
					resultGot,
				)
			}

			deleteTempCsvFile()
		})
	}
}

func TestWriteCsvFile(test *testing.T) {
	testGroup := []struct {
		name        string
		fileName    string
		fileContent [][]string
		text        []string
		errorWanted error
	}{
		{
			"Write data into csv file with success",
			"temp.csv",
			[][]string{},
			[]string{"OR", "DE", "11"},
			nil,
		},
		{
			"Append data into csv file with success",
			"temp.csv",
			[][]string{
				{"A", "C", "1"},
				{"B", "A", "5"},
				{"C", "B", "3"},
				{"B", "D", "2"},
			},
			[]string{"F", "H", "90"},
			nil,
		},
	}

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			createTempCsvFile(testCase.fileContent)

			readBefore, _ := ReadCsvFile(testCase.fileName)
			err := WriteCsvFile(testCase.fileName, testCase.text)

			if err != nil {
				test.Errorf("WriteCsvFile() got an unexpected error: %v", err)
			}

			readAfter, _ := ReadCsvFile(testCase.fileName)

			expected := append(readBefore, testCase.text)
			if !reflect.DeepEqual(expected, readAfter) {
				test.Errorf(
					"WriteCsvFile() got an unexpected result - want: <%v> but got: <%v>",
					expected,
					readAfter,
				)
			}

			deleteTempCsvFile()
		})
	}
}
