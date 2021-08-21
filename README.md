Conversion between Japanese characters

## Kana to romaji
```go
package main

import "github.com/MyNihongo/kanaconv"

func main() {
	res := kanaconv.KanaToRomaji("ひらがな") // hiragana
	res = kanaconv.KanaToRomaji("カタカナ") // katakana
	res = kanaconv.KanaToRomaji("ひらがな・カタカナ") // hiraganakatakana
}
```
