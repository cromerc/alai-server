package utils

import (
	"bytes"
	"strconv"
)

func toInt(toConvert string) int {
	converted, _ := strconv.Atoi(toConvert)
	return converted
}

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func InsertNth(s string, n int, symbol rune) string {
	var buffer bytes.Buffer
	var n_1 = n - 1
	var l_1 = len(s) - 1
	for i, rune := range s {
		buffer.WriteRune(rune)
		if i%n == n_1 && i != l_1 {
			buffer.WriteRune(symbol)
		}
	}
	return buffer.String()
}
