package main

func longestCommonPrefix(strs []string) string {
	// commonPrefix := make(map[string]int)
	if len(strs) <= 0 {
		return ""
	}
	minlen := 999999999
	for _, str := range strs {
		if len(str) < minlen {
			minlen = len(str)
		}
	}
	longestCommon := ""
	for i := 0; i < minlen; i++ {
		common := ""
		isHaveCommon := true

		for _, str := range strs {
			if common == "" {
				common = string(str[i])
			}
			if common != "" {
				if string(str[i]) != common {
					if i == 0 {
						return ""
					}
					isHaveCommon = false
					break
				}
			}
		}

		if isHaveCommon {
			str := strs[0][i]
			longestCommon += string(str)
		}
	}
	if longestCommon != "" {
		return longestCommon
	}
	return ""
}

// ["flower","flow","flight"]
// ["dog","racecar","car"]
// ["aca","cba"]
