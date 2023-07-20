package parser

import (
	"errors"
	"os"
	"regexp"

)

type InputOptions struct {
	Code string
	Date string
}

func ScanInput() ([]string, error){
	if len(os.Args) < 3 {
		return nil, errors.New("not enough input parameters")
	}
	input := os.Args[1:]
	return input, nil
}

func ParseInput(input []string) (InputOptions, error){
	c, err := parseCode(input[0])
	if err != nil {
		return InputOptions{}, err
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

func parseCode(input string) (string, error){
	m, _ := regexp.MatchString("--code=", input)
	// --code=USD
	if !m || len(input) != 10  {
		return "", errors.New("incorrect code flag syntax")
	}
	code := input[7:] 
	return code, nil
}

func parseDate(date string) (string, error) {
	m, _ := regexp.MatchString("--date=", date)
	if !m || len(date) != 17 {
		return "", errors.New("incorrect date flag syntax")
	}
	return date[7:], nil
}