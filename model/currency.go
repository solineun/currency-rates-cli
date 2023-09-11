package model

import (
	"encoding/xml"
	"errors"
)

type Date struct {
	Year int
	Month int
	Day int
}

type InputOptions struct {
	Code string
	Date
}

type CurrencyRates struct {
	Currencies []Currency `xml:"Valute"`
}

type Currency struct {
	NumCode int `xml:"NumCode"`
	CharCode string `xml:"CharCode"`
	Nominal int `xml:"Nominal"`
	Name string `xml:"Name"`
	Value string `xml:"Value"`
}

func GetCurrency(options InputOptions) (Currency, error) {
	cr, err := getCurrencyRates(options.Date)
	if err != nil {
		return Currency{}, err
	}
	curr, err := findCurrByCode(options.Code, cr.Currencies)
	return curr, err
}

func getCurrencyRates(date Date) (CurrencyRates, error) {
	ratesXml, err := getXml(date)
	if err != nil {
		return CurrencyRates{}, err
	}
	cr, err := getRatesFromXml(ratesXml)
	return cr, err
}

func getRatesFromXml(ratesXml []byte) (CurrencyRates, error){
	cr := new(CurrencyRates)
	err := xml.Unmarshal(ratesXml, cr)
	return *cr, err
}

func findCurrByCode(code string, currencies []Currency) (Currency, error) {
	for _, curr := range currencies {
		if curr.CharCode == code {
			return curr, nil
		}
	}
	return Currency{}, errors.New("no such currency by this code")
}