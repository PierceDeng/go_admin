package utils

import (
	"strings"
	"unicode"
)

func UniqueStrings(s []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(s))
	for _, v := range s {
		if !seen[v] {
			seen[v] = true
			result = append(result, v)
		}
	}
	return result
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s) // 使用 rune 切片处理多字节字符
	// 将第一个字符转为大写
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

const (
	HTTP  = "http://"
	HTTPS = "https://"
)

// StartsWithAny 检查字符串 s 是否以 prefixes 中的任意一个字符串开头
func StartsWithAny(s string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(s, prefix) {
			return true
		}
	}
	return false
}
