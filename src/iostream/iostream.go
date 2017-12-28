package iostream

import (
	"io/ioutil"
	"os"
)


func ReadFile(filePath string) (string, error) {
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(b), nil
}


func WriteFile(filePath string, values string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	file.WriteString(values)
	defer file.Close()
	file.Sync()
	//file.Close()
	return nil
}


func WriteByteSliceToFile(file *os.File, byteSlice []int) {
	b := []byte{0}
	for _, it := range byteSlice {
		tmp := it
		for ; tmp != 0; {
			if tmp > 127 {
				b[0] = uint8(tmp & 127 | 128)
			} else {
				b[0] = uint8(tmp)
			}
			file.Write(b)
			tmp = tmp >> 7
		}
	}
	file.Write([]byte{0})
}


func ReadByteSliceFromFile(pathToFile string) [][]int {
	buffer, _ := ioutil.ReadFile(pathToFile)
	currentInt := 0
	currentBitCount := uint8(0)
	var slice []int
	var slices [][]int

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


func ReadUncompressedFile(pathToFile string) []string {
	buffer, err := ioutil.ReadFile(pathToFile)
	if err != nil || len(buffer) == 0 {
		panic("Error acquired during reading file" + pathToFile)
	}
	return separateInSeveralStreamsByStringSize(buffer, 2000000)
}

func separateInSeveralStreamsByStringSize(buffer []byte, maxStringSize int) []string {
	var chunks []string
	for offset := 0; offset < len(buffer); offset += maxStringSize {
		if len(buffer)-offset < maxStringSize {
			chunks = append(chunks, string(buffer[offset:]))
		} else {
			chunks = append(chunks, string(buffer[offset:(maxStringSize + offset)]))
		}
	}
	return chunks
}
