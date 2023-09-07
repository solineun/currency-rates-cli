package controller

import (
	"github.com/golang-module/carbon/v2"
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

func ParseInput(input []string) (InputOptions, error){
	c, err := parseCode(input[0])
	if err != nil {
		return InputOptions{}, err
	}
	if len(input) == 1 {
		today := Date {
			Year: carbon.Now().Year(),
			Month: carbon.Now().Month(),
			Day: carbon.Now().Day(),
		}
		inops := InputOptions {
			Code: c,
			Date: today,
		}
		return inops, nil
	}
	d, err := parseDate(input[1])
	if err != nil {
		return InputOptions{}, err
	}
	inops := InputOptions {
		c,
		d, 		
	}
	return inops, nil
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
	return date[7:], nil
}