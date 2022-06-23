package main

import "fmt"

func main() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)
	fmt.Println(m["James"])
	// prints zero value of value type for a key that is not present.
	fmt.Println(m["Bananas"])
	// , ok idiom to check if key is present in map or not
	v, ok := m["Bananas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["James"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
	// adding an elemnt to maps
	m["todd"] = 33

	for k, v := range m {
		fmt.Println(k, v)
	}
	xi := []int{4, 5, 6, 7, 8, 9, 10}
	for i, v := range xi {
		fmt.Println(i, v)
	}
	fmt.Println(m)
	delete(m, "James")
	fmt.Println(m)
	// deleting a key that doesn't exist does not give an error
	delete(m, "Parul Sahoo")
	fmt.Println(m)

	fmt.Println(m["Miss Moneypenny"])
	fmt.Println(m["parul Sahoo"])

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}
	fmt.Println(m)
	// returns an empty stmt

}
