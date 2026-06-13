func longestCommonPrefix(strs []string) string {

	out := ""
	if len(strs) == 1 {
		return strs[0]
	}

	correct := true
	var next rune

	//going through letters of the words
	for j := 0; j < len(strs[0]); j++ {

		//going throught all words for letter possition j
		for i := 1; i <= len(strs)-1; i++ {

            //checking for the shortest word
			if len(strs[i]) <= j {
                correct = false
				break
			}

			if []rune(strs[i-1])[j] != []rune(strs[i])[j] {
				correct = false
				break
			} else {
				next = []rune(strs[i-1])[j]
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
