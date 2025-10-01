package utils

import (
	"html"
	"math/rand"
	"net/url"
	"regexp"
	"strconv"
)

// StrToUInt 字符串转 uint
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
}

func StrToInt(str string) int {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return i
}

// RemoveRepeatedElement 数组去重
// func RemoveRepeatedElement(arr []string) (newArr []string) {
// 	newArr = make([]string, 0)
// 	for i := 0; i < len(arr); i++ {
// 		repeat := false
// 		for j := i + 1; j < len(arr); j++ {
// 			if arr[i] == arr[j] {
// 				repeat = true
// 				break
// 			}
// 		}
// 		if !repeat {
// 			newArr = append(newArr, arr[i])
// 		}
// 	}
// 	return
// }

// RemoveRepeated 数组去重
func RemoveRepeated(s []string) []string {
	result := make([]string, 0)
	temp := map[string]interface{}{}
	for _, v := range s {
		if _, ok := temp[v]; !ok {
			temp[v] = nil
			result = append(result, v)
		}
	}
	return result
}

// func AppendAndDistinct(s1 []string, s2 []string) []string {
// 	s := append(s1, s2...)
// 	result := make([]string, 0)
// 	temp := map[string]interface{}{}
// 	for _, v := range s {
// 		if _, ok := temp[v]; !ok {
// 			temp[v] = nil
// 			result = append(result, v)
// 		}
// 	}
// 	return result
// }

// 判断字符串是不是时间戳
func IsTimestamp(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}

// 简单过滤XXS攻击和sql注入
var sqlInjectionRegex = regexp.MustCompile(`(?i)('|--|\|)|((\%27)|(\-\-)|(\%7C))`)

func SanitizeInput(input string) string {
	// Escape the input to prevent XSS attacks
	safeHTML := html.EscapeString(input)

	// URL encode the input to prevent SQL injection attacks
	safeSQL := url.QueryEscape(safeHTML)

	// Further sanitize the input to remove special characters
	// that might lead to SQL injection attacks
	safeSQL = sqlInjectionRegex.ReplaceAllString(safeSQL, "")

	return safeSQL
}

// 随机数
func StringWithCharset(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// func RandString(n int) string {
// 	b := make([]byte, n)
// 	_, _ = rand.Read(b)
// 	return base64.RawURLEncoding.EncodeToString(b)[:n]
// }
