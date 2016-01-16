package problems

import (
	"math"
	"math/big"
)

// Solution16_1 ...
func Solution16_1() int {
	twoToThe1000 := new(big.Int).Exp(big.NewInt(2), big.NewInt(1000), nil).String()

	sum := 0
	for _, v := range twoToThe1000 {
		sum += int(v) - '0'
	}

	return sum
}

// Solution17_1 ...
func Solution17_1() (sum int) {
	for i := 1; i <= 1000; i++ {
		n := NumberToWordToThousand(i)
		numberOfLetters := CountLetters(n)
		sum += numberOfLetters
	}
	return
}

// Solution30_1 ...
func Solution30_1() (sum int64) {
	var i int64
	for i = 2; i < 1653372; i++ {
		digits := DigitsOfNumber(i)
		var sumToThePower int64
		for _, n := range digits {
			sumToThePower += int64(math.Pow(float64(n), 5))
		}

		if sumToThePower == i {
			sum += sumToThePower
		}

	}
	return
}
