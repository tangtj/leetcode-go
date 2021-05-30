package leetcode

func isValid(s string) bool {

	stack := make([]int32, 0)
	top := 0

	for _, c := range s {

		if top <= 0 {
			if c == '(' || c == '[' || c == '{' {
				stack = append(stack, c)
				top++
			} else {
				return false
			}
		} else if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
			top++
		} else {
			t := stack[top-1]
			if c == ')' && t == '(' {
				stack = stack[:top-1]
				top--
			} else if c == ']' && t == '[' {
				stack = stack[:top-1]
				top--
			} else if c == '}' && t == '{' {
				stack = stack[:top-1]
				top--
			} else {
				stack = append(stack, c)
				top++
			}
		}
	}
	return top == 0
}
