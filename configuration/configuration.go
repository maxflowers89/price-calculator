package configuration

import (
	"github.com/magiconair/properties"
	"log"
)

type Configuration struct {
	IoManagerType string   `properties:"io-manager-type,default=file"`
	TaxRates      []string `properties:"tax-rates"`
}

func LoadConfiguration() Configuration {
	p := properties.MustLoadFile("config.properties", properties.UTF8)
	var configuration Configuration
	if err := p.Decode(&configuration); err != nil {
		log.Fatal(err)
	}

	return configuration
}
