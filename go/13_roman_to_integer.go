package main

func romanToInt(s string) int {
	var acc int

	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	prev := 5000

	for _, letter := range s {
		v := m[letter]
		if v > prev {
			acc = acc - prev - prev + v
			continue
		}
		acc += v
		prev = v
	}

	return acc
}
