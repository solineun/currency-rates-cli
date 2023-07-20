package main

import (
	"fmt"

	// "github.com/golang-module/carbon"
	"log"
	"qiwifresh/parser"
)



func main() {
	input, err := parser.ScanInput()
	if err != nil {
		log.Fatal(err)
	}
	options, err := parser.ParseInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(options)
}