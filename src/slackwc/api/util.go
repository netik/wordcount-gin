/*
 * John Adams
 * jna@retina.net
 * 8/30/2017
 *
 * slackwc: util.go
 * Utility Functions
 *
 */

package api

import (
   "strings"
   "unicode"
)

func stripPuncuation(str string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsPunct(r) {
            // if the rune is in unicode category P*, skip this rune.
            return -1
        }
        // else keep it in the string
        return r
    }, str)
}

func WordCounter(s string) map[string]int {
	/* return a map representing a case insensitive word count of words in s */
	/* filters and ignores puncuation */

	words := strings.Fields(stripPuncuation(s))
	wordCountMap := make(map[string]int)

	for _, word := range words {
		wordCountMap[strings.ToLower(word)]++
	}

	return wordCountMap
}
