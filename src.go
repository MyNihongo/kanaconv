//	Package kanaconv converts Japanese kana characters (hiragana & katakana) to English characters (romaji) and vice versa.
package kanaconv

import (
	"errors"
	"strings"
)

//	KanaToRomaji converts kana (hiragana or katakana) to romaji.
//	It returns the converted romaji string and any error encountered.
func KanaToRomaji(str string) (string, error) {
	const byteCount = 3
	if len(str) == 0 {
		return "", nil
	} else if len(str)%byteCount != 0 {
		// kana is a 3-bit unicode char
		return "", errors.New("all characters must be unicode")
	}

	var bld strings.Builder
	bld.Grow(len(str) * 2)

	for i := 0; i < len(str); i += byteCount {
		switch getKanaRune(str[i], str[i+1], str[i+2]) {
		case 'あ', 'ア':
			bld.WriteByte('a')
		case 'い', 'イ':
			bld.WriteByte('i')
		case 'う', 'ウ':
			bld.WriteByte('u')
		case 'え', 'エ':
			bld.WriteByte('e')
		case 'お', 'オ':
			bld.WriteByte('o')
		}
	}

	return bld.String(), nil
}

//	getKanaRune converts a 3-bit hex value to its unicode code point
// 		[1110(0011)]+[10(00 0001)]+[10(00 0010)] -> [(0011)+(00 0001)+(00 0010)]
func getKanaRune(byte1 byte, byte2 byte, byte3 byte) rune {
	int1 := int16(byte1&0b_0000_1111) << 12 // remove 4 leading bytes, move by 12 bytes
	int2 := int16(byte2&0b_0011_1111) << 6  // remove 2 leading bytes, move by 6 bytes
	int3 := int16(byte3 & 0b_0011_1111)     // remove 2 leading bytes
	return rune(int1 + int2 + int3)
}
