package simulator

import (
	"fmt"
	"github.com/vikramjakhr/hiking-simulator/parser"
)

type BridgeFastestTime struct {
	Name string
	Time float64
}

type Response struct {
	BAndTimes []BridgeFastestTime
	TotalTime float64
}

// Actual logic for hiking with shortest time
func Hike(input *parser.SimulationInput) (*Response, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}
	bfts := make([]BridgeFastestTime, len(input.Bridges))

	var totalTime float64
	// get fastest among the all hikers, rest are remaining
	fastestHiker, remainingHikers := fastestOne(input.Hikers)

	// Iterate over all bridges
	for pos, bridge := range input.Bridges {
		var fastestTime float64 = 0
		var crossingTime, returnTime float64

		// Iterate over all hikers except fastest
		for _, hiker := range remainingHikers {
			crossingTime, returnTime = letsCrossTheBridge(fastestHiker, hiker, bridge)
			fastestTime = fastestTime + crossingTime + returnTime
		}

		// Iterate over additional hikers if any else subtract return time for last hiker
		if bridge.AdditionalHikers != nil {
			for index, hiker := range bridge.AdditionalHikers {
				crossingTime, returnTime = letsCrossTheBridge(fastestHiker, hiker, bridge)

				// Check if its last, if so than set return time to 0
				if index == len(bridge.AdditionalHikers)-1 {
					returnTime = 0
				}

				// Calculate fastest time
				fastestTime = fastestTime + crossingTime + returnTime
			}
		} else {
			fastestTime = fastestTime - returnTime
		}

		bfts[pos] = BridgeFastestTime{Name: bridge.Name, Time: fastestTime}
		fmt.Printf("Fastest time for crossing bridge %s: %f\n", bridge.Name, fastestTime)

		// Calculate total time
		totalTime = totalTime + fastestTime
	}
	fmt.Printf("Total time for crossing: %f\n", totalTime)
	return &Response{BAndTimes: bfts, TotalTime: totalTime}, nil
}

func validate(input *parser.SimulationInput) error {
	if len(input.Hikers) < 1 {
		return fmt.Errorf("Atleast one hiker needed")
	}
	return nil
}

// Cross the bridge
func letsCrossTheBridge(fastest, other parser.Hiker, bridge parser.Bridge) (float64, float64) {
	crossingTime := bridge.Length / other.Speed
	returnTime := bridge.Length / fastest.Speed
	return crossingTime, returnTime
}

// Find the fastest hiker and return fastest and except faster hikers
func fastestOne(hikers []parser.Hiker) (parser.Hiker, []parser.Hiker) {
	// Lets assume first hiker is fastest
	fastest := hikers[0]

	// Except faster
	remaining := make([]parser.Hiker, len(hikers)-1)

	// Iterate over hikers except faster
	for key, hiker := range hikers[1:] {
		// if hiker is faster than swap with faster, else add it to remaining
		if fastest.Speed < hiker.Speed {
			remaining[key] = fastest
			fastest = hiker
		} else {
			remaining[key] = hiker
		}
	}
	return fastest, remaining
}
