package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// e := errors.New("More coffee needed")
		// or
		e := fmt.Errorf("Sqrt of postivie values are only possible - input value was %v", f)
		return 0, sqrtError{"50.2289 N", "99.5457 W", e}
	}
	return 42, nil
}
