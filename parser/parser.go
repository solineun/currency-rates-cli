package parser

import (
	"errors"
	"os"
	"regexp"
	"github.com/golang-module/carbon/v2"
)

type InputOptions struct {
	Code string
	Date string
}

func ScanInput() ([]string, error){
	if len(os.Args) < 2 {
		return nil, errors.New("not enough input parameters")
	}
	if len(os.Args) > 3 {
		return nil, errors.New("too much input parameters")
	}
	input := os.Args[1:]
	return input, nil
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
	m, _ := regexp.MatchString("--code=", code)
	// --code=USD
	if !m || len(code) != 10  {
		return "", errors.New("incorrect code flag syntax")
	}
	return code[7:], nil
}

func parseDate(date string) (string, error) {
	m, _ := regexp.MatchString("--date=", date)
	// --date=2022-10-08
	if !m || len(date) != 17 {
		return "", errors.New("incorrect date flag syntax")
	}
	return date[7:], nil
}