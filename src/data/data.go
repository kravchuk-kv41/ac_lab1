package data


var (
	CompressedChunksStorage = make(map[int][]int)
	DecompressedChunksStorage = make(map[int]string)
)