package utils

import "strconv"

// StrToUInt 字符串转 uint
func StrToUInt(str string) uint {
	i, e := strconv.Atoi(str)
	if e != nil {
		return 0
	}
	return uint(i)
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
