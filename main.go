package main

import (
	"example.com/price-calculator/cmdmanager"
	"example.com/price-calculator/configuration"
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/iomanager"
	"example.com/price-calculator/prices"
	"fmt"
	"strconv"
)

//

func main() {
	config := configuration.LoadConfiguration()

	for _, taxRate := range config.TaxRates {
		taxRateAsFloat, _ := strconv.ParseFloat(taxRate, 64)
		iom := getIOManager(config.IoManagerType, taxRateAsFloat)
		priceJob := prices.NewTaxIncludedPriceJob(iom, taxRateAsFloat)
		err := priceJob.Process()
		if err != nil {
			fmt.Println("Could not process the job: ", err)
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
