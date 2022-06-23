package main

import "fmt"

func main() {
	favSport := "Football"
	switch favSport {
	case "Skiing":
		fmt.Println("Go to the mountains")
	case "Swimming":
		fmt.Println("Go to the pool")
	case "Surfing":
		fmt.Println("Go to the beach")
	default:
		fmt.Println("So it's none of those")
	}

}
