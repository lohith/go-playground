package main

import "fmt"

//Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

func fibonacci() func() int {
	prevState := 0
	currState := 1
	return func() int {
		prevState, currState = currState, prevState+currState
		return currState
	}

}
