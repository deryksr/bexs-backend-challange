package service

import (
	"reflect"
	"strconv"
	"testing"
)

func TestGetBestRoute(test *testing.T) {
	testGroup := []struct {
		name           string
		source         string
		target         string
		graphInput     [][]string
		expectedResult Route
	}{
		{
			"The best route is the only route",
			"A",
			"B",
			[][]string{
				{"A", "B", "5"},
			},
			Route{
				Paths: []string{"A - B"},
				Cost:  5,
			},
		},
	}

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			graph := GetGraphInstance()
			for _, line := range testCase.graphInput {
				origin := City{line[0], false, nil}
				destination := City{line[1], false, nil}
				cost, _ := strconv.Atoi(line[2])
				graph.AddRoad(&origin, &destination, cost)
			}

			bestRoute := GetBestRoute(testCase.source, testCase.target)

			if !reflect.DeepEqual(bestRoute, testCase.expectedResult) {
				test.Errorf(
					"TestGetBestRoute() Expected: <%v> but got <%v>",
					testCase.expectedResult,
					bestRoute,
				)
			}
			CleanGraph()
		})
	}
}
