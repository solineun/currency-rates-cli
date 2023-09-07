package main

import (
	"fmt"
	"log"
	v "currecyrates/view"
	c "currecyrates/controller"
)



func main() {
	input, err := v.ScanInput()
	if err != nil {
		log.Fatal(err)
	}
	options, err := c.ParseInput(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(options)
}