package main

import (
		// "fmt"
		"os"
		// "log"
)

func main() {
	// f, err := os.Create("log.txt")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer f.Close()
	// log.SetOutput(f)


	// f2, err := os.Open("no-file.txt")

	//this code piece was for checking save err to a new file

	_, err := os.Open("no-file.txt")
	if err != nil {
		 // fmt.Println("err happened", err) 
		 // helps with customising the err msg

		 // log.Println("err happened", err)
		 // give the date & time of execution as well in the O/p
		 // also can be used to save err msg in a separate file


		 /* 
		 	... the fatal func call os.Exit(1) after writing the log msg
		 */
		// log.Fatalln(err)
		// Fatalln is basically println() followed by a caall to os.Exit(1)
		
		// log.Panicln(err)

		panic(err)

	}
	// defer f2.Close()
	// fmt.Println("check the log.txt file in the directory")
	// to save the err msgs in a file for reviewing later.
}