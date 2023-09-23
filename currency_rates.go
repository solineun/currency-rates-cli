package main

import (
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
	currency, err := m.GetCurrency(options)
	if err != nil {
		log.Fatal(err)
	}
	v.PrintResult(currency)
}