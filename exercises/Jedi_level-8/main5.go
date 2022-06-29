package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ByAge []user

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

type ByLast []user

func (a ByLast) Len() int           { return len(a) }
func (a ByLast) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLast) Less(i, j int) bool { return a[i].Last < a[j].Last }

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
			"Mornin, Hey Everyone",
			"Khair what was I saying?",
		},
	}

	u3 := user{
		First: "Narul",
		Last:  "Prusty",
		Age:   64,
		Sayings: []string{
			"zzzz, Hey fellas!!",
			"Wassup everybody?",
		},
	}

	users := []user{u1, u2, u3}

	for i, user := range users {
		fmt.Println("User #", i)
		fmt.Println("\t", user.First, user.Last, user.Age)
		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}
	}

	fmt.Println("Sorted by Age of user and sorting the sayings strings")

	sort.Sort(ByAge(users))
	for i, user := range users {
		fmt.Println("User #", i)
		fmt.Println("\t", user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}
	}

	fmt.Println("Sorted by Last name of user")
	sort.Sort(ByLast(users))
	for i, user := range users {
		fmt.Println("User #", i)
		fmt.Println("\t", user.First, user.Last, user.Age)
		sort.Strings(user.Sayings)
		for _, saying := range user.Sayings {
			fmt.Println("\t\t", saying)
		}
	}

}
