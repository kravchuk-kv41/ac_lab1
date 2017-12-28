package main

import (
	"os"
	"bufio"
	"strings"
	"iostream"
	"sync"
	"lzw"
	"data"
)

const (
<<<<<<< HEAD
	PATH_TO_ATTACHMENTS = "/home/kate/GolangProjects/lzw/files/"
=======
	PATH_TO_ATTACHMENTS = "/home/sergey/Projects/Golang/LzwCompressor/attachments/"
>>>>>>> 0eb29af9d427e7bdd2593d744ac0360a1d143141
	TEST_FILE_NAME_C_IN   = "test_in.txt"
	TEST_FILE_NAME_C_OUT  = "test_out.lzw"
	TEST_FILE_NAME_D_IN   = "test_out.lzw"
	TEST_FILE_NAME_D_OUT  = "decompress_out.txt"
)


func main() {
	io := bufio.NewReader(os.Stdin)
	for  {
		print("Enter command [c, d, e] > ")
		enteredVal, err := io.ReadString('\n')
		enteredVal = strings.Trim(enteredVal, "\n")
		if err != nil {
			panic(err)
		}
		if enteredVal == "c" {
			compress()
			os.Exit(0)
		}
		if enteredVal == "d" {
			decompress()
			os.Exit(0)
		}
		if enteredVal == "e" {
			os.Exit(0)
		}
		println("Enter correct command")
	}
}


func compress() {
	writer, _ := os.OpenFile(
		PATH_TO_ATTACHMENTS + TEST_FILE_NAME_C_OUT + ".lzw",
		os.O_CREATE | os.O_RDWR,
		0666)
	defer writer.Close()
	res := iostream.ReadUncompressedFile(PATH_TO_ATTACHMENTS + TEST_FILE_NAME_C_IN)
	compressing(res)
	writeCompressed(writer)
}

func decompress() {
	writer, _ := os.OpenFile(
		PATH_TO_ATTACHMENTS + TEST_FILE_NAME_D_OUT,
		os.O_CREATE | os.O_RDWR,
		0666)
	defer writer.Close()
	res := iostream.ReadByteSliceFromFile(PATH_TO_ATTACHMENTS + TEST_FILE_NAME_D_IN + ".lzw")
	decompressing(res)
	writeDecompressed(writer)
}

func compressing(res []string) {
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	wg.Add(len(res))
	for i, j := range res {
		go func(id int) {
			str := lzw.Compress(lzw.GenerateDict(), j)
			mutex.Lock()
			data.CompressedChunksStorage[id] = str
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}


func decompressing(res [][]int) {
	var wg sync.WaitGroup
	mutex := &sync.Mutex{}
	wg.Add(len(res))
	for i, j := range res {
		go func(id int) {
			str := lzw.Decompress(lzw.GenerateDict(), j)
			mutex.Lock()
			data.DecompressedChunksStorage[id] = str
			mutex.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func writeCompressed(file *os.File) {
	for k := 0; k < len(data.CompressedChunksStorage); k++ {
		iostream.WriteByteSliceToFile(file, data.CompressedChunksStorage[k])
	}
}

func writeDecompressed(file *os.File) {
	for k := 0; k < len(data.DecompressedChunksStorage); k++ {
		file.WriteString(data.DecompressedChunksStorage[k])
	}
}
