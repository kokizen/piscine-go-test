package piscine

func Index(s string, toFind string) int {
	sLen := len(s)
	fLen := len(toFind)

	if fLen == 0 {
		return 0
	}

	if fLen > sLen {
		return -1
	}

	for i := 0; i <= sLen-fLen; i++ {
		if s[i:i+fLen] == toFind {
			return i
		}
	}
	return -1
}
