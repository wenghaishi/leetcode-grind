package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("hi")
}

func validPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		left := rune(s[l])
		right := rune(s[r])
		isLeftChar, isRightChar := false, false

		for !isLeftChar {
			if unicode.IsLetter(left) {
				isLeftChar = true
			} else {
				l++
			}
		}

		for !isRightChar {
			if unicode.IsLetter(right) {
				isRightChar = true
			} else {
				r--
			}
		}

	}
}
