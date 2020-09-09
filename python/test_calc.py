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