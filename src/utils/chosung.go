package utils

import (
	"strings"
	"unicode"
)

var (
	hangulCHO = []string{"ㄱ", "ㄲ", "ㄴ", "ㄷ", "ㄸ", "ㄹ", "ㅁ", "ㅂ", "ㅃ", "ㅅ", "ㅆ", "ㅇ", "ㅈ", "ㅉ", "ㅊ", "ㅋ", "ㅌ", "ㅍ", "ㅎ"}
	hangulBASE = rune('가')
)

func ExtractHangul(input string) string {
	words := strings.Fields(input)
	var resultWords []string

	for _, word := range words {
		hangul := ""
		for _, s := range word {
			if unicode.Is(unicode.Hangul, s) {
				hangul += string(s)
			}
		}
		if hangul != "" {
			resultWords = append(resultWords, hangul)
		}
	}

	return strings.Join(resultWords, " ")
}

func ExtractChosung(word string) string {
	hangul := ExtractHangul(word)
	chosung := ""

	if len(hangul) == 0 {
		return ""
	}

	for _, s := range hangul {
		if s != ' ' {
			temp := s - hangulBASE
			indx := temp / 588
			chosung += hangulCHO[indx]
		} else {
			chosung += " "
		}
	}
	return chosung
}
