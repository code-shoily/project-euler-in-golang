// Solution50_1 Consequetive prime sum
func Solution50_1() (prime int) {
	const max = 1000000
	primesBelowMillion, cumulativeSum := []int{}, []int{}
	sum := 0

	for i := 0; i < max; i++ {
		if NaivePrime(i) {
			sum += i
			primesBelowMillion = append(primesBelowMillion, i)
			cumulativeSum = append(cumulativeSum, sum)
		}
	}
	limit := 0

	for i := len(cumulativeSum) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			difference := cumulativeSum[i] - cumulativeSum[j]
			if difference >= max {
				break
			}

			if NaivePrime(difference) {
				if i-j > limit {
					limit = i - j
					prime = difference
				}
			}
		}
	}
	return
}
