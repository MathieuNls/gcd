package gcd

import "strings"

func containsKeyWord(str string, keywords []string) bool {

	for _, keyword := range keywords {
		if keyword == strings.Trim(str, " ") {
			return true
		}
	}
	return false
}

func containsLineBreak(str string) bool {

	return strings.Contains(str, "\n")
}

func removeLineBreaks(str string) string {
	return strings.Replace(strings.Replace(str, "\n", "", -1), "\t", "", -1)
}

func trimCarriageReturn(str []string) []string {
	for i := 0; i < len(str); i++ {

		if strings.HasPrefix(str[i], "\n\t") {
			str[i] = strings.Replace(str[i], "\n\t", "", 1)
		} else if strings.HasPrefix(str[i], "\n") {
			str[i] = strings.Replace(str[i], "\n", "", 1)
		} else if strings.HasPrefix(str[i], "\t") {
			str[i] = strings.Replace(str[i], "\t", "", 1)
		}
	}

	return str
}

func encodeCode(code string) (string, []int, []int) {

	keywords := []string{"{", "}", "+", "=", "]", "[", ")", "(", ";", "for", "if", "while", "<", ".", "-", ">"}

	for _, keyword := range keywords {
		code = strings.Replace(code, keyword, " "+keyword+" ", -1)
	}

	code = strings.Replace(code, "\t", "", -1)
	code = strings.Replace(code, "\n", "", -1)

	encodedCode := ""
	beginBlocks := []int{}
	endBlocks := []int{}

	for _, str := range strings.Split(code, " ") {

		if str != "" {

			if containsKeyWord(str, keywords) {
				encodedCode += strings.Trim(str, " ")

				if containsKeyWord(str, []string{";", "}", "{", "(", ")"}) {

					encodedCode += "\n"
				}

			} else {
				encodedCode += "#"
			}
		}

	}

	for index, char := range encodedCode {

		if char == '{' {
			beginBlocks = append(beginBlocks, index)
		} else if char == '}' {
			endBlocks = append(endBlocks, index)
		}
	}

	return encodedCode, beginBlocks, endBlocks
}

func LCS(first, second string) (int, string) {

	a, b := []rune(first), []rune(second)

	lengths := make([][]int, len(a)+1)
	for i := 0; i <= len(a); i++ {
		lengths[i] = make([]int, len(b)+1)
	}

	// row 0 and column 0 are initialized to 0 already
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				lengths[i+1][j+1] = lengths[i][j] + 1
			} else if lengths[i+1][j] > lengths[i][j+1] {
				lengths[i+1][j+1] = lengths[i+1][j]
			} else {
				lengths[i+1][j+1] = lengths[i][j+1]
			}
		}
	}

	// read the substring out from the matrix
	s := make([]rune, 0, lengths[len(a)][len(b)])
	for x, y := len(a), len(b); x != 0 && y != 0; {
		if lengths[x][y] == lengths[x-1][y] {
			x--
		} else if lengths[x][y] == lengths[x][y-1] {
			y--
		} else {
			s = append(s, a[x-1])
			x--
			y--
		}
	}

	// reverse string
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return len(s), string(s)
}
