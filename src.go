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

	var rPrev string
	for i := 0; i < len(str); i += byteCount {
		var rStr string
		var rByte byte

		switch getKanaRune(str[i], str[i+1], str[i+2]) {
		// basic
		case 'あ', 'ア':
			rByte = 'a'
			goto RomajiByte
		case 'い', 'イ':
			rByte = 'i'
			goto RomajiByte
		case 'う', 'ウ':
			rByte = 'u'
			goto RomajiByte
		case 'え', 'エ':
			rByte = 'e'
			goto RomajiByte
		case 'お', 'オ':
			rByte = 'o'
			goto RomajiByte
		case 'ん', 'ン':
			rByte = 'n'
			goto RomajiByte
		// basic dakuten
		case 'ゔ', 'ヴ':
			rStr = "vu"
			goto RomajiString
		// k
		case 'か', 'カ':
			rStr = "ka"
			goto RomajiString
		case 'き', 'キ':
			rStr = "ki"
			goto RomajiString
		case 'く', 'ク':
			rStr = "ku"
			goto RomajiString
		case 'け', 'ケ':
			rStr = "ke"
			goto RomajiString
		case 'こ', 'コ':
			rStr = "ko"
			goto RomajiString
		// k dakuten (g)
		case 'が', 'ガ':
			rStr = "ga"
			goto RomajiString
		case 'ぎ', 'ギ':
			rStr = "gi"
			goto RomajiString
		case 'ぐ', 'グ':
			rStr = "gu"
			goto RomajiString
		case 'げ', 'ゲ':
			rStr = "ge"
			goto RomajiString
		case 'ご', 'ゴ':
			rStr = "go"
			goto RomajiString
		// s
		case 'さ', 'サ':
			rStr = "sa"
			goto RomajiString
		case 'し', 'シ':
			rStr = "shi"
			goto RomajiString
		case 'す', 'ス':
			rStr = "su"
			goto RomajiString
		case 'せ', 'セ':
			rStr = "se"
			goto RomajiString
		case 'そ', 'ソ':
			rStr = "so"
			goto RomajiString
		// s dakuten (z)
		case 'ざ', 'ザ':
			rStr = "za"
			goto RomajiString
		case 'じ', 'ジ':
			rStr = "ji"
			goto RomajiString
		case 'ず', 'ズ':
			rStr = "zu"
			goto RomajiString
		case 'ぜ', 'ゼ':
			rStr = "ze"
			goto RomajiString
		case 'ぞ', 'ゾ':
			rStr = "zo"
			goto RomajiString
		// t
		case 'た', 'タ':
			rStr = "ta"
			goto RomajiString
		case 'ち', 'チ':
			rStr = "chi"
			goto RomajiString
		case 'つ', 'ツ':
			rStr = "tsu"
			goto RomajiString
		case 'て', 'テ':
			rStr = "te"
			goto RomajiString
		case 'と', 'ト':
			rStr = "to"
			goto RomajiString
		// t dakuten (d)
		case 'だ', 'ダ':
			rStr = "da"
			goto RomajiString
		case 'ぢ', 'ヂ':
			rStr = "ji"
			goto RomajiString
		case 'づ', 'ヅ':
			rStr = "zu"
			goto RomajiString
		case 'で', 'デ':
			rStr = "de"
			goto RomajiString
		case 'ど', 'ド':
			rStr = "do"
			goto RomajiString
		// n
		case 'な', 'ナ':
			rStr = "na"
			goto RomajiString
		case 'に', 'ニ':
			rStr = "ni"
			goto RomajiString
		case 'ぬ', 'ヌ':
			rStr = "nu"
			goto RomajiString
		case 'ね', 'ネ':
			rStr = "ne"
			goto RomajiString
		case 'の', 'ノ':
			rStr = "no"
			goto RomajiString
		// h
		case 'は', 'ハ':
			rStr = "ha"
			goto RomajiString
		case 'ひ', 'ヒ':
			rStr = "hi"
			goto RomajiString
		case 'ふ', 'フ':
			rStr = "fu"
			goto RomajiString
		case 'へ', 'ヘ':
			rStr = "he"
			goto RomajiString
		case 'ほ', 'ホ':
			rStr = "ho"
			goto RomajiString
		// h dakuten (b)
		case 'ば', 'バ':
			rStr = "ba"
			goto RomajiString
		case 'び', 'ビ':
			rStr = "bi"
			goto RomajiString
		case 'ぶ', 'ブ':
			rStr = "bu"
			goto RomajiString
		case 'べ', 'ベ':
			rStr = "be"
			goto RomajiString
		case 'ぼ', 'ボ':
			rStr = "bo"
			goto RomajiString
		// h dakuten (p)
		case 'ぱ', 'パ':
			rStr = "pa"
			goto RomajiString
		case 'ぴ', 'ピ':
			rStr = "pi"
			goto RomajiString
		case 'ぷ', 'プ':
			rStr = "pu"
			goto RomajiString
		case 'ぺ', 'ペ':
			rStr = "pe"
			goto RomajiString
		case 'ぽ', 'ポ':
			rStr = "po"
			goto RomajiString
		// m
		case 'ま', 'マ':
			rStr = "ma"
			goto RomajiString
		case 'み', 'ミ':
			rStr = "mi"
			goto RomajiString
		case 'む', 'ム':
			rStr = "mu"
			goto RomajiString
		case 'め', 'メ':
			rStr = "me"
			goto RomajiString
		case 'も', 'モ':
			rStr = "mo"
			goto RomajiString
		// y
		case 'や', 'ヤ':
			rStr = "ya"
			goto RomajiString
		case 'ゆ', 'ユ':
			rStr = "yu"
			goto RomajiString
		case 'よ', 'ヨ':
			rStr = "yo"
			goto RomajiString
		// r
		case 'ら', 'ラ':
			rStr = "ra"
			goto RomajiString
		case 'り', 'リ':
			rStr = "ri"
			goto RomajiString
		case 'る', 'ル':
			rStr = "ru"
			goto RomajiString
		case 'れ', 'レ':
			rStr = "re"
			goto RomajiString
		case 'ろ', 'ロ':
			rStr = "ro"
			goto RomajiString
		// w
		case 'わ', 'ワ':
			rStr = "wa"
			goto RomajiString
		case 'ゐ', 'ヰ':
			rStr = "wi"
			goto RomajiString
		case 'ゑ', 'ヱ':
			rStr = "we"
			goto RomajiString
		case 'を', 'ヲ':
			rStr = "wo"
			goto RomajiString
		}

	RomajiByte:
		if len(rPrev) != 0 {
			sb.WriteString(rPrev)
		}

		sb.WriteByte(rByte)
		continue
	RomajiString:
		if len(rPrev) != 0 {
			sb.WriteString(rPrev)
		}

		rPrev = rStr
		continue
	}

	if len(rPrev) != 0 {
		sb.WriteString(rPrev)
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
