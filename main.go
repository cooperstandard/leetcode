package main

import (
	"fmt"
	"slices"

	
)

func main() {
	input := []int{7, 7}
	output := topKFrequent(input, 1)
	fmt.Printf("%v: %v\n", slices.Equal(output, []int{7}), output)
}

