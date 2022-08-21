package main

import (
	"fmt"
	"os"
	"strconv"

	echo "github.com/metalvexis/gobasics/lib/echo"
	basics "github.com/metalvexis/gobasics/lib/fizzbuzz"
)

/*
	Print integers 1 to N,
	but print
	“Fizz” if an integer is divisible by 3,
	“Buzz” if an integer is divisible by 5,
	and “FizzBuzz” if an integer is divisible by both 3 and 5.
*/

func main() {
	var input, _ = strconv.Atoi(os.Args[1])
	fmt.Println(basics.FizzBuzz(input))
	fmt.Println(echo.Echo(os.Args[1]))
}
