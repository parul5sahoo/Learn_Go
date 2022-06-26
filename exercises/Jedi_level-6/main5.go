package main

import (
	"fmt"; 
	"math"
) 

type square struct {
	side float64
}

type circle struct {
	rad float64
}

type shape interface {
	areaCalc() float64
}

func (c circle) areaCalc() float64 {
	return 2 * math.Pi * c.rad
}

func (s square) areaCalc() float64 {
	return s.side * s.side
}

func info(sh shape) {
	switch sh.(type) {
	case circle:
		area := sh.areaCalc()
		fmt.Printf("The area of the circle is %v", area) 
	case square:
		sh.areaCalc()
		area := sh.areaCalc()
		fmt.Printf("The area of the square is %v", area)
	}

}

func main() {
	c1 := circle{
		rad: 14,
	}

	s1 := square{
		side: 22,
	}

	info(c1)
	fmt.Println(" and ")
	info(s1)

}
