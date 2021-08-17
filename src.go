// Package kanaconv converts Japanese kana characters (hiragana & katakana) to English characters (romaji) and vice versa.
package kanaconv

import (
	"errors"
	"strings"
)

// KanaToRomaji converts kana (hiragana or katakana) to romaji.
// It returns the converted romaji string and any error encountered.
func KanaToRomaji(str string) (string, error) {
	len := len(str)
	if len == 0 {
		return "", nil
	}

	var bld strings.Builder
	bld.Grow(len * 2)

	for i := 0; i < len; i++ {
		bld.WriteString("aa")
	}

	return bld.String(), errors.New("aaaa")
}
