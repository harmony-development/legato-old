package postgres

const (
	minChar = byte('0')
	maxChar = byte('z')
)

func Rank(prev, next string) string {
	if prev == "" {
		prev = string(minChar)
	}
	if next == "" {
		next = string(maxChar)
	}

	rank := ""
	i := 0

	for {
		prevChar := getChar(prev, i, minChar)
		nextChar := getChar(next, i, maxChar)

		if prevChar == nextChar {
			rank += string(prevChar)
			i++
			continue
		}

		midChar := mid(prevChar, nextChar)
		if midChar == prevChar || midChar == nextChar {
			rank += string(prevChar)
			i++
			continue
		}

		rank += string(midChar)
		break
	}

	if rank >= next {
		return prev
	}

	return rank
}

func mid(prev, next byte) byte {
	return (prev + next) / 2
}

func getChar(s string, i int, defaultChar byte) byte {
	if i >= len(s) {
		return defaultChar
	}
	return s[i]
}
