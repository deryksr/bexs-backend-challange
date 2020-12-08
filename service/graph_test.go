package service

import "testing"

func TestGraphAddCity(test *testing.T) {
	testGroup := []struct {
		name  string
		input []*City
	}{
		{
			"Do not adds any city",
			nil,
		},
		{
			"Add a city without any connection",
			[]*City{&City{"B", false, nil}},
		},
		{
			"Add a city with connections",
			[]*City{
				&City{"B", false, []*Road{
					&Road{&City{"D", false, nil}, 10},
					&Road{&City{"E", false, nil}, 20},
				}},
			},
		},
	}

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			graph := GetGraphSingleton()
			for _, city := range testCase.input {
				graph.AddCity(city)
			}

			if len(graph.Cities) != len(testCase.input) {
				test.Errorf(
					"TestGraphAddCity() expected: <%d> itens on graph but got: %d",
					len(testCase.input),
					len(graph.Cities),
				)
			}
			CleanGraph()
		})
	}

}

func TestGraphAddRoad(test *testing.T) {

}
