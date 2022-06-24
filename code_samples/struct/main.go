package main

import "fmt"

type person struct {
	first string
	last  string
	age int
}

type secretAgent struct {
	person
	// just to check name collision
	first string
	ltk   bool
}

func main() {
	p1 := person{
		first: "Parul",
		last:  "Sahoo",
		age: 32,
	}
	// p2 := person{
	// 	first: "Narul",
	// 	last:  "Prusty",
	// }

	sa1 := secretAgent{
		person: person{
			first: "Narul",
			last:  "Prusty",
			age:   29,
		},
		first: "something collision",
		ltk:   true,
	}

	// declaring an anonymous struct
	p3 := struct {
		first string
		last string
		age int
	}{
		first: "Nidhi",
		last: "Bisht",
		age: 13,
	}

	fmt.Println(p1)
	fmt.Println(p3)
	fmt.Println(sa1)
	fmt.Println(sa1.person.first, sa1.first, sa1.ltk)
	fmt.Printf("The first name of first person is %v and last name is %v \n", p1.first, p1.last)
	fmt.Printf("The first name of second person is %v and last name is %v \n", sa1.first, sa1.last)


}
