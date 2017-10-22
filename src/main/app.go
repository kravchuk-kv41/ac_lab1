package main

import (
	"os"
	"strings"
	"bufio"
	"lzw"
)


func main() {
	io := bufio.NewReader(os.Stdin)
	for  {
		print("1 -- compress\n2 -- decompress\n> ")
		code, err := io.ReadString('\n')
		code = strings.Trim(code, "\n")
		if err != nil {
			panic(err)
		}
		print("Enter file IN > ")
		filePathIn := getFilePath(*io)
		print("Enter file OUT > ")
		filePathOut := getFilePath(*io)
		if code == "1" {
			lzw.Compress(filePathIn, filePathOut)
			println("compression completed successfully")
			os.Exit(0)
		}
		if code == "2" {
			lzw.Decompress(filePathIn, filePathOut)
			println("decompression completed successfully")
			os.Exit(0)
		}
		println("Enter correct command")
	}

}


func getFilePath(io bufio.Reader) string {
	filePath, _ := io.ReadString('\n')
	filePath = strings.Trim(filePath, "\n")
	return filePath
}
