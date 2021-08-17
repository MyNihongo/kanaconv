package kanaconv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	const expectedRes = "123"

	res, _ := KanaToRomaji("aa")
	assert.Equal(t, expectedRes, res)
}
