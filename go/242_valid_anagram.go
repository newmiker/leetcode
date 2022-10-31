package main

//func isAnagram(s string, t string) bool {
//	counters := make(map[rune]int32)
//
//	// Fill counter-map with first string.
//	for _, r := range []rune(s) {
//		counters[r]++
//	}
//	// Decrement going through the second string.
//	for _, r := range []rune(t) {
//		counters[r]--
//	}
//	// Iterate and find non-zero values.
//	for _, v := range counters {
//		if v != 0 {
//			return false
//		}
//	}
//
//	return true
//}

func isAnagramOptimized(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	counters := [26]uint16{}

	for i := range s {
		counters[s[i]%26]++
		counters[t[i]%26]--
	}
	for _, v := range counters {
		if v != 0 {
			return false
		}
	}

	return true
}
