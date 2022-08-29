package basics

import (
	"errors"
)

/*
	Find the nearest whole number

	num: 13
	nearest: 4

	possible nearest values: 12, 16
*/

func RoundBy(num int, roundNearest int) (int, error) {
	if roundNearest <= 0 {
		return 0, errors.New("roundNearest > 0")
	}
	var mod int = num % roundNearest

	if mod == 0 {
		return num, nil
	}

	var nearestLower int = num - mod
	var lowerGap int = absInteger(num - nearestLower)
	var nearestHigher int = nearestLower + roundNearest
	var higherGap int = absInteger(num - nearestHigher)

	// fmt.Printf("%d : %d <-> %d : %d\n", nearestLower, lowerGap, nearestHigher, higherGap)

	// round up if equally near
	if lowerGap == higherGap {
		return nearestLower, nil
	}

	if lowerGap < higherGap {
		return nearestLower, nil
	}

	return nearestHigher, nil
}

func absInteger(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
