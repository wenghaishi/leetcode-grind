package main

import "fmt"

func generateParenthesis(n int) []string {
	var solution []string

	var helper func(str string)
	helper = func(str string) {
		openCount, closeCount := 0, 0
		for _, i := range str {
			if i == '(' {
				openCount++
			}
			if i == ')' {
				closeCount++
			}
		}

		// Reached n number of parenthesis
		if openCount == n && closeCount == n {
			solution = append(solution, str)
		}

		if closeCount < openCount {
			helper(str + ")")
		}

		if openCount < n {
			helper(str + "(")
		}
	}

	helper("")

	return solution
}

func main() {
	test := generateParenthesis(5)
	fmt.Println(test)
}
