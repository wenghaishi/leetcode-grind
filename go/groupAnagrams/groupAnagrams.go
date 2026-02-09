package main

import "sort"

func groupAnagrams(strs []string) [][]string {

	strsMap := make(map[string][]string)

	for _, v := range strs {
		strRune := []rune(v)

		sort.Slice(strRune, func(i,j int) bool {
			return strRune[i] < strRune[j]
		} )

		strString := string(strRune)
		strsMap[strString] = append(strsMap[strString], v)
	}
	

	var solution [][]string

	for _, value := range strsMap {
		solution = append(solution, value)
	}

	return solution

}
