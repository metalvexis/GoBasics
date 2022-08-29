package basics

import (
	"testing"
)

type roundbyTest struct {
	arg1     int
	arg2     int
	expected int
}

var roundbyTestArgs = []roundbyTest{
	{15, 5, 15},
	{13, 5, 15},
	{23, 5, 25},
	{11, 4, 12},
	{12, 4, 12},
	{13, 4, 12},
	{11, 3, 12},
	{12, 3, 12},
	{13, 3, 12},
}

func TestRoundby1(t *testing.T) {
	for _, test := range roundbyTestArgs {
		if result, _ := RoundBy(test.arg1, test.arg2); result != test.expected {
			t.Errorf("%d RoundBy %d Output %d not equal to expected %d", test.arg1, test.arg2, result, test.expected)
		}
	}
}

func TestRoundby2(t *testing.T) {
	var _, err = RoundBy(13, 0)
	if err == nil {
		t.Errorf("nearest number divisible by zero should be an error")
	}
}
