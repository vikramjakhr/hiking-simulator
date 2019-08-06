package simulator

import (
	"github.com/vikramjakhr/hiking-simulator/parser"
	"testing"
)

func Test(t *testing.T) {
	input := &parser.SimulationInput{
		Bridges: []parser.Bridge{
			{Name: "X", Length: 100, AdditionalHikers: []parser.Hiker{
				{"A", 20},
			}},
			{Name: "Y", Length: 200},
			{Name: "Z", Length: 300, AdditionalHikers: []parser.Hiker{
				{"Q", 10},
				{"P", 20},
			}},
		},
		Hikers: []parser.Hiker{
			{Name: "P", Speed: 10},
			{Name: "Q", Speed: 20},
			{Name: "R", Speed: 50},
		},
	}
	response := Hike(input)

	if len(response.BAndTimes) != 3 {
		t.Errorf("Expected %d fastest response, found %d", len(input.Bridges), len(response.BAndTimes))
	}

	if response.BAndTimes[0].Name != "X" || response.BAndTimes[0].Time != 24 {
		t.Errorf("Expected %f fastest time for bridge %s, Found %f", float64(24), "X", response.BAndTimes[0].Time)
	}

	if response.BAndTimes[1].Name != "Y" || response.BAndTimes[1].Time != 34 {
		t.Errorf("Expected %f fastest time for bridge %s, Found %f", float64(34), "Y", response.BAndTimes[1].Time)
	}

	if response.BAndTimes[2].Name != "Z" || response.BAndTimes[2].Time != 108 {
		t.Errorf("Expected %f fastest time for bridge %s, Found %f", float64(108), "Z", response.BAndTimes[2].Time)
	}
}
