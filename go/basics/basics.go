package main

import "fmt"

type gasEngine struct{
	mgp uint8
	gallons uint8
	owner
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

type owner struct{
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mgp
}

func (e electricEngine) milesLeft() uint8 {
	return e.mpkwh * e.kwh
}

type engine interface {
	milesLeft() uint8
}



func main() {
	engine1 := gasEngine{
		1,1, owner{"sam"},
	}

	engine2 := electricEngine{
		1,2,
	}

	fmt.Println(engine1.milesLeft())

	fmt.Println(engine2.milesLeft())
}