// Solution34_1 ...
func Solution34_1() (sum int) {
	for i := 3; i < 1000000; i++ {
		digits := DigitsOfNumber(int64(i))
		factorialSum := 0

		for _, v := range digits {
			factorialSum += Factorial(v)
		}

		if factorialSum == i {
			sum += factorialSum
		}
	}

	return
}
