package main

func goodWaysToSplitString(s string) int {
	count := 0
	length := len(s)
	leftMap := make(map[byte]int)
	rightMap := make(map[byte]int)

	for i := 0; i < length; i++ {
		rightMap[s[i]]++
	}

	for j := 0; j < length-1; j++ {
		leftMap[s[j]]++
		rightMap[s[j]]--

		if rightMap[s[0]] < 0 {
			delete(rightMap, s[0])
		}

		if len(leftMap) == len(rightMap) {
			count++
		}
	}

	return count
}

aaaaa