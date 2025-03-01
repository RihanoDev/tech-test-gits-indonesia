package main

import "fmt"

func isBalanced(s string) string {
	stack := []rune{}

	bracketPair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	for _, ch := range s {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' || ch == ']' || ch == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != bracketPair[ch] {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	fmt.Println("First Input Result: ", isBalanced("{ [ ( ) ] }"))
	fmt.Println("Second Input Result: ", isBalanced("{ [ ( ] ) }"))
	fmt.Println("Third Input Result: ", isBalanced("{ ( ( [ ] ) [ ] ) [ ] }"))
}
