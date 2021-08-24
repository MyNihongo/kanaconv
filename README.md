Conversion between Japanese characters

## Kana to romaji
```go
package main

import "github.com/MyNihongo/kanaconv"

func main() {
	res, err := kanaconv.KanaToRomaji("ひらがな") // hiragana
	res, err = kanaconv.KanaToRomaji("カタカナ") // katakana
	res, err = kanaconv.KanaToRomaji("ひらがな・カタカナ") // hiraganakatakana
}
```
