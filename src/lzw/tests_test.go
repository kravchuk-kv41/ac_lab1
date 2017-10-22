package lzw

import (
	"strconv"
	"testing"
)


func TestGenerateDict(t *testing.T) {
	generated := initMapWithLettersAndCodes(1, 129)

	for i := 1; i < 129; i++ {
		if generated[string(i)] != i {
			t.Error("Expected " + strconv.Itoa(i), generated[string(i)])
		}
	}
}


func TestCompress(t *testing.T) {
	compressed := compress(initMapWithLettersAndCodes(1, 129), "1234567789")
	encoded := []int{49, 50, 51, 52, 53, 54, 55, 55, 56, 57}
	for i := range compressed {
		if compressed[i] != encoded[i] {
			t.Error("Expected " + strconv.Itoa(encoded[i]), compressed[i])
		}
	}
}


func TestDecompress(t *testing.T) {
	decompressed := decompress(initMapWithLettersAndCodes(1, 129), []int{49, 50, 51, 52, 53, 54, 55, 55, 56, 57})
	decoded := "1234567789"
	for i := range decompressed {
		if decompressed[i] != decoded[i] {
			t.Error("Expected " + string(decoded[i]), decompressed[i])
		}
	}
}
