package main

func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for _, s := range strs[1:] {
		prefix = commonPrefix(prefix, s)
	}

	return prefix
}

func commonPrefix(s1, s2 string) string {
	shortest := []rune(s1)
	longest := []rune(s2)
	if len(s2) < len(s1) {
		shortest = []rune(s2)
		longest = []rune(s1)
	}

	for i := 0; i < len(shortest); i++ {
		if shortest[i] != longest[i] {
			return string(shortest[:i])
		}
	}

	return string(shortest)
}
