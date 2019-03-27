package main

func letterCombinations(digits string) []string {
	result := []string{}

	var phonePad = map[string][]string{
		"0": []string{},
		"1": []string{},
		"2": []string{"a", "b", "c"},
		"3": []string{"d", "e", "f"},
		"4": []string{"g", "h", "i"},
		"5": []string{"j", "k", "l"},
		"6": []string{"m", "n", "o"},
		"7": []string{"p", "q", "r", "s"},
		"8": []string{"t", "u", "v"},
		"9": []string{"w", "x", "y", "z"},
	}
	for len(digits) > 0 {
		if string(digits[0]) == "0" {
			continue
		}
		keys := phonePad[string(digits[0])]
		digits = digits[1:]
		if len(result) == 0 {
			result = append(result, keys...)
			continue
		}
		temp := make([]string, len(result))
		copy(temp, result)
		result = []string{}
		for len(temp) > 0 {
			s := temp[0]
			temp = temp[1:]
			for _, key := range keys {
				result = append(result, s+key)
			}
		}

	}

	return result
}

// "23": ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"]
// "2" : ["a","b","c"]
// "234" : ["adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"]
