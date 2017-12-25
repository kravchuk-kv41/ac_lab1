package lzw

import (
	"io/ioutil"
	"os"
)


func Compress(pathToInFile string, pathToOutFile string) {
	outFile, _:= os.OpenFile(pathToOutFile, os.O_CREATE|os.O_RDWR, 0666)
	defer outFile.Close()
	strings := readUncompressedFile(pathToInFile)
	dict := initMapWithLettersAndCodes(1, 129)
	println(len(strings))
	for _, str := range strings {
		writeByteSliceToFile(outFile, compress(dict, str))
	}
}


func Decompress(pathToInFile string, pathToOutFile string) {
	outFile, _:= os.OpenFile(pathToOutFile, os.O_CREATE|os.O_RDWR, 0666)
	defer outFile.Close()
	compressed := readByteSliceFromFile(pathToInFile)
	println(len(compressed))
	dict := reverseMap(initMapWithLettersAndCodes(1, 129))
	for i, byteSlice := range compressed {
		dec := decompress(dict, byteSlice)
		print("operation: "); println(i)
		outFile.WriteString(dec)
		outFile.Sync()
	}
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


func decompress(lzwMap map[int]string, compressed []int) string {
	dictSize := 128
	w := string(compressed[0])
	result := w
	for _, k := range compressed[1:] {
		var entry string
		if x, ok := lzwMap[k]; ok { entry = x } else if k == dictSize { entry = w + w[:1] }
		result += entry
		lzwMap[dictSize] = w + entry[:1]
		dictSize++
		w = entry
	}
	return result
}


func initMapWithLettersAndCodes(from int, to int) map[string]int {
	lzwMap := map[string]int {}
	for i := from; i < to; i++ {
		lzwMap[string(i)] = i
	}
	return lzwMap
}


func reverseMap(m map[string]int) map[int]string {
	n := make(map[int]string)
	for k, v := range m { n[v] = k }
	return n
}

// ---------------------------------------------------------------------------------------------------------------------

const (
	maxChunkSize = 2000000
)

func readUncompressedFile(pathToFile string) []string {
	buffer, _ := ioutil.ReadFile(pathToFile)
	return separateInSeveralStreamsByStringSize(buffer, maxChunkSize)
}

func separateInSeveralStreamsByStringSize(buffer []byte, maxStringSize int) []string {
	var chunks []string
	for offset := 0; offset < len(buffer); offset += maxStringSize {
		if len(buffer)-offset < maxStringSize {
			chunks = append(chunks, string(buffer[offset:]))
		} else {
			chunks = append(chunks, string(buffer[offset:(maxChunkSize + offset)]))
		}
	}
	return chunks
}

// ---------------------------------------------------------------------------------------------------------------------

func readByteSliceFromFile(pathToFile string) [][]int {
	buffer, _ := ioutil.ReadFile(pathToFile)
	var (
		slice []int
		slices [][]int
		currentInt int
		currentBitCount uint8 = 0
	)

	for _, currentByte := range buffer {
		if currentByte == 0 {
			slices = append(slices, slice)
			slice = []int{}
		} else {
			currentInt = currentInt | (int(currentByte & 127) << currentBitCount)
			if currentByte & 128 == 0 {
				slice = append(slice, currentInt)
				currentInt = 0
				currentBitCount = 0
			} else {
				currentBitCount += 7
			}
		}
	}
	return slices
}

func writeByteSliceToFile(file *os.File, byteSlice []int) {
	b := []byte{0}
	for _, it := range byteSlice {
		tmp := it
		for ; tmp != 0; tmp = tmp >> 7 {
			if tmp > 127 { b[0] = uint8(tmp & 127 | 128) } else { b[0] = uint8(tmp)}
			file.Write(b)
		}
	}
	file.Write([]byte{0})
	file.Sync()
}
