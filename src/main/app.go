package main

import (
	"os"
	"strings"
	"bufio"
	"lzw"
	"time"
	"fmt"
)


func main() {
	io := bufio.NewReader(os.Stdin)
	for {
		print("1 -- compress\n2 -- decompress\n> ")
		code := read(io)
		print("Enter file IN > ")
		filePathIn := read(io)
		print("Enter file OUT > ")
		filePathOut := read(io)
		if code == "1" {
			Timestamp()
			lzw.Compress(filePathIn, filePathOut)
			Stamp("compression completed successfully")
			os.Exit(0)
		}
		if code == "2" {
			Timestamp()
			lzw.Decompress(filePathIn, filePathOut)
			Stamp("decompression completed successfully")
			os.Exit(0)
		}
		println("Enter correct command")
	}

}

func read(io *bufio.Reader) string {
	code, _ := io.ReadString('\n')
	code = strings.Trim(code, "\n")
	return code
}

var timestamp int64

func Timestamp() {
	timestamp = time.Now().Unix()
}

func Stamp(str string) {
	fmt.Printf("%s: %d\n", str, time.Now().Unix() - timestamp)
}
