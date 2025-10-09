package main

import (
	"fmt"
	"strings"
)

func reverseVowel (s string)string {
	a := strings.Split(s, "")

	i,j := 0, len(a) -1

	for {
		if i >= j {
			break
		}

		if !isVowel(a[i]) && !isVowel(a[j]) {
			i++
			j--
		} else if isVowel(a[i]) && isVowel(a[j]) {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		} else if isVowel(a[i]) {
			j--
		} else if isVowel(a[j]) {
			i++
		}
	}

	return strings.Join(a, "")

}


func isVowel(s string) bool {
	vowel := "AaEeIiOoUu"
	return strings.Contains(vowel, s)
}


func main() {
	fmt.Println(reverseVowel("Abcda"))
}