// Example-3 custom type 

package main 

import (
	"fmt"
	"log"
)

// implementing the errorString interface since it has the method Error() string
type norgateMathError struct {
	lat  string
	long string
	err  error
}

func (n norgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long,  n.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
			log.Println(err)
		}

}

func sqrt(f float64) (float64, error) {
		if f < 0 {
			nme := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
			return 0, norgateMathError{`50.2289 N`, `99.4656 W`, nme}
		}
		return 42, nil
}



// Example-2 Using fmr.Errorf

// package main

// import (
// 	"fmt"
// 	"log"
// )

// func main() {

// 	_, err := sqrt(-10.23)
// 	if err != nil {
// 			log.Fatalln(err)
// 		}
// }

// func sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		ErrNorgateMath := fmt.Errorf("norgate math again: square root of negative number: %v", f)
// 		return 0, ErrNorgateMath
// 	}
// 	return 42, nil
// }





// Example 1 - Understanding errors new()

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"log"
// )

// var ErrNorgateMath = errors.New("norgate math: square root of negative number")

// func main() {
// 	fmt.Printf("%T\n", ErrNorgateMath)
// 	_, err := sqrt(-10)

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// }

// func sqrt(f float64) (float64, error) {
// 	if f < 0 {
// 		return 0, ErrNorgateMath
// 	}
// 	return 42, nil
// }

// Output - 1
// *errors.errorString
// 2022/07/03 13:40:40 norgate math: square root of negativ
// e number
// exit status 1
