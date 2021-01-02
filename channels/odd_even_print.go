package main

import "fmt"

const max = 20

//Two go routines are printing odd & even numbers.
//Co-ordinate & control them to print the numbers sequentially
func main() {

	oddC := make(chan int)
	evenC := make(chan int)

	go emitOdd(oddC)
	go emitEven(evenC)

	for i := 1; i < max; i += 2 {
		fmt.Println(<-oddC)
		fmt.Println(<-evenC)
	}
}

func emitOdd(c chan<- int) {

	for i := 1; i < max; i += 2 {
		c <- i
	}
	close(c)

}

func emitEven(c chan<- int) {

	for i := 2; i <= max; i += 2 {
		c <- i
	}
	close(c)
}
