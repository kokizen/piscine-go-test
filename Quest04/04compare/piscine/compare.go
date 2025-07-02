package piscine

func Compare(a, b string) int {
	alen := len(a)
	blen := len(b)
	minlen := alen
	if blen < alen {
		minlen = blen
	}

	for i := 0; i < minlen; i++ {
		if a[i] < b[i] {
			return -1
		}
		if a[i] > b[i] {
			return 1
		}
	}

	if alen < blen {
		return -1
	}
	if alen > blen {
		return 1
	}
	return 0
}
