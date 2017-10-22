package lzw

import (
	"testing"
	"time"
	"fmt"
)

func BenchmarkLzwCompress(b *testing.B) {
	start := time.Now()
	Compress(
		"/home/kate/Go/lzw/files/testfile_in.txt",
		"/home/kate/Go/lzw/files/testfile_out.txt")
	fmt.Println(time.Since(start))
	//os.Remove("/home/kate/Go/lzw/files/testfile_in-res.txt")
}


func BenchmarkLzwDecompress(b *testing.B) {
	start := time.Now()
	Decompress(
		"/home/kate/Go/lzw/files/testfile_out.txt",
		"/home/kate/Go/lzw/files/testfile_in.txt")
	fmt.Println(time.Since(start))
	//os.Remove("/home/kate/Go/lzw/files/testfile_out-res.txt")
}
