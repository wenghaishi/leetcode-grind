package main

import "sort"

func reorganiseString(s string) string {
	byteMap := make(map[byte]int)

	for i :=0; i < len(s); i++ {
		byteMap[s[i]]++
	}

	var maxKey byte
	maxVal := -1

	for k, v := range byteMap {
		if v > maxVal {
			maxKey = k
			maxVal = v
		}
	}

	if maxVal > (len(s) + 1) /2 {
		return ""
	}

	var solution []byte

	for i := 0; i < len(s); i+=2 {
		if maxVal > 0 {
			solution[i] = maxKey
		} else {
			
		}
	}



	
}
