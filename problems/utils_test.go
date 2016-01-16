package problems

import (
	"reflect"
	"testing"
)

func TestFactorial(t *testing.T) {
	type testpair struct {
		arg    int
		output int
	}

	var testdata = []testpair{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
	}

	for _, pair := range testdata {
		v := Factorial(pair.arg)
		if v != pair.output {
			t.Error("For", pair.arg, "Expected", pair.output, "Got", v)
		}
	}
}

func TestContainer(t *testing.T) {
	type teststruct struct {
		haystack []int
		needle   int
		output   bool
	}

	var testdata = []teststruct{
		{[]int{1, 2, 3}, 2, true},
		{[]int{0, 2, 3}, -1, false},
	}

	for _, tuple := range testdata {
		v := Contains(tuple.haystack, tuple.needle)

		if v != tuple.output {
			t.Error("For", tuple.haystack, "and", tuple.needle, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestDigitsOfNumber(t *testing.T) {
	type teststruct struct {
		number int64
		output []int
	}

	var testdata = []teststruct{
		{1234567, []int{1, 2, 3, 4, 5, 6, 7}},
		{1, []int{1}},
		{0, []int{0}},
	}

	for _, tuple := range testdata {
		v := DigitsOfNumber(tuple.number)

		if !reflect.DeepEqual(v, tuple.output) {
			t.Error("For", tuple.number, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestDigitsOfNumberReversed(t *testing.T) {
	type teststruct struct {
		number int64
		output []int
	}

	var testdata = []teststruct{
		{1234567, []int{7, 6, 5, 4, 3, 2, 1}},
	}

	for _, tuple := range testdata {
		v := DigitsOfNumberReversed(tuple.number)

		if !reflect.DeepEqual(v, tuple.output) {
			t.Error("For", tuple.number, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestReverseDigits(t *testing.T) {
	type teststruct struct {
		number int
		output int
	}

	var testdata = []teststruct{
		{1234567, 7654321},
		{1000, 1},
		{101, 101},
	}

	for _, tuple := range testdata {
		v := ReverseDigits(tuple.number)

		if v != tuple.output {
			t.Error("For", tuple.number, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestPalindromicNumber(t *testing.T) {
	type teststruct struct {
		number int
		output bool
	}

	var testdata = []teststruct{
		{101, true},
		{203313, false},
		{77, true},
		{233, false},
	}

	for _, tuple := range testdata {
		v := IsPalindromicNumber(tuple.number)

		if v != tuple.output {
			t.Error("For", tuple.number, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestSum(t *testing.T) {
	type teststruct struct {
		numbers []int
		output  int
	}

	var testdata = []teststruct{
		{[]int{1, 2, 3, 4, 5}, 15},
	}

	for _, tuple := range testdata {
		v := Sum(tuple.numbers...)

		if v != tuple.output {
			t.Error("For", tuple.numbers, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestMultiply(t *testing.T) {
	type teststruct struct {
		numbers []int
		output  int
	}

	var testdata = []teststruct{
		{[]int{1, 2, 3, 4, 5}, 120},
	}

	for _, tuple := range testdata {
		v := Multiply(tuple.numbers...)

		if v != tuple.output {
			t.Error("For", tuple.numbers, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestNumberToWordToThousand(t *testing.T) {
	type teststruct struct {
		number int
		output string
	}

	var testdata = []teststruct{
		{1, "one"},
		{19, "nineteen"},
		{25, "twenty-five"},
		{42, "forty-two"},
		{120, "one hundred and twenty"},
		{342, "three hundred and forty-two"},
		{500, "five hundred"},
		{999, "nine hundred and ninety-nine"},
		{1000, "one thousand"},
	}

	for _, tuple := range testdata {
		v := NumberToWordToThousand(tuple.number)

		if v != tuple.output {
			t.Error("For", tuple.number, "Expected", tuple.output, "Got", v)
		}
	}

}

func TestCountLetters(t *testing.T) {
	type teststruct struct {
		sentence string
		output   int
	}

	var testdata = []teststruct{
		{"Project", 7},
		{"Project Euler", 12},
		{"Forty-Two", 8},
		{"One Hundred and Eleven", 19},
	}

	for _, tuple := range testdata {
		v := CountLetters(tuple.sentence)

		if v != tuple.output {
			t.Error("For", tuple.sentence, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestGetCharNumberIgnoreCase(t *testing.T) {
	type teststruct struct {
		c      rune
		output int
	}

	var testdata = []teststruct{
		{'a', 1},
		{'i', 9},
		{'m', 13},
		{'s', 19},
		{'z', 26},
	}

	for _, tuple := range testdata {
		v := GetCharacterNumberIgnoreCase(tuple.c)

		if v != tuple.output {
			t.Error("For", tuple.c, "Expected", tuple.output, "Got", v)
		}
	}
}

func TestIsTriangularNumber(t *testing.T) {
	type teststruct struct {
		c      int
		output bool
	}

	var testdata = []teststruct{
		{1, true},
		{3, true},
		{6, true},
		{11, false},
		{22, false},
		{55, true},
	}

	for _, tuple := range testdata {
		v := IsTriangleNumber(tuple.c)

		if v != tuple.output {
			t.Error("For", tuple.c, "Expected", tuple.output, "Got", v)
		}
	}
}
