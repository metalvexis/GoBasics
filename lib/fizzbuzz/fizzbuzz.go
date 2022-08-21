package basics

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
