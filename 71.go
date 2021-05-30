package leetcode

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0)

	var a []int32
	for _, i := range path {
		if i == '/' {

			if l := len(a); l > 0 {
				//不需要弹出
				if l == 1 && a[0] == '.' {
					a = nil
				} else if l == 2 && a[0] == '.' {
					if j := len(stack); j > 0 {
						stack = stack[0 : j-1]
					}
					a = nil
				} else {
					stack = append(stack, string(a))
					a = nil
				}
			}
			continue
		} else {
			a = append(a, i)
		}
	}
	if l := len(a); l > 0 {
		//不需要弹出
		if l == 1 && a[0] == '.' {
			a = nil
		} else if l == 2 && a[0] == '.' {
			if j := len(stack); j > 0 {
				stack = stack[0 : j-1]
			}
			a = nil
		} else {
			stack = append(stack, string(a))
			a = nil
		}
	}
	var sb strings.Builder
	if len(stack) > 0 {
		for _, s := range stack {
			sb.WriteByte('/')
			sb.WriteString(s)
		}
	} else {
		sb.WriteByte('/')
	}

	return sb.String()
}
