package kata

import (
	"strings"
)

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func SpinWords(str string) string {
	sentence := strings.Fields(str)
	result := ""
	for _, word := range sentence{
		if len(word) >= 5 {
			result += Reverse(word) + " "
		} else {
			result += word + " "
		}
	}
	result = strings.TrimSpace(result)
	return result
}
