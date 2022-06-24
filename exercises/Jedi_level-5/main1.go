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
		lastName:    "Sahoo",
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

}
