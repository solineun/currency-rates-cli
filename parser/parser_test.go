package parser_test

import (
	"fmt"
	"qiwifresh/parser"
	"testing"
)

func TestParseInput(t *testing.T) {
	got, err := parser.ParseInput([]string{"--code=USD", "--date=2022-10-08"})
	want := struct {
		Code string
		Date string
	}{
		"USD",
		"2022-10-08",
	}

	if got != want || err != nil {
		fmt.Println(err)
		t.Errorf("got %q, wanted %q", got, want)
	}

	got, err = parser.ParseInput([]string{"-code=USD", "--date=2022-10-08"})
	want = struct {
		Code string
		Date string
	}{}

	if got != want {
		fmt.Println(err)
		t.Errorf("got %q, wanted %q", got, want)
	}
}