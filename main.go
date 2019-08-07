package main

import (
	"flag"
	"fmt"
	"github.com/vikramjakhr/hiking-simulator/parser"
	"github.com/vikramjakhr/hiking-simulator/simulator"
	"log"
	"os"
)

const usage = `hiking-simulator, finding shortest time for crossing a barrier.

Usage:

  hiking-simulator
`

func usageExit(rc int) {
	fmt.Println(usage)
	os.Exit(rc)
}

func main() {
	flag.Usage = func() { usageExit(0) }
	flag.Parse()

	// Read simulation input
	simulatorInput := parser.NewSimulationInput()
	err := simulatorInput.LoadSimulationInput()
	if err != nil {
		log.Fatal("E! " + err.Error())
	}

	// lets go hiking
	_, err = simulator.Hike(simulatorInput)
	if err != nil {
		log.Fatal(err.Error())
	}

}
