package main

func longestCommonPrefix(strs []string) string {
	// commonPrefix := make(map[string]int)
	minlen := 999999999
	for _, str := range strs {
		if len(str) < minlen {
			minlen = len(str)
		}
	}

	return ""
}

// ["flower","flow","flight"]
// ["dog","racecar","car"]
