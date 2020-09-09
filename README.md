# Comparison of Unit-Testing between  [Python](https://www.python.org/) and [Go](https://golang.org/)
 [![license: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) ![GitHub go.mod Go version (branch & subfolder of monorepo)](https://img.shields.io/github/go-mod/go-version/c0dehard/testing-go-vs-python/master?filename=golang%2Fgo.mod)

With built-in libraries only.


<small>Credits: Python part is based on [CoreyMSchafer's Snippet](https://github.com/CoreyMSchafer/code_snippets/tree/master/Python-Unit-Testing)</small>

## Why

Tests are usually very meaningful once you have taken the time and trouble to write them. For example if a method has changed it's logic (whether on purpose or not) in the future, unit-tests will save a lot of headaches and time. 
Simply run the tests to see if the code still works as asserted in the test and everything is OK!

> Note: If you are more experienced i'd recommend you to take a look at [Test-driven Development](https://en.wikipedia.org/wiki/Test-driven_development) It can be very strong if it is done properly, i.e. writing the test first and then trying to write suitable code.

### But for now let's give it a try by writing a simple calculator module first.

[In Python](/python/calc.py)  

```python
def add (x,y):
	return x + y

def substract(x,y):
	return x - y

def multiply(x,y):
	return x * y

def divide(x,y):
	if y == 0:
		raise ValueError("Can't divide by zero!")
	return x / y
```

[In Golang](/golang/calc.go) 

```go
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

```



## Required naming convention in terms of testing

| Python                                                       | Golang                                                       |
| ------------------------------------------------------------ | ------------------------------------------------------------ |
| test_WhatYouAreTesting.py                                    | WhatYouAreTesting_test.go                                    |
| `class TestClalc(unittest.TestCase):` `                                                      test_NameOfTest(self):` | `func TestNameOfTest(t *testing.T)`                                                         `func Test_name(t *testing.T)` |

### [Python test](/python/test_calc.py)

> Don't forget it's necessary to `import unittest` and the `module you'd like to test.`

```python
import unittest
import calc

# First create class that inherits from (unittest.TestCase).
class TestClalc(unittest.TestCase):
  
	# Make sure to use the naming convention, it wont run methods otherwise.
	def test_add(self):
		#  Through the inheritance of unittest.TestCase it's possible,
		#  to make all kinds of assertions that can be proven later. 
    
    # In this case i assert that (addition(x+y),sahll_return_value).
		self.assertEqual(calc.add(10,5),15)
		self.assertEqual(calc.add(-1,1),0)
		self.assertEqual(calc.add(-1,-1),-2)
	
  # Add all methods you'd like to test.
	def test_substract(self):
		self.assertEqual(calc.substract(10,5),5)
		self.assertEqual(calc.substract(-1,1),-2)
		self.assertEqual(calc.substract(-1,-1),0)
	
	def test_multiply(self):
		self.assertEqual(calc.multiply(10,5),50)
		self.assertEqual(calc.multiply(-1,1),-1)
		self.assertEqual(calc.multiply(-1,-1),1)
	
	def test_divide(self):
		self.assertEqual(calc.divide(10,5),2)
		self.assertEqual(calc.divide(-1,1),-1)
		self.assertEqual(calc.divide(-1,-1),1)

		# To catch if the error was thrown works a bit different
		# assertRaises(ErrorType,module.Method,params as arguments).
		# 
		# The more pleasant solution is using `with` to access the context manager,
		# in this case its possible to call the method like regularly.
		with self.assertRaises(ValueError):
			calc.divide(10,0)
    
"""
Usage via terminal
$ python3 -m unittest test_calc.py
""" 
if __name__ == "__main__":
	# Using unittest's main method allows to run within the editor.
	unittest.main()
```

### [Go test](/golang/calc_test.go)

> Don't forget it's necessary to `import testing` and the module you'd like to test (if external). As mentioned before i'm not using any external libraries, just teh built-in stuff. I know there's [testify](https://github.com/stretchr/testify) and its rly cool too!

But with that said, let's write an equivalent test in Go.

Example to run a general test via termianl: `$ go test`

```go
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

```






