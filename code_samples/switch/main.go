package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this shouldn't print")
	case (2 == 4):
		fmt.Println("this shouldn't print 2")
	case (3 == 3):
		fmt.Println("this should print")
	case (4 == 4):
		fmt.Println("also true, does it print")
	}
	switch {
	case (9 == 9):
		fmt.Println("this should definitely print")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, it print as well coz we have fall through")
		fallthrough
	case (9 == 3):
		fmt.Println("this shouldn't print generally but here it will")
		fallthrough
	case (11 == 4):
		fmt.Println("also false, does it print: yes")
		fallthrough
	case (15 == 15):
		fmt.Println("yeh to print hoga hi")
		// if there would have been a fallthrough here to default case
		//would have been executed
	default:
		fmt.Println("this is default")
	}

	// switch case for specific values of a var
	switch "Bond" {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond James")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("this is default")
	}
	// or this can be done as
	n := "Moneypenny"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money for n")
	case "Bond":
		fmt.Println("bond James for n")
	case "Q":
		fmt.Println("this is Q for n")
	default:
		fmt.Println("this is default for n")
	}
	// adding multiple values
	m := "Moneywise"
	switch m {
	case "Moneypenny", "Moneywise", "Bond":
		fmt.Println("miss money or Moneywise or Bond for m")
	}
	// missing switch exp is like a bopol value true
	switch {
		case false:
			fmt.Println("this should not print for missing val")
		
		case true:
		fmt.Println("this should print for missing val")
    }

}
