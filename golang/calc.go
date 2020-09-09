package main

import "errors"

type calculator struct{
}

func (c calculator)add(x int,y int) int{
	return  x + y 
}

func (c calculator)substract(x int, y int) int {
	return  x - y 
}

func (c calculator)multiply(x int, y int) int {
	return x * y
}

func (c calculator)divide(x int, y int) int {
	if y == 0{
		errors.New("Can't divide by zero!")
	}
	return x / y
}
