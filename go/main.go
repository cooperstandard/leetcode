package main

import "fmt"

func main() {
	// input := []int{7, 7}
	// output := topKFrequent(input, 1)
	// fmt.Printf("%v: %v\n", slices.Equal(output, []int{7}), output)
	fmt.Println(productExceptSelf([]int{-1,0,1,2,3}))
}

func productExceptSelf(nums []int) []int {
    product := 1
    for _,v := range nums {
        product *= v
    }

    for i,_ := range nums {
        nums[i] = product / nums[i]
    }

    return nums
}

func longestConsecutive(nums []int) int {
	vals := make(map[int]int)
	for _, v := range nums {
		vals[v] += 1
	}

	sequences := make(map[int]int)
	keys := make([]int, len(vals))

	i := 0
	for k := range vals {
		keys[i] = k
		i++
	}
	
	start := keys[0]
	for _, v := range keys {
		if _, ok := vals[v-1]; !ok {
			start = v
			sequences[start] = 1
		}
		sequences[start]++
	}


	maxSequences := 0
	for _, length := range sequences {
		maxSequences = max(length, maxSequences)
	}

	return maxSequences


}
