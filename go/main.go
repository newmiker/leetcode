package main

import (
	"fmt"
)

func main() {
	fmt.Println(reverseVowels("hello"))
}

func reverseVowels(s string) string {
	runes := []rune(s)
	endIdx := len(runes) - 1
	for i, r := range runes {
		if !vowels[r] {
			continue
		}
		if i >= endIdx {
			return string(runes)
		}

		for j := endIdx; ; j-- {
			if j == i {
				return string(runes)
			}
			if vowels[runes[j]] {
				runes[i], runes[j] = runes[j], runes[i]
				endIdx = j - 1
				break
			}
		}
	}

	return string(runes)
}

var vowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
	'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
}
