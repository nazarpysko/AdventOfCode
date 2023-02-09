package main

import "testing"

func TestAddPath(t *testing.T) {
	tests := []struct {
		name                string
		path                string
		expectedCoordinates []Coordinate
	}{
		{"Right path", "498,4 -> 498,6",
			[]Coordinate{{498, 4}, {498, 5}, {498, 6}}},
		{"Upper path", "498,6 -> 496,6",
			[]Coordinate{{498, 6}, {497, 6}, {496, 6}}},
		{"Up - right path", "498,4 -> 498,6 -> 496,6",
			[]Coordinate{{498, 4}, {498, 5}, {498, 6}, {497, 6}, {496, 6}}},
		{"Down - left path", "503,4 -> 502,4 -> 502,6 -> 500,6",
			[]Coordinate{{503, 4}, {502, 4}, {502, 5}, {502, 6}, {501, 6}, {500, 6}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			cave := (&Cave{}).newCave()
			cave.addPath(test.path)
			for _, point := range test.expectedCoordinates {
				value, exists := cave.structure[point]
				if !exists {
					t.Errorf("Point %s is expected to be defined in cave", point)
				}

				if value != rock {
					t.Errorf("Point %s is expected to be a rock", value)
				}
			}
		})
	}
}

func TestGetCoordinate(t *testing.T) {
	tests := []struct {
		name       string
		strPoint   string
		coordinate Coordinate
	}{
		{"simple test", "456,55", Coordinate{456, 55}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := getCoordinate(test.strPoint)
			if got != test.coordinate {
				t.Errorf("%v should be equal to %v, but it is not", got, test.coordinate)
			}
		})
	}
}
