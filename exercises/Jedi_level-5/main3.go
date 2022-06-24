package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "White",
		},
		fourWheels: false,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "Blue",
		},
		luxury: true,
	}
	fmt.Println(truck1)
	fmt.Println(sedan1)

	fmt.Printf("Sedan 1 has %v doors, is of color %v and is %vly luxurious \n", sedan1.vehicle.doors, sedan1.vehicle.color, sedan1.luxury)
	fmt.Printf("Truck 1 has %v doors, is of color %v and has %v wheels \n", truck1.vehicle.doors, truck1.vehicle.color, truck1.fourWheels)

}
