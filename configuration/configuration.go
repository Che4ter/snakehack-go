package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Port string

	Color          string
	Name           string
	Taunt          string
	HeadType       HeadType
	TailType       TailType
	SecondaryColor string
}

func ParseConfiguration() (configuration Configuration, err error) {
	// Create a default configuration.
	config := Configuration{"4242", "#FF0000", "Funky Snake", "I eat you", PIXELHEAD, PIXELTAIL, "#00FF00"}

	// Open the configuration file.
	fmt.Printf("load config file: %v", "configuration.json")
	file, err := os.Open("configuration.json")
	if err != nil {
		fmt.Println("\nerror: could not open config file")
		fmt.Println("-> use default values")
		return config, err
	}

	// Parse JSON in the configuration file.
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("error: config file not valid")
		return config, err
	}
	fmt.Println(" -> parsing successful")

	return config, nil
}
