package kanaconv

type youon int8

const (
	youonYa youon = iota
	youonYu
	youonYo
	youonA
	youonI
	youonU
	youonE
	youonO
	youonWa
)

func (youonType youon) char() string {
	switch youonType {
	case youonYa, youonA:
		return "a"
	case youonYu, youonU:
		return "u"
	case youonYo, youonO:
		return "o"
	case youonI:
		return "i"
	case youonE:
		return "e"
	case youonWa:
		return "wa"
	}

	panic("unsupported youon character passed")
}
