package dynamic

import (
	"fmt"
)

// Given a list of denominations (e.g., [1, 2, 5] means you have coins worth
// 1 penny, 2 pennies, and 5 pennies) and a target number M, find all possible
// combinations, if any, of coins in the given denominations that add up to M.
// You can use coins of the same denomination more than once.
// Time complexity: O(N), where N is the number of distinct type of coins. I did not include the sort in the complexity calculation

func pettyCash(denominations []int, m int) (map[int]int, error) {
	den := insertionSort(denominations)
	out := make(map[int]int)

	for i := len(den) - 1; i >= 0; i-- {
		coin := den[i]
		nb := int(m / coin)
		out[coin] = nb
		m -= nb * coin
	}

	if m > 0 {
		return nil, fmt.Errorf("couldn't find a way to deal with the remaining value of %d", m)
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
