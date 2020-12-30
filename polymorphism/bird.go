package main

import (
	"fmt"
)

type flight interface {
	fly()
}
type bird struct {
	name    string
	species string
}

func (b bird) fly() {
	fmt.Println("Some birds can fly")
}

type duck struct {
	bird
	canSwim bool
}

func (d duck) fly() {
	fmt.Println("Yes yes, duck can fly, though short distance")
}

func doFly(f flight) {
	f.fly()
}

func main() {

	b := bird{
		name:    "crow",
		species: "blaky",
	}
	d := duck{
		bird: bird{
			name:    "duky",
			species: "cutie",
		},
		canSwim: true,
	}

	doFly(b)
	doFly(d)
}
