package main

import "fmt"

func main() {

	m := map[string][]string{
		`Parul Sahoo`: []string{`Ice Cream`, `Gupchup`, `momos`},
		`Narul Sahoo`: []string{`Achar`, `Papad`, `Dhokla`},
		`Neera Sahoo`: []string{`Juice`, `fruits`, `cereals`},
	}
	// adding values to a map
	m[`Nidhi Sahoo`] = []string{`Daal`, `Rice`, `Roti`}
	// deleting values from maps
	delete(m, `Nidhi Sahoo`)
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}

}
