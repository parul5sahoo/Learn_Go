// opening the file we just created in the prev example.
package main

import (
	 "fmt"
	 "os"
	 "io/ioutil"
)

func main() {
	f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// close the file to stop consumption of unneccesary resources
	defer f.Close()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))
}



// Examples 3
// package main

// import (
// 	"fmt"
// 	"io"
// 	"strings"
// 	"os"
// )

// func main() {
// 		f, err := os.Create("names.txt")

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		defer f.Close()

// 		r := strings.NewReader("Wassup")
// 		// creates a new file names.txt and add this string into it
// 		io.Copy(f, r)

// }







// Example 2


// package main

// import "fmt"

// func main() {
// 	var ans1, ans2, ans3 string

// 	fmt.Print("Name: ")
// 	_, err := fmt.Scan(&ans1)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Print("Fav Food: ")
// 	_, err = fmt.Scan(&ans2)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Print("Fav sport: ")
// 	_, err = fmt.Scan(&ans3)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(ans1, ans2, ans3)

// }

// Example 1

// package main

// import "fmt"

// func main() {
// 	n , err := fmt.Println("Hello")

// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(n) 
// 	// the o/p is 6 because hello has 5 letters and the new line char takes the 6th position.
// }
