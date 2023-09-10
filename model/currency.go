package model

import "encoding/xml"

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
	
}

func GetCurrencyRates(options InputOptions) (CurrencyRates, error) {
	ratesXml, err := getXml(options.Date)
	if err != nil {
		return CurrencyRates{}, err
	}
	cr, err := getRatesFromXml(ratesXml)
	if err != nil{
		return CurrencyRates{}, err
	}
	return cr, nil
}

func getRatesFromXml(ratesXml []byte) (CurrencyRates, error){
	cr := new(CurrencyRates)
	err := xml.Unmarshal(ratesXml, cr)
	return *cr, err
}