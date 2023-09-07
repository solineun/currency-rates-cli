package controller

import (
	"errors"
	"regexp"
	"github.com/golang-module/carbon/v2"
)

type InputOptions struct {
	Code string
	Date string
}

func ParseInput(input []string) (InputOptions, error){
	c, err := parseCode(input[0])
	if err != nil {
		return InputOptions{}, err
	}
	if len(input) == 1 {
		return InputOptions {
			c,
			carbon.Now().ToDateString(), 		
		}, nil
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
	err := validateCode(code)
	if err != nil {
		return "", err
	}
	return code[7:], nil
}

func parseDate(date string) (string, error) {
	m, err := regexp.MatchString("--date=", date)
	if err != nil {
		return "", err
	}
	// --date=2022-10-08
	if !m || len(date) != 17 {
		return "", errors.New("incorrect date flag syntax")
	}
	return date[7:], nil
}