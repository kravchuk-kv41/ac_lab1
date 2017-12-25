package lzw

import (
	"strconv"
	"strings"
)


func Compress(dict map[string]int, inStream string) []int {
	var tmp string
	var result []int
	for _, symbol := range inStream {
		tmp, result = compressMainLogic(dict, tmp, symbol, result)
	}
	return append(result, dict[tmp])
}


func compressMainLogic(dict map[string]int, tmp string, symbol int32, result []int) (string, []int) {
	if _, keyExists := dict[tmp+string(symbol)]; keyExists {
		tmp += string(symbol)
	} else {
		dict[tmp + string(symbol)] = len(dict)
		result = append(result, dict[tmp])
		tmp = string(symbol)
	}
	return tmp, result
}


func Decompress(dict map[string]int, inStream []int) string {
	var oldCode = inStream[0]
	var result = getKeyByValue(oldCode, dict)
	for _, code := range inStream[1:] {
		key := getKeyByValue(code, dict)
		if key != "" {result += key} else {
			key = getKeyByValue(oldCode, dict); result += key + string(key[0])
		}
		dict[getKeyByValue(oldCode, dict) + string(key[0])] = len(dict); oldCode = code
	}
	return result
}


func StringFromIntSlice(stream []int) string  {
	var result string

	for _, i := range stream {
		result += strconv.Itoa(i) + " "
	}
	return result
}


func SplitStringToSlice(str string) []int {
	slice := strings.Fields(str)
	var res []int
	for _, j := range slice {
		i, _ := strconv.Atoi(j)
		res = append(res, i)
	}
	return res
}


func GenerateDict() map[string]int {
	dict := map[string]int {}
	for i := 1; i < 129; i++ {
		dict[string(i)] = i
	}
	return dict
}


func getKeyByValue(value int, dict map[string]int) string {
	for k, v := range dict {
		if value == v {
			return k
		}
	}
	return ""
}

