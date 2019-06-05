package dynamic

import (
	"fmt"
)

// Given a list of denominations (e.g., [1, 2, 5] means you have coins worth
// $1, $2, and $5) and a target number M, find all possible combinations, if
// any, of coins in the given denominations that add up to M. You can use coins
// of the same denomination more than once.

// Time complexity: O(MN), where N is the number of distinct type of coins.
// Space complexity: O(M).

func pettyCash(denominations []int, m int) (map[int]int, error) {
	den := insertionSort(denominations)
	out := make(map[int]int)
	for m > 0 {
		lastval := m
		for i := len(den)-1; i >= 0; i-- {
			coin := den[i]
			if coin <= m {
				out[coin]++
				m -= coin
				break;
			}
		}
		if m == lastval {
			return nil, fmt.Errorf("couldn't find a way to deal with the remaining value of %d", m)
		}
	}

	return out, nil
}

func insertionSort(arr []int) []int {
	out := append([]int(nil), arr...)
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0 && out[j-1] > out[j]; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}
