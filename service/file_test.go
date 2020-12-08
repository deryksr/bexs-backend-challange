package service

import (
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
	}

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			resultGot, err := ReadCsvFile(testCase.args)
			if err != nil {
				test.Errorf("ReadCsvFile got an error: %v", err)
			}
			if !reflect.DeepEqual(resultGot, testCase.resultWanted) {
				test.Errorf(
					"ReadCsvFile got an unexpected result - want: <%v> but got: <%v>",
					resultGot,
					testCase.resultWanted,
				)
			}
		})
	}

}
