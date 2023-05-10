package main

import (
	"testing"
)

func TestGetSensorAndBeacon(t *testing.T) {
	tests := []struct {
		name   string
		line   string
		sensor Coordinate
		beacon Coordinate
	}{
		{"Low numbers", "Sensor at x=2, y=18: closest beacon is at x=-2, y=15", Coordinate{2, 18}, Coordinate{-2, 15}},
		{"High numbers", "Sensor at x=3844106, y=3888618: closest beacon is at x=3225436, y=4052707\n", Coordinate{3844106, 3888618}, Coordinate{3225436, 4052707}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSensor, gotBeacon := getSensorAndBeacon(test.line)
			if gotSensor != test.sensor {
				t.Errorf("Want %d but got %d", test.sensor, gotSensor)
			}

			if gotBeacon != test.beacon {
				t.Errorf("Want %d but got %d", test.beacon, gotBeacon)
			}
		})
	}
}
