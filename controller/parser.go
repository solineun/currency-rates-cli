package controller

import (
	"github.com/golang-module/carbon/v2"
)
type Date struct {
	Year int
	Month int
	Day int
}

func today() Date {
	y, m, d := carbon.Now().Date()
	return Date {
		Year: y,
		Month: m,
		Day: d,
	}
}

type InputOptions struct {
	Code string
	Date
}

func ParseInput(input []string) (InputOptions, error){
	var code string
	var date Date
	code, err := parseCode(input[0])
	if err != nil {
		return InputOptions{}, err
	}
	if len(input) == 1 {
		date = today()
	} else {
		date, err = parseDate(input[1])
	}
	if err != nil {
		return InputOptions{}, err
	}
	return  InputOptions {
		code,
		date, 		
	}, nil
}

func parseCode(code string) (string, error){
	err := checkCode(code)
	if err != nil {
		return "", err
	}
	return code[7:], nil
}

func parseDate(date string) (Date, error) {
	err := checkDate(date)
	if err != nil {
		return Date{}, err
	}
	date = date[7:]
	return getDateFromString(date), nil
}

func getDateFromString(str string) Date {
	y, m, d := carbon.Parse(str).Date()
	return Date{
		y, 
		m,
		d,
	}
}