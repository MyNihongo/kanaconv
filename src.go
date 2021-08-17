package kanaconv

import (
	"errors"
	"strings"
)

func KanaToRomaji(str string) (string, error) {
	len := len(str)
	if len == 0 {
		return "", nil
	}

	var bld strings.Builder
	bld.Grow(len * 2)

	for i := 0; i < len; i++ {
		bld.WriteString("aa")
	}

	return bld.String(), errors.New("aaaa")
}
