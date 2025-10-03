package main

import (
	"fmt"
)

func main() {
	s := "abcabcbb"

	maxLen, start := 0, 0
	lastIdx := make([]int, 128) // Store the last index we seen for current char

	for end := 0; end < len(s); end++ {
		currentC := s[end] // get the current char

		println("Loop:", end)
		println("Top Start: ", start)
		println("Top End: ", end)

		fmt.Println("Current Char:", currentC)
		fmt.Println("Last Idx:", lastIdx[currentC])

		// Modify when we have repeat char, slide forward the window
		if lastIdx[currentC] > start {
			println("Inside first IF")
			start = lastIdx[currentC]
		}

		if end-start+1 > maxLen {
			println("Inside second IF")
			maxLen = end - start + 1
		}

		lastIdx[currentC] = end + 1 // Always plus 1 because we want to move forward and update the char if has repeat

		println("Bottom Start: ", start)
		println("Bottom End: ", end)
		fmt.Println("Last Idx:", lastIdx[currentC])
		println()
	}
}

// Can simplify the above implementation using an while loop
// if strings.Contains(s[start:end], string(s[end])) {
// 		start++
// 		continue
// 	}

// 	end++
// 	maxLen = max(maxLen, end-start)
