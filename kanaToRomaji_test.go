package kanaconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	const expectedRes = "aiueon"

	for _, v := range [2]string{"あいうえおん", "アイウエオン"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestBasicDakuten(t *testing.T) {
	const expectedRes = "vu"

	for _, v := range [2]string{"ゔ", "ヴ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestK(t *testing.T) {
	const expectedRes = "kakikukeko"

	for _, v := range [2]string{"かきくけこ", "カキクケコ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestKDakuten(t *testing.T) {
	const expectedRes = "gagigugego"

	for _, v := range [2]string{"がぎぐげご", "ガギグゲゴ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestS(t *testing.T) {
	const expectedRes = "sashisuseso"

	for _, v := range [2]string{"さしすせそ", "サシスセソ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestSDakuten(t *testing.T) {
	const expectedRes = "zajizuzezo"

	for _, v := range [2]string{"ざじずぜぞ", "ザジズゼゾ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestT(t *testing.T) {
	const expectedRes = "tachitsuteto"

	for _, v := range [2]string{"たちつてと", "タチツテト"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestTDakuten(t *testing.T) {
	const expectedRes = "dajizudedo"

	for _, v := range [2]string{"だじづでど", "ダジヅデド"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestN(t *testing.T) {
	const expectedRes = "naninuneno"

	for _, v := range [2]string{"なにぬねの", "ナニヌネノ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestH(t *testing.T) {
	const expectedRes = "hahifuheho"

	for _, v := range [2]string{"はひふへほ", "ハヒフヘホ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestHDakutenB(t *testing.T) {
	const expectedRes = "babibubebo"

	for _, v := range [2]string{"ばびぶべぼ", "バビブベボ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestHDakutenP(t *testing.T) {
	const expectedRes = "papipupepo"

	for _, v := range [2]string{"ぱぴぷぺぽ", "パピプペポ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestM(t *testing.T) {
	const expectedRes = "mamimumemo"

	for _, v := range [2]string{"まみむめも", "マミムメモ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestY(t *testing.T) {
	const expectedRes = "yayuyo"

	for _, v := range [2]string{"やゆよ", "ヤユヨ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestR(t *testing.T) {
	const expectedRes = "rarirurero"

	for _, v := range [2]string{"らりるれろ", "ラリルレロ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestW(t *testing.T) {
	const expectedRes = "wawiwewo"

	for _, v := range [2]string{"わゐゑを", "ワヰヱヲ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}

func TestYouon(t *testing.T) {
	const expectedRes = "youon cannot be the first character in a kana block"

	input := []string{
		"ゃき", "ゅき", "ょき", "ぁき", "ぃき", "ぅき", "ぇき", "ぉき", "ゎき",
		"ャキ", "ュキ", "ョキ", "ァキ", "ィキ", "ゥキ", "ェキ", "ォキ", "ヮキ",
	}

	for _, v := range input {
		_, res := KanaToRomaji(v)
		assert.EqualError(t, res, expectedRes)
	}
}

// func TestYouonK(t *testing.T) {
// 	const expectedRes = "kyakyukyokyakyikyukyekyo"

// 	for _, v := range [2]string{"きゃきゅきょきぁきぃきぅきぇきぉ", "キャキュキョキァキィキゥキェキォ"} {
// 		res, _ := KanaToRomaji(v)
// 		assert.Equal(t, expectedRes, res)
// 	}
// }

// func TestYouonKSpecial(t *testing.T) {
// 	const expectedRes = "kakikukekokwaki"

// 	for _, v := range [2]string{"くぁくぃくぅくぇくぉくゎけぃ", "クァクィクゥクェクォクヮケィ"} {
// 		res, _ := KanaToRomaji(v)
// 		assert.Equal(t, expectedRes, res)
// 	}
// }
