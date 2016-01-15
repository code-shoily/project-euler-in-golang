package problems

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"unicode"
)

// Fibonacci Returns the nth Fibonacci number
func Fibonacci(limit int) int {
	first, second := 1, 1

	for idx := 3; idx <= limit; idx++ {
		temp := first + second
		first = second
		second = temp
	}
	return second
}

// GCD Returns the greatest common divisor for a and b
func GCD(a, b int) int {
	rem := a % b
	for rem != 0 {
		a, b = b, rem
		rem = a % b
	}
	return b
}

// LCM2 computes the least common multiplier of 2 numbers
func LCM2(a, b int) int {
	return int(a * b / GCD(a, b))
}

// LCM computes LCM of many numbers
func LCM(nums ...int) int {
	if len(nums) == 2 {
		return LCM2(nums[0], nums[1])
	}
	return LCM(nums[0], LCM(nums[1:]...))
}

// NaivePrime Outputs if the number is a prime or not
func NaivePrime(number int) bool {
	if number <= 1 || (number%2 == 0 && number != 2) {
		return false
	}

	for i := 3; float64(i) <= math.Sqrt(float64(number)); i += 2 {
		if number%i == 0 {
			return false
		}
	}

	return true
}

// ReverseDigits reverses the digits of a number
func ReverseDigits(n int) (output int) {
	for n > 0 {
		output = output*10 + n%10
		n /= 10
	}
	return
}

// IsPalindromicNumber tests if number is palindromic
func IsPalindromicNumber(n int) bool {
	return n == ReverseDigits(n)
}

// Multiply multiplies all numbers given
func Multiply(numbers ...int) (result int) {
	result = 1
	for _, v := range numbers {
		result *= v
	}
	return
}

// Sum totals all numbers given
func Sum(numbers ...int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return
}

// DigitsOfNumber returns the digits of a number in order
func DigitsOfNumber(n int64) []int {
	limit, ctr := int(math.Log10(float64(n))), 0
	digits := make([]int, limit+1)
	for ctr <= limit {
		digits[limit-ctr] = int(n % 10)
		n /= 10
		ctr++
	}

	return digits
}

// DigitsOfNumberReversed returns the digits of a number in reverse order
func DigitsOfNumberReversed(n int64) []int {
	limit, ctr := int(math.Log10(float64(n))), 0
	digits := make([]int, limit+1)
	for ctr <= limit {
		digits[ctr] = int(n % 10)
		n /= 10
		ctr++
	}

	return digits
}

// Contains tells whether or not a number is contained within a slice
func Contains(list []int, n int) bool {
	for _, v := range list {
		if v == n {
			return true
		}
	}

	return false
}

// Factorial calculates the factorial of any number
func Factorial(number int) int {
	fact := 1
	for i := 1; i <= number; i++ {
		fact *= i
	}
	return fact
}

// NumberToWordToThousand returns the (in-word) of any number
func NumberToWordToThousand(number int) string {
	if number == 1000 {
		return "one thousand"
	}
	var presets = []string{
		"zero", // NOT IMPORTANT HERE
		"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen",
		"seventeen", "eighteen", "nineteen"}
	var tens = []string{
		"twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety"}

	if number < 20 {
		return presets[number]
	} else if number < 100 {
		idx10 := int(number / 10)
		rem10 := int(number % 10)
		if rem10 == 0 {
			return fmt.Sprintf("%s", tens[idx10-2])
		}
		return fmt.Sprintf("%s-%s", tens[idx10-2], presets[rem10])
	}

	idx100 := int(number / 100)
	rem100 := number % 100
	if rem100 == 0 {
		return fmt.Sprintf("%s hundred", presets[idx100])
	}
	return fmt.Sprintf("%s hundred and %s", presets[idx100], NumberToWordToThousand(rem100))
}

// CountLetters counts the letters in a sentence
func CountLetters(sentence string) (sum int) {
	for _, v := range sentence {
		if unicode.IsLetter(v) {
			sum++
		}
	}
	return
}

// GetCharacterNumberIgnoreCase Character Number returns the position of a character
func GetCharacterNumberIgnoreCase(c rune) int {
	c = unicode.ToLower(c)
	return 1 + int(c)%97
}

// IsTriangleNumber checks if a number is triangle number
func IsTriangleNumber(n int) bool {
	eightNPlusOne := int64(8*n + 1)
	if eightNPlusOne%2 == 0 {
		return false
	}

	sqrtOfEightNPlusOne := math.Sqrt(float64(eightNPlusOne))
	return sqrtOfEightNPlusOne == math.Floor(sqrtOfEightNPlusOne)
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines. FIXME EXTREMELY BAD CODE GO HOME!
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
