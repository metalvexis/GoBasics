package main

import "fmt"

/*
Print integers 1 to N,
but print
“Fizz” if an integer is divisible by 3,
“Buzz” if an integer is divisible by 5,
and “FizzBuzz” if an integer is divisible by both 3 and 5.
*/

func main() {
	fmt.Println("Hello, 世界")

	input := 3

	if input%3 == 0 && input%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if input%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println("Buzz")
	}
}
