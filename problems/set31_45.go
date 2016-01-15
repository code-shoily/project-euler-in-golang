package problems

import "strings"

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

//Solution42_1 ...
func Solution42_1() (count int) {
	words, err := ReadLines("./p042_words.txt")
	r := strings.NewReplacer("\"", "")
	words = strings.Split(r.Replace(words[0]), ",")

	if err != nil {
		panic(err)
	}
	for _, word := range words {
		tmp := 0
		for _, r := range word {
			tmp += GetCharacterNumberIgnoreCase(r)
		}

		if IsTriangleNumber(tmp) {
			count++
		}
	}
	return
}
