package dynamic

import (
	"testing"
)

func TestSort (t *testing.T) {
	arr := []int{5, 4, 3, 2, 1}

	out := insertionSort(arr)

	for i := 1; i < len(out); i++ {
		if out[i] < out[i-1] {
			t.Errorf("out[%d] = %d should be smaller than out[%d] = %d", i-1, out[i-1], i, out[i])
		}
	}
}

func TestPettyCash (t *testing.T) {
	coins := []int{1, 2, 5, 10, 20, 50, 100, 200}
	change := 1799

	actual, err := pettyCash(coins, change)

	if err != nil {
		t.Errorf("the function failed with error: %v", err)
		return
	}

	expected := map[int]int{
		200: 8,
		100: 1,
		50: 1,
		20: 2,
		10: 0,
		5: 1,
		2: 2,
		1: 0,
	}

	for coin, nb := range expected {
		if actual[coin] != nb {
			t.Errorf("Expected %d coins of value %d, got %d instead", nb, coin, actual[coin])
		}
	}
}

func TestFailOnImpossibleCases (t *testing.T) {
	coins := []int{5, 10, 20, 50, 100, 200}
	change := 1799

	if out, err := pettyCash(coins, change); err == nil {
		t.Errorf("expected the function to return with an error, instead return with solution: %v", out)
	}
}