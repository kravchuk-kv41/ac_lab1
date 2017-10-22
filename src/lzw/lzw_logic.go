package lzw

import (
	"strconv"
	"strings"
	"io/ioutil"
	"os"
)


func Compress(pathToInFile string, pathToOutFile string) {
	var strSlice []string
	for _, item := range compress(initMapWithLettersAndCodes(1, 129), readFile(pathToInFile)) {
		strSlice = append(strSlice, strconv.Itoa(item))
	}
	writeFile(strings.Join(strSlice, " "), pathToOutFile)
}


func Decompress(pathToInFile string, pathToOutFile string) {
	var res []int
	for _, j := range strings.Split(readFile(pathToInFile), " ") {
		i, _ := strconv.Atoi(j)
		res = append(res, i)
	}
	writeFile(decompress(initMapWithLettersAndCodes(1, 129), res), pathToOutFile)
}

// ---------------------------------------------------------------------------------------------------------------------

func compress(lzwMap map[string]int, uncompressedString string) []int {
	var (tmp string; result []int)
	for _, symbol := range uncompressedString {
		if _, keyExists := lzwMap[tmp + string(symbol)]; keyExists {
			tmp += string(symbol)
		} else {
			lzwMap[tmp + string(symbol)] = len(lzwMap)
			result = append(result, lzwMap[tmp])
			tmp = string(symbol)
		}
	}
	return append(result, lzwMap[tmp])
}


func decompress(lzwMap map[string]int, compressedString []int) string {
	var (beforeCode int = compressedString[0]; result string = getKeyFromMap(beforeCode, lzwMap))
	for _, nextCode := range compressedString[1:] {
		key := getKeyFromMap(nextCode, lzwMap)
		key, result = decompressUpdateResult(key, result, beforeCode, lzwMap)
		lzwMap[getKeyFromMap(beforeCode, lzwMap) + string(key[0])] = len(lzwMap)
		beforeCode = nextCode
	}
	return result
}


func decompressUpdateResult(key string, result string, beforeCode int, lzwMap map[string]int) (string, string) {
	if key != "" {
		result += key
	} else {
		key = getKeyFromMap(beforeCode, lzwMap)
		result += key + string(key[0])
	}
	return key, result
}


func initMapWithLettersAndCodes(from int, to int) map[string]int {
	lzwMap := map[string]int {}
	for i := from; i < to; i++ {
		lzwMap[string(i)] = i
	}
	return lzwMap
}


func getKeyFromMap(code int, lzwMap map[string]int) string {
	for key, value := range lzwMap {
		if value == code {
			return key
		}
	}
	return ""
}

// ---------------------------------------------------------------------------------------------------------------------

func readFile(filePath string) string {
	readString, err := ioutil.ReadFile(filePath)
	if err != nil {
		os.Exit(1)
	}
	return string(readString)
}


func writeFile(values string, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	file.WriteString(values)
	defer file.Close()
	file.Sync()
	return nil
}