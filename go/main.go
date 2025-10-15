package main

func main() {
	// input := []int{7, 7}
}

func lengthOfLongestSubstring(s string) int {
	longest := 0
	current := 0
	seen := make(map[rune]bool)
	for _, c := range s {
		if !seen[c] {
			seen[c] = true
			current += 1
		} else {
			current = 1
			seen = make(map[rune]bool)
			seen[c] = true
		}
		longest = max(current, longest)
	}

	return longest
}

