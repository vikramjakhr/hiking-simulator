package main

import (
	"github.com/vikramjakhr/hiking-simulator/parser"
	"github.com/vikramjakhr/hiking-simulator/simulator"
	"log"
)

func main() {
	// Read simulation input
	simulatorInput := parser.NewSimulationInput()
	err := simulatorInput.LoadSimulationInput()
	if err != nil {
		log.Fatal("E! " + err.Error())
	}

	// lets go hiking
	simulator.Hike(simulatorInput)

}
