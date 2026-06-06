func longestCommonPrefix(strs []string) string {

	out := ""
    if len(strs) == 1 {
        return strs[0]
    }

	// check the size of the shortest word
	max := 2000
	for _, val := range strs {
		if max > len(val) {
			max = len(val)
		}
	}

	correct := true
	var next rune

	//going through letters of the word
	for j := 0; j < max; j++ {

		//going throught all words for letter possition j
		for i := 1; i <= len(strs)-1; i++ {

			strArr1 := []rune(strs[i-1])
			strArr2 := []rune(strs[i])

			if strArr1[j] != strArr2[j] {
				correct = false
				break
			} else {
				next = strArr1[j]
			}
		}

		if correct == true {
			out = out + string(next)
		} else {
			break
		}

	}

	return out
}
