package main

import (
	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/configuration"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	config := configuration.LoadConfiguration()
	doneChannels := make([]chan bool, len(config.TaxRates))
	errorChannels := make([]chan error, len(config.TaxRates))

	for index, taxRate := range config.TaxRates {
		doneChannels[index] = make(chan bool)
		errorChannels[index] = make(chan error)
		iom := getIOManager(config.IoManagerType, taxRate)
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRate)
		go priceJob.Process(doneChannels[index], errorChannels[index])
	}

	for index := range config.TaxRates {
		select {
		case err := <-errorChannels[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChannels[index]:
			fmt.Println("Done!")
		}
	}
}

func getIOManager(IOManagerType string, taxRate float64) iomanager.IOManager {
	switch IOManagerType {
	case "cmd":
		return cmdmanager.New()
	case "file":
		return filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
	default:
		return cmdmanager.New()
	}
}
