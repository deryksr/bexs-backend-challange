package service

import "testing"

type Fixture struct {
	name  string
	input []*City
}

func getFixture() []Fixture {
	cityA := City{"A", false, nil}
	cityB := City{"B", false, []*Road{
		{&cityA, 10},
	}}
	cityC := City{"C", false, []*Road{
		{&cityA, 10},
		{&cityB, 30},
	}}

	return []Fixture{
		{
			"Do not adds any city",
			nil,
		},
		{
			"Add a city without any connection",
			[]*City{&cityA},
		},
		{
			"Add a city with connections",
			[]*City{&cityA, &cityB, &cityC},
		},
	}
}

func TestGraphAddCity(test *testing.T) {
	testGroup := getFixture()

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

	test.Run("Do not add the same city twice", func(test *testing.T) {
		graph := GetGraphSingleton()
		inputCity := City{"B", false, nil}
		graph.AddCity(&inputCity)
		graph.AddCity(&inputCity)
		graph.AddCity(&inputCity)

		if len(graph.Cities) != 1 {
			test.Errorf(
				"TestGraphAddCity() expected: <%d> itens on graph but got: %d",
				1,
				len(graph.Cities),
			)
		}
		CleanGraph()
	})
}

func TestGraphAddRoad(test *testing.T) {
	testGroup := getFixture()

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			graph := GetGraphSingleton()
			graph.Cities = testCase.input

			newCity := City{"X", false, nil}
			for _, currentCity := range graph.Cities {
				expected := len(currentCity.Connections) + 1
				graph.AddRoad(currentCity, &newCity, 95)
				graph.AddRoad(currentCity, &newCity, 95)

				if expected != len(currentCity.Connections) {
					test.Errorf(
						"TestGraphAddRoad() expected: <%d> itens on Connections but got: %d",
						expected,
						len(currentCity.Connections),
					)
				}
			}
			CleanGraph()
		})
	}

}
