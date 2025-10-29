package main

import (
	"fmt"
	"strings"
)

func lss(s string) int {
	stringArr := strings.Split(s, "")
	maxLength := 0
	
	for i := 0; i < len(stringArr); i++ {
		mySet := make(map[string]struct{})
		currentLength := 0

		for j := i; j < len(stringArr); j++ {
			_, exists := mySet[stringArr[j]]

			if exists {
				break
			} else {
				mySet[stringArr[j]] = struct{}{}
				currentLength++
				if currentLength > maxLength {
					maxLength = currentLength
				}
			}
		}
	}

	return maxLength
}


func main() {
	test1 := "abcdab"

	fmt.Println(lss(test1))

	test2 := "aaaaaa"

	fmt.Println(lss(test2))
} 