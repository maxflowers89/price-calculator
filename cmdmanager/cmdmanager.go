package cmdmanager

import (
	"fmt"
)

type CmdManager struct{}

func (cmd CmdManager) ReadLines() ([]string, error) {
	fmt.Println("Please enter your prices and confirm with ENTER")

	var prices []string

	for {
		var price string
		fmt.Print("Price: ")
		_, _ = fmt.Scan(&price)

		if price == "0" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CmdManager) WriteData(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CmdManager {
	return CmdManager{}
}
