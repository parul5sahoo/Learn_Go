package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}

	u2 := user{
		First: "Parul",
		Last:  "Sahoo",
		Age:   20,
		Sayings: []string{
			"Hey Everyone",
			"Khair what was I saying?",
		},
	}

	u3 := user{
		First: "Narul",
		Last:  "Prusty",
		Age:   64,
		Sayings: []string{
			"Hey fellas!!",
			"Wassup everybody?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)
	// turning data into json again with NewEncoder
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:", err)
	}
}
