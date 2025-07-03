package piscine

func TrimAtoi(s string) int {
	result := 0
	isNegative := false
	digitFound := false

	for _, r := range s {
		if r >= '0' && r <= '9' {
			result = result*10 + int(r-'0')
			digitFound = true
		} else if r == '-' && !digitFound {
			isNegative = true
		}
	}

	if result == 0 && !digitFound {
		return 0
	}

	if isNegative {
		return -result
	}
	return result
}
