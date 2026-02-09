package main

import "fmt"

func ransomNote(ransomNote string, magazine string) bool {
	letters := make(map[rune]int)

	for _, v := range magazine {
		letters[v]++
	}

	for _, r := range ransomNote {
		letters[r]--

		if letters[r] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(ransomNote("aaa", "ab"))
}
