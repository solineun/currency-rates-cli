package view

import (
	"errors"
	"os"
)

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

func PrintResult([]string) {

}