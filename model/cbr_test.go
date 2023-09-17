package model

import (
	"fmt"
	"testing"
)

func TestMakeRequestNotNil(t *testing.T) {
	var tests = []struct {
		name string
		url string
	}{
		{"yandex", "https://ya.ru/"},
		{"google", "https://www.google.com/"},
		{"cbr", "https://www.cbr.ru/scripts/XML_daily.asp?date_req=",},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := makeRequest(tt.url)
			if err != nil {
				t.Errorf(err.Error())
			}
			if res == nil {
				t.Errorf("got nil as response body")
			}
		})
	}
}

func TestConcatUrlWithDate(t *testing.T) {
	var tests = []struct {
		Date
		want string
	}{
		{
			Date{2022, 8, 6}, 
			"https://www.cbr.ru/scripts/XML_daily.asp?date_req=06/08/2022",
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.Date), func(t *testing.T) {
			res := concatUrlWithDate(tt.Date)
			if res != tt.want {
				t.Errorf("Error in date formating")
			}
		})
	}
}