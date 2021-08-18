package kanaconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVowels(t *testing.T) {
	const expectedRes = "aiueo"

	for _, v := range [2]string{"あいうえお", "アイウエオ"} {
		res, _ := KanaToRomaji(v)
		assert.Equal(t, expectedRes, res)
	}
}
