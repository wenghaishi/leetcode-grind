package main

import "fmt"

func editDistance(word1, word2 string) int {
	word1Length, word2Length := len(word1), len(word2)
	maxMatches := 0
	if word1Length > word2Length {
		beginning, ending := 0, word2Length-1
		for ending < word1Length {
			matches := 0

			j := beginning

			for _, v := range word2 {
				if v == int32(word1[j]) {
					matches++
				}
				j++
			}

			if matches > maxMatches {
				maxMatches = matches
			}
			beginning++
			ending++
		}
	}

	additions := word1Length - word2Length
	replacements := word2Length - maxMatches

	return additions + replacements
}

func main() {
	str1, str2 := "horse", "ros"

	res := editDistance(str1, str2)
	fmt.Println(res)
}
