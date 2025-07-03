package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	newWord := true

	for i, r := range runes {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			if newWord {
				// If it's a letter and start of a word, make uppercase
				if r >= 'a' && r <= 'z' {
					runes[i] = r - 32
				}
				newWord = false
			} else {
				// If it's a letter, ensure it's lowercase
				if r >= 'A' && r <= 'Z' {
					runes[i] = r + 32
				}
			}
		} else {
			newWord = true // reset at non-alphanumeric boundary
		}
	}
	return string(runes)
}
