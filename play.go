package main

import "fmt"

func main() {
	a := "au"
	fmt.Println(lengthOfLongestSubstring(a))
}

func lengthOfLongestSubstring(s string) int {
	start := -1
	end := 0
	max_len := 0
	charMap := make(map[rune]int)
	i := len([]rune(s))
	if i <= 1 {
		return i
	}
	for i, c := range s {
		place, exists := charMap[c]
		if exists && place > start {
			start = place
			end = i
			fmt.Println("start: ", start)
			fmt.Println("end: ", end)
		} else {
			end = i
		}
		charMap[c] = i
		if end-start > max_len {
			max_len = end - start
		}
	}
	return max_len
}
