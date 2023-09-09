package model

type Date struct {
	Year int
	Month int
	Day int
}

type InputOptions struct {
	Code string
	Date
}

type Currency struct {
	
}

func GetCurrency(options InputOptions) (Currency, error) {
	
	return Currency{}, nil
}

func getCurrencyFromXml() {

}