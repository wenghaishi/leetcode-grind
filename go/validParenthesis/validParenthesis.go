package main

func validParenthesis(s string) bool {
	pairs := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}

	var stack []byte

	for i := 0; i < len(s); i++ {
		c := s[i]

		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(s) <= 0 {
				return false
			}

			if c != pairs[s[i]] {
				return false
			}

			stack = stack[:len(stack) -1]
		}
	}

	return len(stack) == 0
}