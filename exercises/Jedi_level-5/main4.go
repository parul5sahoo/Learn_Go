package main

import "fmt"

func main() {

	ano1 := struct {
		Name    string
		Age     int
		Details map[string]string
	}{
		Name: "Parul Sahoo",
		Age:  20,
		Details: map[string]string{
			"Gender":   "Female",
			"Caste":    "General",
			"Religion": "Hindu",
		},
	}

	fmt.Println(ano1)

	fmt.Printf("ano1 is %v of age %v and their details are as follows:- \n", ano1.Name, ano1.Age)

	for k, v := range ano1.Details {
		fmt.Printf("\t%v:- %v \n", k, v)
	}

}


