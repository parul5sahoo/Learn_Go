package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	favFlavours []string
}

func main() {
	p1 := person{
		firstName:   "Nidhi",
		lastName:    "Sahoo",
		favFlavours: []string{"Vanila", "Chocolate", "Orange"},
	}
	p2 := person{
		firstName:   "Anisha",
		lastName:    "Prusty",
		favFlavours: []string{"Mango", "Chocolate Chip", "Tutty Fruity"},
	}

	fmt.Printf("%v 's fav flavours are\n", p1.firstName)
	for i, p := range p1.favFlavours {

		fmt.Printf("\t%d %v", i+1, p)
	}
	fmt.Println("")
	fmt.Printf("%v 's fav flavours are\n", p2.firstName)

	for i, p := range p2.favFlavours {
		fmt.Printf("\t%d %v", i+1, p)
	}
	fmt.Println("")

	personMap := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}
	// second work around for a solution.
	//[]string{
	//"Sahoo":  []string{"Nidhi", "Vanila", "Chocolate", "Orange"},
	//"Prusty": []string{"Anisha", "Mango", "Chocolate chip", "Tutty Fruity"},
	//}
	fmt.Println(personMap)
	fmt.Println("")

	for k, v := range personMap {
		fmt.Println("for key:", k)
		fmt.Printf("%v %v's fav flavours are: %v \n", v.firstName, v.lastName, v.favFlavours)
	}
}
