package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' 9%d bytes) \n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	// チャンク格納
	var chunks []io.Reader
	// 最初の8バイトをとばす
	// 最初の8バイトはシグニチャ
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		// 次のチャンクの先頭に追加
		// 現在位置は長さを読み終わった位置なので
		// チャンク名(4バイト) + データ長 + CRC(4バイト)先に移動
		// 4バイトごとに情報が格納されている
		// | 8バイト(シグニチャ)| 4バイト(長さ) | 4バイト(種類) | データ | 4バイト(CRC) | ...
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

func main() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
