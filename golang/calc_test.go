package main

import (
	"testing" 
)

var c = &calculator{}

func Test_add(t *testing.T){
	additions := []int{
		c.add(10,5),
		c.add(-1,1),
		c.add(-1,-1),
	}
	expected := []int{15,0,-2}
	
	for i,sum := range additions{
		if sum != expected[i]{
			t.Errorf("Addition was incorrect, got %d, want: %d.",
			sum,expected[i])
		}
	}
}

func Test_subtract(t *testing.T){
	subtractions := []int{
		c.substract(10,5),
		c.substract(-1,1),
		c.substract(-1,-1),
	}
	expected := []int{5,-2,0}
	
	for i,difference := range subtractions{
		if difference != expected[i]{
			t.Errorf("Substraction was incorrect, got %d, want: %d.",
			difference,expected[i])
		}
	}
}

func Test_multiply(t *testing.T){
	multipications := []int{
		c.multiply(10,5),
		c.multiply(-1,1),
		c.multiply(-1,-1),
	}
	expected := []int{50,-1,1}
	
	for i,product := range multipications{
		if product != expected[i]{
			t.Errorf("Product was incorrect, got %d, want: %d.",
			product,expected[i])
		}
	}
}

func Test_divide(t *testing.T){
	divisions := []int{
		c.divide(10,5),
		c.divide(-1,1),
		c.divide(-1,-1),
	}
	expected := []int{2,-1,1}
	
	for i,quotient := range divisions{
		if quotient != expected[i]{
			t.Errorf("Division was incorrect, got %d, want: %d.",
			quotient,expected[i])
		}
	}
}
