package main

import (
	"fmt"
	"strings"
)

func mergeAlt(w1, w2 string) string {
	w1Arr := strings.Split(w1, "")
	w2Arr := strings.Split(w2, "")

	i,j := 0,0

	var combinedArr []string

	for i < len(w1) || j < len(w2) {
		if i < len(w1) {
			combinedArr = append(combinedArr, w1Arr[i])
			i++
		}

		if j < len(w2) {
			combinedArr = append(combinedArr, w2Arr[j])
			j++
		}
	}

	return strings.Join(combinedArr, "")

}


func main() {
	result := mergeAlt("ac", "bdefg")
	fmt.Println(result) // Output: apbqcr
}