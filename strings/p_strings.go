package strings

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

type BijectiveMorphisme struct {
	morphismes          []*BijectiveReplacement
	source              string
	target              string
	transformed         string
	LevenshteinDistance int
	lcs                 int
	lcsString           string
}

type BijectiveReplacement struct {
	from string
	to   string
}

func (bm *BijectiveMorphisme) find() {
	alphaAlphabet := strings.Split(bm.source, " ")
	betaAlphabet := strings.Split(bm.target, " ")

	bm.morphismes = LevenshteinDistance2(alphaAlphabet, len(alphaAlphabet), betaAlphabet, len(betaAlphabet))
	bm.reduceMorphisme()
	printBijectivesReplacements(bm.morphismes)
	bm.apply()
	fmt.Printf("source = %s,\n target = %s,\n transo = %s\n", bm.source, bm.target, bm.transformed)

	printAlphabets(alphaAlphabet, betaAlphabet)
}

func (bm *BijectiveMorphisme) reduceMorphisme() error {

	knownMorphisme := make(map[string]*BijectiveReplacement)
	tmpReplacements := bm.morphismes
	bm.morphismes = nil

	for _, morphisme := range tmpReplacements {

		if _, contains := knownMorphisme[morphisme.from]; contains && knownMorphisme[morphisme.from].to != morphisme.to {
			return errors.New("Add")
		}
		bm.morphismes = append(bm.morphismes, morphisme)

	}

	return nil

}

func (bm *BijectiveMorphisme) apply() {

	bm.transformed = bm.source
	for _, br := range bm.morphismes {
		bm.transformed = strings.Replace(bm.transformed, br.from, br.to, -1)
	}
}

func printBijectivesReplacements(replacements []*BijectiveReplacement) {
	for _, replacement := range replacements {
		fmt.Printf("%s -> %s,", replacement.from, replacement.to)
	}
}

func printAlphabets(alphabets ...[]string) {

	for _, alphabet := range alphabets {
		fmt.Println()
		for _, token := range alphabet {
			fmt.Printf("%s,", token)
		}
	}
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

// LevenshteinDistance len_s and len_t are the number of characters in string s and t respectively
func LevenshteinDistance2(s []string, lenS int, t []string, lenT int) []*BijectiveReplacement {

	var replacements []*BijectiveReplacement

	/* base case: empty strings */
	if lenS == 0 {
		return nil
	} else if lenT == 0 {
		return nil
	}

	/* test if last characters of the strings match */
	if s[lenS-1] != t[lenT-1] {

		replacements = append(replacements, &BijectiveReplacement{
			string(s[lenS-1]),
			string(t[lenT-1]),
		})
	}

	ldFromS := append(replacements, LevenshteinDistance2(s, lenS-1, t, lenT)...)
	ldFromT := append(replacements, LevenshteinDistance2(s, lenS, t, lenT-1)...)
	ldFromSAndT := append(replacements, LevenshteinDistance2(s, lenS-1, t, lenT-1)...)

	return minimumBijectiveReplacement(ldFromS, ldFromT, ldFromSAndT)
}

// LevenshteinDistance len_s and len_t are the number of characters in string s and t respectively
func LevenshteinDistance(s string, lenS int, t string, lenT int) int {

	var cost int

	/* base case: empty strings */
	if lenS == 0 {
		return lenT
	} else if lenT == 0 {
		return lenS
	}

	/* test if last characters of the strings match */
	if s[lenS-1] == t[lenT-1] {
		cost = 0
	} else {
		cost = 1
	}

	/* return minimum of delete char from s, delete char from t, and delete char from both */
	return minimum(LevenshteinDistance(s, lenS-1, t, lenT)+1,
		LevenshteinDistance(s, lenS, t, lenT-1)+1,
		LevenshteinDistance(s, lenS-1, t, lenT-1)+cost)
}

func minimumBijectiveReplacement(replacements ...[]*BijectiveReplacement) []*BijectiveReplacement {

	minimum := math.MaxInt32
	var minimumReplacement []*BijectiveReplacement

	for _, value := range replacements {
		if len(value) < minimum {
			minimum = len(replacements)
			minimumReplacement = value
		}
	}

	return minimumReplacement
}

func minimum(nums ...int) int {

	minimum := math.MaxInt32

	for _, value := range nums {

		if value < minimum {
			minimum = value
		}
	}
	return minimum
}

func pString(a string, b string) {

}
