package main

func twoSum(nums []int, target int) []int {
	pairs := make(map[int]int)

	for i, v := range nums { // you could do this in one pass but it doesn't hurt big O to do it in two
		pairs[target-v] = i
	}

	for i, v := range nums {
		if pairs[v] != 0 && i != pairs[v] {
			return []int{min(i, pairs[v]), max(i, pairs[v])}
		}
	}

	return nil
}

