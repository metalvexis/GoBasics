package basics

import "testing"

type argMain struct {
	arg1     int
	expected string
}

var mainTest = []argMain{
	{3, "Fizz"},
	{5, "Buzz"},
	{9, "Fizz"},
	{15, "FizzBuzz"},
	{4, ""},
}

func TestFizzBuzz(t *testing.T) {
	t.Log("Testing FizzBuzz")
	for _, test := range mainTest {
		if result := FizzBuzz(test.arg1); result != test.expected {
			t.Errorf("Output %q not equal to expected %q", result, test.expected)
		}
	}
}
