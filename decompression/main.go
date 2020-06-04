package decompression

import (
	"strconv"
)

// Decompress 2[3[a]b] -> aaabaaab
func Decompress(compressed string) string {
	tokenized := tokenize(compressed)
	return expand(tokenized)
}

func expand(target []string) string {
	decompressed := ""
	for i := 0; i < len(target); i++ {
		currentChar := target[i]
		// 普通の文字
		if !isInt(currentChar) && !isBracket(currentChar) {
			decompressed += currentChar
			continue
		}
		// 数値
		if isInt(currentChar) {
			repeatCount := toInt(currentChar)
			j := i + 2 // [の次の文字
			insideBracket := []string{}
			bracketStack := 1
			for ; ; j++ {
				if target[j] == "[" {
					bracketStack++
					insideBracket = append(insideBracket, target[j])
					continue
				}
				if target[j] == "]" {
					bracketStack--
					if bracketStack == 0 {
						break
					}
					insideBracket = append(insideBracket, target[j])
					continue
				}
				insideBracket = append(insideBracket, target[j])
			}
			decompressed += repeatString(repeatCount, expand(insideBracket))
			i = j
		}
	}
	return decompressed
}

// 3[abc]4[ab]c -> ["3", "[", "abc", "]", "4", "[", "ab", "]", "c"]
func tokenize(target string) []string {
	tokens := []string{}
	for _, r := range target {
		char := string(r)
		if len(tokens) == 0 {
			tokens = append(tokens, char)
			continue
		}
		// []の場合
		if char == "[" || char == "]" {
			tokens = append(tokens, char)
			continue
		}
		// 数値の場合
		last := tokens[len(tokens)-1]
		if isInt(char) {
			// tokensの最後の要素が数値の場合はそこに追加
			if isInt(last) {
				tokens[len(tokens)-1] = last + char
			} else {
				tokens = append(tokens, char)
			}
			continue
		}
		// 文字列の場合
		// // tokensの最後の要素が文字の場合はそこに追加
		if isInt(last) || last == "[" || last == "]" {
			tokens = append(tokens, char)
		} else {
			tokens[len(tokens)-1] = last + char
		}
	}
	return tokens
}

func isInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func toInt(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func isBracket(str string) bool {
	return str == "[" || str == "]"
}

func repeatString(count int, str string) string {
	result := ""
	for ; count > 0; count-- {
		result += str
	}
	return result
}
