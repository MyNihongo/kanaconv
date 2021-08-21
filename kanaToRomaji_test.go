package kanaconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	const want = "aiueon"

	for _, v := range [2]string{"あいうえおん", "アイウエオン"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestBasicDakuten(t *testing.T) {
	const want = "vu"

	for _, v := range [2]string{"ゔ", "ヴ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestK(t *testing.T) {
	const want = "kakikukeko"

	for _, v := range [2]string{"かきくけこ", "カキクケコ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestKDakuten(t *testing.T) {
	const want = "gagigugego"

	for _, v := range [2]string{"がぎぐげご", "ガギグゲゴ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestS(t *testing.T) {
	const want = "sashisuseso"

	for _, v := range [2]string{"さしすせそ", "サシスセソ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSDakuten(t *testing.T) {
	const want = "zajizuzezo"

	for _, v := range [2]string{"ざじずぜぞ", "ザジズゼゾ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestT(t *testing.T) {
	const want = "tachitsuteto"

	for _, v := range [2]string{"たちつてと", "タチツテト"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestTDakuten(t *testing.T) {
	const want = "dajizudedo"

	for _, v := range [2]string{"だじづでど", "ダジヅデド"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestN(t *testing.T) {
	const want = "naninuneno"

	for _, v := range [2]string{"なにぬねの", "ナニヌネノ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestH(t *testing.T) {
	const want = "hahifuheho"

	for _, v := range [2]string{"はひふへほ", "ハヒフヘホ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestHDakutenB(t *testing.T) {
	const want = "babibubebo"

	for _, v := range [2]string{"ばびぶべぼ", "バビブベボ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestHDakutenP(t *testing.T) {
	const want = "papipupepo"

	for _, v := range [2]string{"ぱぴぷぺぽ", "パピプペポ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestM(t *testing.T) {
	const want = "mamimumemo"

	for _, v := range [2]string{"まみむめも", "マミムメモ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestY(t *testing.T) {
	const want = "yayuyo"

	for _, v := range [2]string{"やゆよ", "ヤユヨ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestR(t *testing.T) {
	const want = "rarirurero"

	for _, v := range [2]string{"らりるれろ", "ラリルレロ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestW(t *testing.T) {
	const want = "wawiwewo"

	for _, v := range [2]string{"わゐゑを", "ワヰヱヲ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonInvalid(t *testing.T) {
	const want = "yōon cannot be the first character in a kana block"

	input := []string{
		"ゃき", "ゅき", "ょき", "ぁき", "ぃき", "ぅき", "ぇき", "ぉき", "ゎき",
		"ャキ", "ュキ", "ョキ", "ァキ", "ィキ", "ゥキ", "ェキ", "ォキ", "ヮキ",
	}

	for _, v := range input {
		got, err := KanaToRomaji(v)
		assert.Empty(t, got)
		assert.EqualError(t, err, want)
	}
}

func TestYouonSpecialInvalid(t *testing.T) {
	const want = "unrecognised yōon vowel"

	input := []string{
		"あぃ", "えぃ", "おぃ",
		"アィ", "エィ", "オィ",
	}

	for _, v := range input {
		got, err := KanaToRomaji(v)
		assert.Empty(t, got)
		assert.EqualError(t, err, want)
	}
}

func TestYouonI(t *testing.T) {
	const want = "ye"

	for _, v := range [2]string{"いぇ", "イェ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonU(t *testing.T) {
	const want = "wiwe"

	for _, v := range [2]string{"うぃうぇ", "ウィウェ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonV(t *testing.T) {
	const want = "vyavyuvyovavivuvevo"

	for _, v := range [2]string{"ゔゃゔゅゔょゔぁゔぃゔぅゔぇゔぉ", "ヴャヴュヴョヴァヴィヴゥヴェヴォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonK(t *testing.T) {
	const want = "kyakyukyokyakyikyukyekyo"

	for _, v := range [2]string{"きゃきゅきょきぁきぃきぅきぇきぉ", "キャキュキョキァキィキゥキェキォ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonKSpecial(t *testing.T) {
	const want = "kakikukekokwaki"

	for _, v := range [2]string{"くぁくぃくぅくぇくぉくゎけぃ", "クァクィクゥクェクォクヮケィ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonG(t *testing.T) {
	const want = "gyagyugyogyagyigyugyegyo"

	for _, v := range [2]string{"ぎゃぎゅぎょぎぁぎぃぎぅぎぇぎぉ", "ギャギュギョギァギィギゥギェギォ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonGSpecial(t *testing.T) {
	const want = "gagigugegogwagi"

	for _, v := range [2]string{"ぐぁぐぃぐぅぐぇぐぉぐゎげぃ", "グァグィグゥグェグォグヮゲィ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonS(t *testing.T) {
	const want = "shashushoshashishushesho"

	for _, v := range [2]string{"しゃしゅしょしぁしぃしぅしぇしぉ", "シャシュショシァシィシゥシェシォ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, want, res)
	}
}

func TestYouonSSpecial(t *testing.T) {
	const want = "sasisusesoswasi"

	for _, v := range [2]string{"すぁすぃすぅすぇすぉすゎせぃ", "スァスィスゥスェスォスヮセィ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonZ(t *testing.T) {
	const want = "jajujojajijujejo"

	for _, v := range [2]string{"じゃじゅじょじぁじぃじぅじぇじぉ", "ジャジュジョジァジィジゥジェジォ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonZSpecial(t *testing.T) {
	const want = "zazizuzezozwazi"

	for _, v := range [2]string{"ずぁずぃずぅずぇずぉずゎぜぃ", "ズァズィズゥズェズォズヮゼィ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonT(t *testing.T) {
	const want = "chachuchochachichuchecho"

	for _, v := range [2]string{"ちゃちゅちょちぁちぃちぅちぇちぉ", "チャチュチョチァチィチゥチェチォ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonTSpecial(t *testing.T) {
	const want = "tsatsitsutsetsotswatityutu"

	for _, v := range [2]string{"つぁつぃつぅつぇつぉつゎてぃてゅとぅ", "ツァツィツゥツェツォツヮティテュトゥ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonDSpecial(t *testing.T) {
	const want = "zazizuzezozwadidyudu"

	for _, v := range [2]string{"づぁづぃづぅづぇづぉづゎでぃでゅどぅ", "ヅァヅィヅゥヅェヅォヅヮディデュドゥ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonN(t *testing.T) {
	const want = "nyanyunyonyanyinyunyenyo"

	for _, v := range [2]string{"にゃにゅにょにぁにぃにぅにぇにぉ", "ニャニュニョニァニィニゥニェニォ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonNSpecial(t *testing.T) {
	const want = "naninunenonwani"

	for _, v := range [2]string{"ぬぁぬぃぬぅぬぇぬぉぬゎねぃ", "ヌァヌィヌゥヌェヌォヌヮネィ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonH(t *testing.T) {
	const want = "hyahyuhyohyahyihyuhyehyo"

	for _, v := range [2]string{"ひゃひゅひょひぁひぃひぅひぇひぉ", "ヒャヒュヒョヒァヒィヒゥヒェヒォ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}

func TestYouonHSpecial(t *testing.T) {
	const want = "fafifufefofwafyuhi"

	for _, v := range [2]string{"ふぁふぃふぅふぇふぉふゎふゅへぃ", "ファフィフゥフェフォフヮフュヘィ"} {
		res, err := KanaToRomaji(v)
		assert.Equal(t, want, res)
		assert.Nil(t, err)
	}
}
