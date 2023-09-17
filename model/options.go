package model

import "fmt"

type Date struct {
	Year int
	Month int
	Day int
}

type InputOptions struct {
	Code string
	Date
}

func (d Date) cbrString() string {
	return fmt.Sprintf("%s/%s/%d", format(d.Day), format(d.Month), d.Year)
}

func format(i int) string {
	if i < 10 {
		return fmt.Sprintf("0%d", i)
	}
	return fmt.Sprintf("%d", i)
}