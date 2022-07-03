package main

import (
	"encoding/json"
	"fmt"
	"log"
	"errors"
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
	bs, err := toJSON(p1)
	if err != nil {
		log.Println(err)
		return // to finish the program

		// or use log.Fatalln to end the program
	}

	fmt.Println(string(bs))
}

// to JSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		// return []byte{}, fmt.Errorf("There was an error in to JSON: %v", err)
		// we cant use Println instead of Errof coz Pritnln return an int and error

		// However we can do
		return []byte{}, errors.New(fmt.Sprintf("There was an error in to JSON %v", err))
	}
	return bs, nil

}
