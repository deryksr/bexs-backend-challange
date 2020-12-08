package service

import (
	"errors"
	"reflect"
	"testing"
)

func TestReadCsvFile(test *testing.T) {
	testGroup := []struct {
		name         string
		args         string
		resultWanted [][]string
		errorWanted  error
	}{
		{
			"Read input.sv file with success",
			"../input.csv",
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
		})
	}

}
