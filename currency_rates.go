package main

import (
	"fmt"
	"log"
	v "currencyrates/view"
	c "currencyrates/controller"
	m "currencyrates/model"
)



func main() {
	input, err := v.ScanInput()
	if err != nil {
		log.Fatal(err)
	}
	options, err := c.ParseInput(input)
	if err != nil {
		log.Fatal(err)
	}
	currRates, err := m.GetCurrencyRates(options)
	fmt.Println(currRates)
}