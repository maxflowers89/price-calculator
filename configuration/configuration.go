package configuration

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Configuration struct {
	IoManagerType string    `yaml:"io-manager-type"`
	TaxRates      []float64 `yaml:"tax-rates"`
}

func LoadConfiguration() Configuration {
	data, err := os.ReadFile("configuration.yaml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var configuration Configuration
	err = yaml.Unmarshal(data, &configuration)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return configuration
}
