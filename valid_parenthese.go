package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	if len(s)%2 != 0 {
		return false
	}
	if s[0] == 41 || s[0] == 93 || s[0] == 125 {
		return false
	}
	stack := []rune{}
	for _, v := range s {
		if v == 40 || v == 91 || v == 123 {
			stack = append(stack, v)
			continue
		}
		if v == 41 {
			if stack[len(stack)-1] == 40 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if v == 93 {
			if stack[len(stack)-1] == 91 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
		if v == 125 {
			if stack[len(stack)-1] == 123 {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}

	}
	// log.Println(stack)
	if len(stack) > 0 {
		return false
	}
	return true
}

// "()", true
// "()[]{}", true
// "(]", false
// "([)]", false
// "{[]}", true
// "(([]){})", true
// "){", false
