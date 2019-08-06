package parser

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var inputFileToLoad = "input.yaml"

type Bridge struct {
	Name             string  `yaml:"name"`
	Length           float64 `yaml:"length"`
	AdditionalHikers []Hiker `yaml:"additionalHikers"`
}

type Hiker struct {
	Name  string  `yaml:"name"`
	Speed float64 `yaml:"speed"`
}

type SimulationInput struct {
	Bridges []Bridge `yaml:"bridges"`
	Hikers  []Hiker  `yaml:"hikers"`
}

// Return new Simulation
func NewSimulationInput() (*SimulationInput) {
	return &SimulationInput{}
}

// Load simulation YAML file
func (input *SimulationInput) LoadSimulationInput() error {
	file, err := ioutil.ReadFile(inputFileToLoad)
	if err != nil {
		return fmt.Errorf("Error while loading input. Error: %v ", err)
	}
	err = yaml.Unmarshal(file, input)
	if err != nil {
		return fmt.Errorf("Error while parsing input YAML. Error: %v ", err)
	}
	return nil
}
