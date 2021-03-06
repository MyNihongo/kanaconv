package kanaconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inp struct {
	input string
	want  string
}

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
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
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
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonKSpecial(t *testing.T) {
	const want = "kakikukekokwaki"

	for _, v := range [2]string{"くぁくぃくぅくぇくぉくゎけぃ", "クァクィクゥクェクォクヮケィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonG(t *testing.T) {
	const want = "gyagyugyogyagyigyugyegyo"

	for _, v := range [2]string{"ぎゃぎゅぎょぎぁぎぃぎぅぎぇぎぉ", "ギャギュギョギァギィギゥギェギォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonGSpecial(t *testing.T) {
	const want = "gagigugegogwagi"

	for _, v := range [2]string{"ぐぁぐぃぐぅぐぇぐぉぐゎげぃ", "グァグィグゥグェグォグヮゲィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonS(t *testing.T) {
	const want = "shashushoshashishushesho"

	for _, v := range [2]string{"しゃしゅしょしぁしぃしぅしぇしぉ", "シャシュショシァシィシゥシェシォ"} {
		got, _ := KanaToRomaji(v)
		assert.Equal(t, want, got)
	}
}

func TestYouonSSpecial(t *testing.T) {
	const want = "sasisusesoswasi"

	for _, v := range [2]string{"すぁすぃすぅすぇすぉすゎせぃ", "スァスィスゥスェスォスヮセィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonZ(t *testing.T) {
	const want = "jajujojajijujejo"

	for _, v := range [2]string{"じゃじゅじょじぁじぃじぅじぇじぉ", "ジャジュジョジァジィジゥジェジォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonZSpecial(t *testing.T) {
	const want = "zazizuzezozwazi"

	for _, v := range [2]string{"ずぁずぃずぅずぇずぉずゎぜぃ", "ズァズィズゥズェズォズヮゼィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonT(t *testing.T) {
	const want = "chachuchochachichuchecho"

	for _, v := range [2]string{"ちゃちゅちょちぁちぃちぅちぇちぉ", "チャチュチョチァチィチゥチェチォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonTSpecial(t *testing.T) {
	const want = "tsatsitsutsetsotswatityutu"

	for _, v := range [2]string{"つぁつぃつぅつぇつぉつゎてぃてゅとぅ", "ツァツィツゥツェツォツヮティテュトゥ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonDSpecial(t *testing.T) {
	const want = "zazizuzezozwadidyudu"

	for _, v := range [2]string{"づぁづぃづぅづぇづぉづゎでぃでゅどぅ", "ヅァヅィヅゥヅェヅォヅヮディデュドゥ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonN(t *testing.T) {
	const want = "nyanyunyonyanyinyunyenyo"

	for _, v := range [2]string{"にゃにゅにょにぁにぃにぅにぇにぉ", "ニャニュニョニァニィニゥニェニォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonNSpecial(t *testing.T) {
	const want = "naninunenonwani"

	for _, v := range [2]string{"ぬぁぬぃぬぅぬぇぬぉぬゎねぃ", "ヌァヌィヌゥヌェヌォヌヮネィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonH(t *testing.T) {
	const want = "hyahyuhyohyahyihyuhyehyo"

	for _, v := range [2]string{"ひゃひゅひょひぁひぃひぅひぇひぉ", "ヒャヒュヒョヒァヒィヒゥヒェヒォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonHSpecial(t *testing.T) {
	const want = "fafifufefofwafyuhi"

	for _, v := range [2]string{"ふぁふぃふぅふぇふぉふゎふゅへぃ", "ファフィフゥフェフォフヮフュヘィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonB(t *testing.T) {
	const want = "byabyubyobyabyibyubyebyo"

	for _, v := range [2]string{"びゃびゅびょびぁびぃびぅびぇびぉ", "ビャビュビョビァビィビゥビェビォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonBSpecial(t *testing.T) {
	const want = "babibubebobwabyubi"

	for _, v := range [2]string{"ぶぁぶぃぶぅぶぇぶぉぶゎぶゅべぃ", "ブァブィブゥブェブォブヮブュベィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonP(t *testing.T) {
	const want = "pyapyupyopyapyipyupyepyo"

	for _, v := range [2]string{"ぴゃぴゅぴょぴぁぴぃぴぅぴぇぴぉ", "ピャピュピョピァピィピゥピェピォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonPSpecial(t *testing.T) {
	const want = "papipupepopwapyupi"

	for _, v := range [2]string{"ぷぁぷぃぷぅぷぇぷぉぷゎぷゅぺぃ", "プァプィプゥプェプォプヮプュペィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonM(t *testing.T) {
	const want = "myamyumyomyamyimyumyemyo"

	for _, v := range [2]string{"みゃみゅみょみぁみぃみぅみぇみぉ", "ミャミュミョミァミィミゥミェミォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonMSpecial(t *testing.T) {
	const want = "mamimumemomwamyumi"

	for _, v := range [2]string{"むぁむぃむぅむぇむぉむゎむゅめぃ", "ムァムィムゥムェムォムヮムュメィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonYSpecial(t *testing.T) {
	const want = "yayiyuyeyoywa"

	for _, v := range [2]string{"ゆぁゆぃゆぅゆぇゆぉゆゎ", "ユァユィユゥユェユォユヮ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonR(t *testing.T) {
	const want = "ryaryuryoryaryiryuryeryo"

	for _, v := range [2]string{"りゃりゅりょりぁりぃりぅりぇりぉ", "リャリュリョリァリィリゥリェリォ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonRSpecial(t *testing.T) {
	const want = "rarirurerorwaryuri"

	for _, v := range [2]string{"るぁるぃるぅるぇるぉるゎるゅれぃ", "ルァルィルゥルェルォルヮルュレィ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestYouonWSpecial(t *testing.T) {
	const want = "wa"

	for _, v := range [2]string{"わぁ", "ワァ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestPunctuationMark(t *testing.T) {
	const want = "jiendo"
	got, err := KanaToRomaji("ジ・エンド")

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestSokuonInvalid(t *testing.T) {
	const want = "sokuon cannot precede a vowel"

	input := []string{
		"っあ", "っい", "っう", "っえ", "っお",
		"ッア", "ッイ", "ッウ", "ッエ", "ッオ",
	}

	for _, v := range input {
		got, err := KanaToRomaji(v)
		assert.EqualError(t, err, want)
		assert.Empty(t, got)
	}
}

func TestSokuon(t *testing.T) {
	const want = "nn"

	for _, v := range []string{"っん", "ッン"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonK(t *testing.T) {
	const want = "kkakkikkukkekko"

	for _, v := range []string{"っかっきっくっけっこ", "ッカッキックッケッコ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonG(t *testing.T) {
	const want = "ggaggigguggeggo"

	for _, v := range []string{"っがっぎっぐっげっご", "ッガッギッグッゲッゴ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonS(t *testing.T) {
	const want = "ssasshissussesso"

	for _, v := range []string{"っさっしっすっせっそ", "ッサッシッスッセッソ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonZ(t *testing.T) {
	const want = "zzajjizzuzzezzo"

	for _, v := range []string{"っざっじっずっぜっぞ", "ッザッジッズッゼッゾ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonT(t *testing.T) {
	const want = "ttatchittsuttetto"

	for _, v := range []string{"ったっちっつってっと", "ッタッチッツッテット"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonD(t *testing.T) {
	const want = "ddajjizzuddeddo"

	for _, v := range []string{"っだっぢっづっでっど", "ッダッヂッヅッデッド"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonN(t *testing.T) {
	const want = "nnanninnunnenno"

	for _, v := range []string{"っなっにっぬっねっの", "ッナッニッヌッネッノ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonH(t *testing.T) {
	const want = "hhahhiffuhhehho"

	for _, v := range []string{"っはっひっふっへっほ", "ッハッヒッフッヘッホ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonB(t *testing.T) {
	const want = "bbabbibbubbebbo"

	for _, v := range []string{"っばっびっぶっべっぼ", "ッバッビッブッベッボ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonP(t *testing.T) {
	const want = "ppappippuppeppo"

	for _, v := range []string{"っぱっぴっぷっぺっぽ", "ッパッピップッペッポ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonM(t *testing.T) {
	const want = "mmammimmummemmo"

	for _, v := range []string{"っまっみっむっめっも", "ッマッミッムッメッモ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonY(t *testing.T) {
	const want = "yyayyuyyo"

	for _, v := range []string{"っやっゆっよ", "ッヤッユッヨ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonR(t *testing.T) {
	const want = "rrarrirrurrerro"

	for _, v := range []string{"っらっりっるっれっろ", "ッラッリッルッレッロ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestSokuonW(t *testing.T) {
	const want = "wwawwo"

	for _, v := range []string{"っわっを", "ッワッヲ"} {
		got, err := KanaToRomaji(v)
		assert.Equal(t, want, got)
		assert.Nil(t, err)
	}
}

func TestChouonpuA(t *testing.T) {
	input := []inp{
		{input: "アー", want: "aa"},
		{input: "カーカ", want: "kaaka"},
		{input: "キャーカ", want: "kyaaka"},
		{input: "ガーカ", want: "gaaka"},
		{input: "ギャーカ", want: "gyaaka"},
		{input: "サーカ", want: "saaka"},
		{input: "シャーカ", want: "shaaka"},
		{input: "ザーカ", want: "zaaka"},
		{input: "ジャーカ", want: "jaaka"},
		{input: "ターカ", want: "taaka"},
		{input: "チャーカ", want: "chaaka"},
		{input: "ダーカ", want: "daaka"},
		{input: "ヂャーカ", want: "jaaka"},
		{input: "ナーカ", want: "naaka"},
		{input: "ニャーカ", want: "nyaaka"},
		{input: "ハーカ", want: "haaka"},
		{input: "ヒャーカ", want: "hyaaka"},
		{input: "バーカ", want: "baaka"},
		{input: "ビャーカ", want: "byaaka"},
		{input: "パーカ", want: "paaka"},
		{input: "ピャーカ", want: "pyaaka"},
		{input: "マーカ", want: "maaka"},
		{input: "ミャーカ", want: "myaaka"},
		{input: "ヤーカ", want: "yaaka"},
		{input: "ラーカ", want: "raaka"},
		{input: "リャーカ", want: "ryaaka"},
		{input: "ワーカ", want: "waaka"},
	}

	for _, v := range input {
		got, err := KanaToRomaji(v.input)
		assert.Equal(t, v.want, got)
		assert.Nil(t, err)
	}
}

func TestChouonpuI(t *testing.T) {
	input := []inp{
		{input: "イー", want: "ii"},
		{input: "キーカ", want: "kiika"},
		{input: "キィーカ", want: "kyiika"},
		{input: "ギーカ", want: "giika"},
		{input: "ギィーカ", want: "gyiika"},
		{input: "シーカ", want: "shiika"},
		{input: "シィーカ", want: "shiika"},
		{input: "ジーカ", want: "jiika"},
		{input: "ジィーカ", want: "jiika"},
		{input: "チーカ", want: "chiika"},
		{input: "チィーカ", want: "chiika"},
		{input: "ヂーカ", want: "jiika"},
		{input: "ヂィーカ", want: "jiika"},
		{input: "ニーカ", want: "niika"},
		{input: "ニィーカ", want: "nyiika"},
		{input: "ヒーカ", want: "hiika"},
		{input: "ヒィーカ", want: "hyiika"},
		{input: "ビーカ", want: "biika"},
		{input: "ビィーカ", want: "byiika"},
		{input: "ピーカ", want: "piika"},
		{input: "ピィーカ", want: "pyiika"},
		{input: "ミーカ", want: "miika"},
		{input: "ミィーカ", want: "myiika"},
		{input: "リーカ", want: "riika"},
		{input: "リィーカ", want: "ryiika"},
		{input: "ヰーカ", want: "wiika"},
	}

	for _, v := range input {
		got, err := KanaToRomaji(v.input)
		assert.Equal(t, v.want, got)
		assert.Nil(t, err)
	}
}

func TestChouonpuU(t *testing.T) {
	input := []inp{
		{input: "ウー", want: "uu"},
		{input: "クーカ", want: "kuuka"},
		{input: "キューカ", want: "kyuuka"},
		{input: "グーカ", want: "guuka"},
		{input: "ギューカ", want: "gyuuka"},
		{input: "スーカ", want: "suuka"},
		{input: "シューカ", want: "shuuka"},
		{input: "ズーカ", want: "zuuka"},
		{input: "ジューカ", want: "juuka"},
		{input: "ツーカ", want: "tsuuka"},
		{input: "チューカ", want: "chuuka"},
		{input: "ヅーカ", want: "zuuka"},
		{input: "ヂューカ", want: "juuka"},
		{input: "ヌーカ", want: "nuuka"},
		{input: "ニューカ", want: "nyuuka"},
		{input: "フーカ", want: "fuuka"},
		{input: "ヒューカ", want: "hyuuka"},
		{input: "ブーカ", want: "buuka"},
		{input: "ビューカ", want: "byuuka"},
		{input: "プーカ", want: "puuka"},
		{input: "ピューカ", want: "pyuuka"},
		{input: "ムーカ", want: "muuka"},
		{input: "ミューカ", want: "myuuka"},
		{input: "ユーカ", want: "yuuka"},
		{input: "ルーカ", want: "ruuka"},
		{input: "リューカ", want: "ryuuka"},
	}

	for _, v := range input {
		got, err := KanaToRomaji(v.input)
		assert.Equal(t, v.want, got)
		assert.Nil(t, err)
	}
}

func TestChouonpuE(t *testing.T) {
	input := []inp{
		{input: "エー", want: "ee"},
		{input: "ケーカ", want: "keeka"},
		{input: "キェーカ", want: "kyeeka"},
		{input: "ゲーカ", want: "geeka"},
		{input: "ギェーカ", want: "gyeeka"},
		{input: "セーカ", want: "seeka"},
		{input: "シェーカ", want: "sheeka"},
		{input: "ゼーカ", want: "zeeka"},
		{input: "ジェーカ", want: "jeeka"},
		{input: "テーカ", want: "teeka"},
		{input: "チェーカ", want: "cheeka"},
		{input: "デーカ", want: "deeka"},
		{input: "ヂェーカ", want: "jeeka"},
		{input: "ネーカ", want: "neeka"},
		{input: "ニェーカ", want: "nyeeka"},
		{input: "ヘーカ", want: "heeka"},
		{input: "ヒェーカ", want: "hyeeka"},
		{input: "ベーカ", want: "beeka"},
		{input: "ビェーカ", want: "byeeka"},
		{input: "ペーカ", want: "peeka"},
		{input: "ピェーカ", want: "pyeeka"},
		{input: "メーカ", want: "meeka"},
		{input: "ミェーカ", want: "myeeka"},
		{input: "レーカ", want: "reeka"},
		{input: "リェーカ", want: "ryeeka"},
		{input: "ヱーカ", want: "weeka"},
	}

	for _, v := range input {
		got, err := KanaToRomaji(v.input)
		assert.Equal(t, v.want, got)
		assert.Nil(t, err)
	}
}

func TestChouonpuO(t *testing.T) {
	input := []inp{
		{input: "オー", want: "oo"},
		{input: "コーカ", want: "kooka"},
		{input: "キョーカ", want: "kyooka"},
		{input: "ゴーカ", want: "gooka"},
		{input: "ギョーカ", want: "gyooka"},
		{input: "ソーカ", want: "sooka"},
		{input: "ショーカ", want: "shooka"},
		{input: "ゾーカ", want: "zooka"},
		{input: "ジョーカ", want: "jooka"},
		{input: "トーカ", want: "tooka"},
		{input: "チョーカ", want: "chooka"},
		{input: "ドーカ", want: "dooka"},
		{input: "ヂョーカ", want: "jooka"},
		{input: "ノーカ", want: "nooka"},
		{input: "ニョーカ", want: "nyooka"},
		{input: "ホーカ", want: "hooka"},
		{input: "ヒョーカ", want: "hyooka"},
		{input: "ボーカ", want: "booka"},
		{input: "ビョーカ", want: "byooka"},
		{input: "ポーカ", want: "pooka"},
		{input: "ピョーカ", want: "pyooka"},
		{input: "モーカ", want: "mooka"},
		{input: "ミョーカ", want: "myooka"},
		{input: "ローカ", want: "rooka"},
		{input: "リョーカ", want: "ryooka"},
		{input: "ヲーカ", want: "wooka"},
	}

	for _, v := range input {
		got, err := KanaToRomaji(v.input)
		assert.Equal(t, v.want, got)
		assert.Nil(t, err)
	}
}

func TestNonKanaCharacters(t *testing.T) {
	inputs := []string{
		"日本",
		"English",
	}

	for _, input := range inputs {
		got, err := KanaToRomaji(input)

		assert.Empty(t, got)
		assert.NotNil(t, err)
	}
}
