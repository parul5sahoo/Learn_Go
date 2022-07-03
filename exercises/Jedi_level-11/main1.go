package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {

	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Anything for you!", "Never say never"},
	}
	bs, err := json.Marshal(p1)
	if err != nil {
		// using log Fatal to end the program
		// coz this is the main aim of it
		log.Fatalln("JSON marshalling failed", err)
	}
	fmt.Println(string(bs))
}
