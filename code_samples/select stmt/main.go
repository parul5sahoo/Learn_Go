package main
import "fmt"

func main() {
	eve := make(chan int)
	odd := make(chan int)
	quit := make(chan int)

	// send
	go send(eve, odd, quit)

	//receive
	receive(eve, odd, quit)
	fmt.Println("about to exit")
	
}

func receive(e, o, q <-chan int) {
	for{
		select {
			case v := <- e:
				fmt.Println("from the even channel:", v)
			case v := <- o:
				fmt.Println("from the odd channel:", v)
			case v := <- q:
				fmt.Println("from the quit channel:", v)
				return
		}
		
	}	
	
}

func send(e, o, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0{
			e<- i
		}else{
			o<- i
		}
	}
	// close(e)
	// close(o) to get rid of zeros from e and o
	q<-0
}

// using comma ok with select stmt with bool...

// package main

// import "fmt"

// func main() {
// 	eve := make(chan int)
// 	odd := make(chan int)
// 	quit := make(chan bool)

// 	// send
// 	go send(eve, odd, quit)

// 	//receive
// 	receive(eve, odd, quit)
// 	fmt.Println("about to exit")

// }

// func receive(e, o <-chan int, q <-chan bool) {
// 	for {
// 		select {
// 		case v := <-e:
// 			fmt.Println("from the even channel:", v)
// 		case v := <-o:
// 			fmt.Println("from the odd channel:", v)
// 		case i, ok := <-q:
// 			if !ok {
// 				fmt.Println("from comma ok:", i)
// 				return
// 			} else {
// 				fmt.Println("from comma ok:", i)
// 			}

// 		}

// 	}

// }

// func send(e, o chan<- int, q chan<- bool) {
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			e <- i
// 		} else {
// 			o <- i
// 		}
// 	}

// 	close(q)
// }


// using comma ok with select stmt and int 

// package main

// import "fmt"

// func main() {
// 	even := make(chan int)
// 	odd := make(chan int)
// 	quit := make(chan int)

// 	// send
// 	go send(even, odd, quit)

// 	//receive
// 	receive(even, odd, quit)
// 	fmt.Println("about to exit")

// }

// func receive(e, o, q <-chan int) {
// 	for {
// 		select {
// 		case v := <-e:
// 			fmt.Println("from the even channel:", v)
// 		case v := <-o:
// 			fmt.Println("from the odd channel:", v)
// 		case i, ok := <-q:
// 			if !ok {
// 				fmt.Println("from comma ok:", i, ok)
// 				return
// 			} else {
// 				fmt.Println("from comma ok bit:", i, ok)
// 			}

// 		}

// 	}

// }

// func send(e, o, q chan<- int) {
// 	for i := 0; i < 100; i++ {
// 		if i%2 == 0 {
// 			e <- i
// 		} else {
// 			o <- i
// 		}
// 	}

// 	close(q)
// }
