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
		return "", errors.New("all characters must be kana (3-bit unicode characters)")
	}

	var sb strings.Builder
	sb.Grow(len(str) * 2)

	for i := 0; i < len(str); i += byteCount {
		var romaji string

		switch getKanaRune(str[i], str[i+1], str[i+2]) {
		// basic
		case 'あ', 'ア':
			sb.WriteByte('a')
			continue
		case 'い', 'イ':
			sb.WriteByte('i')
			continue
		case 'う', 'ウ':
			sb.WriteByte('u')
			continue
		case 'え', 'エ':
			sb.WriteByte('e')
			continue
		case 'お', 'オ':
			sb.WriteByte('o')
			continue
		case 'ん', 'ン':
			romaji = "n"
			goto Romaji
		// basic dakuten
		case 'ゔ', 'ヴ':
			romaji = "vu"
			goto Romaji
		// k
		case 'か', 'カ':
			romaji = "ka"
			goto Romaji
		case 'き', 'キ':
			romaji = "ki"
			goto Romaji
		case 'く', 'ク':
			romaji = "ku"
			goto Romaji
		case 'け', 'ケ':
			romaji = "ke"
			goto Romaji
		case 'こ', 'コ':
			romaji = "ko"
			goto Romaji
		// k dakuten (g)
		case 'が', 'ガ':
			romaji = "ga"
			goto Romaji
		case 'ぎ', 'ギ':
			romaji = "gi"
			goto Romaji
		case 'ぐ', 'グ':
			romaji = "gu"
			goto Romaji
		case 'げ', 'ゲ':
			romaji = "ge"
			goto Romaji
		case 'ご', 'ゴ':
			romaji = "go"
			goto Romaji
		// s
		case 'さ', 'サ':
			romaji = "sa"
			goto Romaji
		case 'し', 'シ':
			romaji = "shi"
			goto Romaji
		case 'す', 'ス':
			romaji = "su"
			goto Romaji
		case 'せ', 'セ':
			romaji = "se"
			goto Romaji
		case 'そ', 'ソ':
			romaji = "so"
			goto Romaji
		// s dakuten (z)
		case 'ざ', 'ザ':
			romaji = "za"
			goto Romaji
		case 'じ', 'ジ':
			romaji = "ji"
			goto Romaji
		case 'ず', 'ズ':
			romaji = "zu"
			goto Romaji
		case 'ぜ', 'ゼ':
			romaji = "ze"
			goto Romaji
		case 'ぞ', 'ゾ':
			romaji = "zo"
			goto Romaji
		// t
		case 'た', 'タ':
			romaji = "ta"
			goto Romaji
		case 'ち', 'チ':
			romaji = "chi"
			goto Romaji
		case 'つ', 'ツ':
			romaji = "tsu"
			goto Romaji
		case 'て', 'テ':
			romaji = "te"
			goto Romaji
		case 'と', 'ト':
			romaji = "to"
			goto Romaji
		// t dakuten (d)
		case 'だ', 'ダ':
			romaji = "da"
			goto Romaji
		case 'ぢ', 'ヂ':
			romaji = "ji"
			goto Romaji
		case 'づ', 'ヅ':
			romaji = "zu"
			goto Romaji
		case 'で', 'デ':
			romaji = "de"
			goto Romaji
		case 'ど', 'ド':
			romaji = "do"
			goto Romaji
		// n
		case 'な', 'ナ':
			romaji = "na"
			goto Romaji
		case 'に', 'ニ':
			romaji = "ni"
			goto Romaji
		case 'ぬ', 'ヌ':
			romaji = "nu"
			goto Romaji
		case 'ね', 'ネ':
			romaji = "ne"
			goto Romaji
		case 'の', 'ノ':
			romaji = "no"
			goto Romaji
		// h
		case 'は', 'ハ':
			romaji = "ha"
			goto Romaji
		case 'ひ', 'ヒ':
			romaji = "hi"
			goto Romaji
		case 'ふ', 'フ':
			romaji = "fu"
			goto Romaji
		case 'へ', 'ヘ':
			romaji = "he"
			goto Romaji
		case 'ほ', 'ホ':
			romaji = "ho"
			goto Romaji
		// h dakuten (b)
		case 'ば', 'バ':
			romaji = "ba"
			goto Romaji
		case 'び', 'ビ':
			romaji = "bi"
			goto Romaji
		case 'ぶ', 'ブ':
			romaji = "bu"
			goto Romaji
		case 'べ', 'ベ':
			romaji = "be"
			goto Romaji
		case 'ぼ', 'ボ':
			romaji = "bo"
			goto Romaji
		// h dakuten (p)
		case 'ぱ', 'パ':
			romaji = "pa"
			goto Romaji
		case 'ぴ', 'ピ':
			romaji = "pi"
			goto Romaji
		case 'ぷ', 'プ':
			romaji = "pu"
			goto Romaji
		case 'ぺ', 'ペ':
			romaji = "pe"
			goto Romaji
		case 'ぽ', 'ポ':
			romaji = "po"
			goto Romaji
		// m
		case 'ま', 'マ':
			romaji = "ma"
			goto Romaji
		case 'み', 'ミ':
			romaji = "mi"
			goto Romaji
		case 'む', 'ム':
			romaji = "mu"
			goto Romaji
		case 'め', 'メ':
			romaji = "me"
			goto Romaji
		case 'も', 'モ':
			romaji = "mo"
			goto Romaji
		// y
		case 'や', 'ヤ':
			romaji = "ya"
			goto Romaji
		case 'ゆ', 'ユ':
			romaji = "yu"
			goto Romaji
		case 'よ', 'ヨ':
			romaji = "yo"
			goto Romaji
		// r
		case 'ら', 'ラ':
			romaji = "ra"
			goto Romaji
		case 'り', 'リ':
			romaji = "ri"
			goto Romaji
		case 'る', 'ル':
			romaji = "ru"
			goto Romaji
		case 'れ', 'レ':
			romaji = "re"
			goto Romaji
		case 'ろ', 'ロ':
			romaji = "ro"
			goto Romaji
		// w
		case 'わ', 'ワ':
			romaji = "wa"
			goto Romaji
		case 'ゐ', 'ヰ':
			romaji = "wi"
			goto Romaji
		case 'ゑ', 'ヱ':
			romaji = "we"
			goto Romaji
		case 'を', 'ヲ':
			romaji = "wo"
			goto Romaji
		}

	Romaji:
		sb.WriteString(romaji)
		continue
	}

	return sb.String(), nil
}

//	getKanaRune converts a 3-bit hex value to its unicode code point
// 		[1110(0011)]+[10(00 0001)]+[10(00 0010)] -> [(0011)+(00 0001)+(00 0010)]
func getKanaRune(byte1 byte, byte2 byte, byte3 byte) rune {
	int1 := uint16(byte1&0b_0000_1111) << 12 // remove 4 leading bytes, move by 12 bytes
	int2 := uint16(byte2&0b_0011_1111) << 6  // remove 2 leading bytes, move by 6 bytes
	int3 := uint16(byte3 & 0b_0011_1111)     // remove 2 leading bytes
	return rune(int1 + int2 + int3)
}
