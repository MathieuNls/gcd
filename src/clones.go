package gcd

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	type2 = iota
	type3 = iota
)

//BijectiveMorphisme stores how to prettyfy a code snippet into another one
type BijectiveMorphisme struct {
	morphismes       []*BijectiveReplacement
	source           string
	encodedSource    string
	cloneType        int
	simlarity        float32
	sourceLines      []string
	sourcePrettyfied []string
	limit            int
}

//BijectiveReplacement contains a from a and a to string for replacements
type BijectiveReplacement struct {
	from string
	to   string
}

func New(source string, limit int) *BijectiveMorphisme {

	bm := &BijectiveMorphisme{
		source: source,
		limit:  limit,
	}

	bm.encodedSource, _, _ = encodeCode(source)
	bm.sourceLines = strings.Split(bm.encodedSource, "\n")
	bm.sourcePrettyfied = prettyfy(bm.source)
	return bm
}

func prettyfy(code string) []string {

	result := strings.Replace(code, "\t", "", -1)
	result = strings.Replace(code, "\n", "", -1)
	result = strings.Replace(code, "	", "", -1)

	for _, keyword := range []string{";", "}", "{", "(", ")"} {
		result = strings.Replace(result, keyword, keyword+"\n", -1)
	}

	results := []string{}

	for _, str := range strings.Split(result, "\n") {
		if str != "" {
			results = append(results, str)
		}
	}

	return results
}

func (bm *BijectiveMorphisme) check(code string) (int, float32, string) {

	targetEncoded, _, _ := encodeCode(code)

	targetLines := strings.Split(targetEncoded, "\n")

	compareCode := func(source []string, target []string) (int, int, int) {

		LCSLines := 0
		LCSStartSource := 0
		LCSStartTarget := 0

		for i := 0; i < len(source); i++ {

			for j := 0; j < len(target); j++ {

				if source[i] == target[j] {

					currentMatchLength := 0

					for k := 0; k+i < len(source) && k+j < len(target); k++ {

						if source[k+i] == target[k+j] {
							currentMatchLength++
						} else {
							break
						}
					}

					if currentMatchLength > LCSLines {
						LCSStartSource = i
						LCSStartTarget = j
						LCSLines = currentMatchLength
					}

				}
			}
		}

		bm.simlarity = float32(LCSLines) / float32(len(source)) * 100.0

		return LCSStartSource, LCSStartTarget, LCSLines
	}

	var startSource, startTarget, lines int

	if len(bm.sourceLines) > len(targetLines) {

		startSource, startTarget, lines = compareCode(bm.sourceLines, targetLines)
	} else {
		startSource, startTarget, lines = compareCode(targetLines, bm.sourceLines)
	}

	if lines > bm.limit {

		spliter := func(code rune) bool {

			if code == ' ' || code == '\n' {
				return true
			}
			return false
		}

		target := prettyfy(code)

		bijs := []*BijectiveReplacement{}

		for i := 0; i < lines; i++ {

			sourceLine := bm.sourcePrettyfied[startSource+i]

			for _, keyword := range []string{";", "}", "{", "(", ")"} {
				sourceLine = strings.Replace(sourceLine, keyword, " ", -1)
			}

			targetLine := target[startTarget+i]

			for _, keyword := range []string{";", "}", "{", "(", ")"} {
				targetLine = strings.Replace(targetLine, keyword, " ", -1)
			}

			sourceLines := strings.FieldsFunc(sourceLine, spliter)
			targetLines := strings.FieldsFunc(targetLine, spliter)

			bijs = append(bijs, LevenshteinDistance(sourceLines, len(sourceLines), targetLines, len(targetLines))...)

		}

		bijs, err := reduceBijections(bijs)

		bm.morphismes = bijs

		if err != nil {
			bm.cloneType = type3
		} else {
			bm.cloneType = type2
		}

	}

	return bm.cloneType, bm.simlarity, bm.transform(code)

}

func (bm *BijectiveMorphisme) transform(code string) string {

	if bm.cloneType == type2 {
		for _, bijr := range bm.morphismes {
			code = strings.Replace(code, bijr.from, bijr.to, -1)
		}
		return code
	}
	return ""

}

func reduceBijections(bijs []*BijectiveReplacement) ([]*BijectiveReplacement, error) {

	bigMap := make(map[string]string)

	for _, bij := range bijs {

		if val, ok := bigMap[bij.from]; ok && val != bij.to {
			return bijs, errors.New("Unconsistent replacement")
		}
		bigMap[bij.from] = bij.to

	}

	results := []*BijectiveReplacement{}

	for k, v := range bigMap {
		results = append(results, &BijectiveReplacement{k, v})
	}

	return results, nil

}

func printBijection(bijs []*BijectiveReplacement) {

	for _, bij := range bijs {
		fmt.Printf("%s -> %s\n", bij.from, bij.to)
	}
}

// LevenshteinDistance len_s and len_t are the number of characters in string s and t respectively
func LevenshteinDistance(s []string, lenS int, t []string, lenT int) []*BijectiveReplacement {

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

	ldFromS := append(replacements, LevenshteinDistance(s, lenS-1, t, lenT)...)
	ldFromT := append(replacements, LevenshteinDistance(s, lenS, t, lenT-1)...)
	ldFromSAndT := append(replacements, LevenshteinDistance(s, lenS-1, t, lenT-1)...)

	return minimumBijectiveReplacement(ldFromS, ldFromT, ldFromSAndT)
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
