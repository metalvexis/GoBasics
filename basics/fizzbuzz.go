package basics

/*
	Print integers 1 to N,
	but print
	“Fizz” if an integer is divisible by 3,
	“Buzz” if an integer is divisible by 5,
	and “FizzBuzz” if an integer is divisible by both 3 and 5.
*/

func FizzBuzz(input int) string {
	var output string
	if input%3 == 0 && input%5 == 0 {
		output = "FizzBuzz"

	} else if input%3 == 0 {
		output = "Fizz"
	} else if input%5 == 0 {
		output = "Buzz"
	} else {
		output = ""
	}

	return output
}
