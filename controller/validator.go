package controller

import (
	"errors"
	"regexp"
)

func checkCode(code string) error {
	var pattern = `(?:\s|^)(?P<primary>--code[=]\S+)(?:\s|$)`
	m, err := matchFlagPattern(code, pattern)
	if err != nil {
		return err
	}
	if !m {
		return errors.New("incorrect code flag syntax")
	}
	if len(code) != 10 {
		return errors.New("incorrect currency code")
	}
	return nil
}

func checkDate(date string) error {
	var pattern =  `(?:\s|^)(?P<primary>--date[=]\S+)(?:\s|$)`
	m, err := matchFlagPattern(date, pattern)
	if err != nil {
		return err
	}
	if !m {
		return errors.New("incorrect date flag syntax")
	}
	if len(date) != 17 {
		return errors.New("incorrect currency date")
	}
	return nil
}

func matchFlagPattern(flag, pattern string) (bool, error) {
	m, err := regexp.MatchString(pattern, flag)
	if err != nil {
		return m, err
	}
	return m, nil
}