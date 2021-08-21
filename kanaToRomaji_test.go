package kanaconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	const want = "aiueon"

	for _, v := range [2]string{"あいうえおん", "アイウエオン"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestBasicDakuten(t *testing.T) {
	const want = "vu"

	for _, v := range [2]string{"ゔ", "ヴ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestK(t *testing.T) {
	const want = "kakikukeko"

	for _, v := range [2]string{"かきくけこ", "カキクケコ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestKDakuten(t *testing.T) {
	const want = "gagigugego"

	for _, v := range [2]string{"がぎぐげご", "ガギグゲゴ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestS(t *testing.T) {
	const want = "sashisuseso"

	for _, v := range [2]string{"さしすせそ", "サシスセソ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestSDakuten(t *testing.T) {
	const want = "zajizuzezo"

	for _, v := range [2]string{"ざじずぜぞ", "ザジズゼゾ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestT(t *testing.T) {
	const want = "tachitsuteto"

	for _, v := range [2]string{"たちつてと", "タチツテト"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestTDakuten(t *testing.T) {
	const want = "dajizudedo"

	for _, v := range [2]string{"だじづでど", "ダジヅデド"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestN(t *testing.T) {
	const want = "naninuneno"

	for _, v := range [2]string{"なにぬねの", "ナニヌネノ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestH(t *testing.T) {
	const want = "hahifuheho"

	for _, v := range [2]string{"はひふへほ", "ハヒフヘホ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestHDakutenB(t *testing.T) {
	const want = "babibubebo"

	for _, v := range [2]string{"ばびぶべぼ", "バビブベボ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestHDakutenP(t *testing.T) {
	const want = "papipupepo"

	for _, v := range [2]string{"ぱぴぷぺぽ", "パピプペポ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestM(t *testing.T) {
	const want = "mamimumemo"

	for _, v := range [2]string{"まみむめも", "マミムメモ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestY(t *testing.T) {
	const want = "yayuyo"

	for _, v := range [2]string{"やゆよ", "ヤユヨ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestR(t *testing.T) {
	const want = "rarirurero"

	for _, v := range [2]string{"らりるれろ", "ラリルレロ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, want, res)
	}
}

func TestW(t *testing.T) {
	const want = "wawiwewo"

	for _, v := range [2]string{"わゐゑを", "ワヰヱヲ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestYouon(t *testing.T) {
	const want = "youon cannot be the first character in a kana block"

	input := []string{
		"ゃき", "ゅき", "ょき", "ぁき", "ぃき", "ぅき", "ぇき", "ぉき", "ゎき",
		"ャキ", "ュキ", "ョキ", "ァキ", "ィキ", "ゥキ", "ェキ", "ォキ", "ヮキ",
	}

	for _, v := range input {
		_, got := KanaToRomaji(v)
		assert.EqualError(t, got, want)
	}
}

func TestYouonI(t *testing.T) {
	const want = "ye"
}

// func TestYouonK(t *testing.T) {
// 	const want = "kyakyukyokyakyikyukyekyo"

// 	for _, v := range [2]string{"きゃきゅきょきぁきぃきぅきぇきぉ", "キャキュキョキァキィキゥキェキォ"} {
// 		res, _ := KanaToRomaji(v)
// 		assert.Equal(t, want, res)
// 	}
// }

// func TestYouonKSpecial(t *testing.T) {
// 	const want = "kakikukekokwaki"

// 	for _, v := range [2]string{"くぁくぃくぅくぇくぉくゎけぃ", "クァクィクゥクェクォクヮケィ"} {
// 		res, _ := KanaToRomaji(v)
// 		assert.Equal(t, want, res)
// 	}
// }
