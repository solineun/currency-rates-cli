package controller

import (
	"github.com/golang-module/carbon/v2"
	"currencyrates/model"
)

type InputOptions model.InputOptions
type Date model.Date

func today() model.Date {
	y, m, d := carbon.Now().Date()
	return model.Date {
		Year: y,
		Month: m,
		Day: d,
	}
}

func ParseInput(input []string) (model.InputOptions, error){
	var code string
	var date model.Date
	code, err := parseCode(input[0])
	if err != nil {
		return model.InputOptions{}, err
	}
	if len(input) == 1 {
		date = today()
	} else {
		date, err = parseDate(input[1])
	}
	if err != nil {
		return model.InputOptions{}, err
	}
	return  model.InputOptions {
		Code: code,
		Date: model.Date(date), 		
	}, nil
}

func parseCode(code string) (string, error){
	err := checkCode(code)
	if err != nil {
		return "", err
	}
	return code[7:], nil
}

func parseDate(date string) (model.Date, error) {
	err := checkDate(date)
	if err != nil {
		return model.Date{}, err
	}
	date = date[7:]
	return getDateFromString(date), nil
}

func getDateFromString(str string) model.Date {
	y, m, d := carbon.Parse(str).Date()
	return model.Date{
		Year: y, 
		Month: m,
		Day: d,
	}
}