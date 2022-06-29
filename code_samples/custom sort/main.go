package main

import (
	"fmt"
	"sort"
)

type Person struct {
	First string
	Age   int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }


type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].First < a[j].First }

func main() {
	p1 := Person{"James", 32}
	p2 := Person{"Parul", 27}
	p3 := Person{"Q", 64}
	p4 := Person{"M", 65}

	people := []Person{p1, p2, p3, p4}
	fmt.Println(people)
	fmt.Println("now sorted by Age")
	sort.Sort(ByAge(people))
	fmt.Println(people)

	fmt.Println("now sorted by Name")
	sort.Sort(ByName(people))
	fmt.Println(people)
}
