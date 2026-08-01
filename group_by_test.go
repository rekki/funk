package funk

import (
	"reflect"
	"testing"
)

type numberClass struct {
	number int
	label  string
}

func TestGroupBy(t *testing.T) {
	keyCalls := 0
	numbers, naNumbers, primeNumbers, semiPrimeNumbers := []numberClass{
		{0, "N/A"},
		{1, "N/A"},
		{2, "prime"},
		{3, "prime"},
		{4, "semi-prime"},
		{5, "prime"},
		{6, "semi-prime"},
	},
		[]numberClass{
			{0, "N/A"},
			{1, "N/A"},
		},
		[]numberClass{
			{2, "prime"},
			{3, "prime"},
			{5, "prime"},
		},
		[]numberClass{
			{4, "semi-prime"},
			{6, "semi-prime"},
		}

	numberMap := GroupBy(numbers, func(i numberClass) string {
		keyCalls++
		return i.label
	})
	if keyCalls != len(numbers) {
		t.Fatalf("GroupBy() called key %d times, want %d", keyCalls, len(numbers))
	}

	testNumbers := func(expectedNumbers []numberClass, numberType string) {
		if !reflect.DeepEqual(expectedNumbers, numberMap[numberType]) {
			t.Fatalf("GroupBy()[%q] = %v, want %v", numberType, numberMap[numberType], expectedNumbers)
		}
	}

	testNumbers(naNumbers, "N/A")
	testNumbers(primeNumbers, "prime")
	testNumbers(semiPrimeNumbers, "semi-prime")
}
