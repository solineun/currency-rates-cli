package controller

import (
	"errors"
	"regexp"
)

func validateCode(code string) error {
	var pattern = `(?:\s|^)(?P<primary>--code[=]\S+)(?:\s|$)`
	m, err := matchCodePattern(code, pattern)
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

func matchCodePattern(code, pattern string) (bool, error) {
	m, err := regexp.MatchString(pattern, code)
	if err != nil {
		return m, err
	}
	return m, nil
}