package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no good"]},{"First":"Miss","Last":"Parul","Age":20,"Sayings":["Hey all","How you doing??"]},{"First":"Narul","Last":"Sahoo","Age":64,"Sayings":["Good mornin","wassup??"]}]`
	fmt.Println(s)

	var people []person

	err := json.Unmarshal([]byte(s), &people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(people)

	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}

}
