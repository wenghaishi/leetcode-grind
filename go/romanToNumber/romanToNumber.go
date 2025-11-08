package main

func romanToNumber(s string) int {
	values := map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var number int
	for i := 0; i < len(s); i++ {
		if i + 1 < len(s) && values[s[i]] < values[s[i]] {
			number -= values[s[i]]
		} else {
			number += values[s[i]]
		}
	}

	return number
}