package service

import "testing"

func TestParseCitiesToString(test *testing.T) {
	testGroup := []struct {
		name     string
		input    []*City
		expected string
	}{
		{
			"Get empty string from empty input",
			[]*City{},
			"",
		},
		{
			"Get two cities as string",
			[]*City{&City{"A", false, nil}, &City{"B", false, nil}},
			"A - B",
		},
		{
			"Get two cities as string",
			[]*City{
				&City{"A", false, nil},
				&City{"B", false, nil},
				&City{"C", false, nil},
				&City{"D", false, nil},
			},
			"A - B - C - D",
		},
	}

	for _, testCase := range testGroup {
		test.Run(testCase.name, func(test *testing.T) {
			result := parseCitiesToString(testCase.input)

			if result != testCase.expected {
				test.Errorf("Expected <%v> but got <%v>", testCase.expected, result)
			}
		})
	}
}
