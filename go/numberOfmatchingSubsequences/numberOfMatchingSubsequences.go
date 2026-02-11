package main

func numMatchingSubseq(s string, words []string) int {
	count := 0

	for _, v := range words {
		if len(v) > len(s) {
			continue
		}

		l, r := 0,0

		for l < len(s) {
			if s[l] == v[r] {
				l++
				r++

				if r == len(v) {
					count++
					break
				}
			} else {
				l++
			}
		}
	}

	return count
}