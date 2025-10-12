package main

func topKFrequent(nums []int, k int) []int {
	frequencyTracker := make(map[int]int)
	buckets := make([][]int, len(nums)+1)

	for _, v := range nums {
		frequencyTracker[v]++
	}

	for val, freq := range frequencyTracker {
		buckets[freq] = append(buckets[freq], val)
	}

	var mostFrequent []int

	i := len(buckets) - 1
	for len(mostFrequent) <= k && i > 0 {
		for j := 0; j < len(buckets[i]); j++ {
			if len(mostFrequent) == k {
				break
			}
			mostFrequent = append(mostFrequent, buckets[i][j])
		}
		i--

	}
	return mostFrequent

}

