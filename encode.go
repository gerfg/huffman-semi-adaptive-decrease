package main

import (
	"fmt"
)

func encodeFile(fileName string) {
	frequency, size, data := getFrequencySlice(fileName)

	fmt.Println("\n --\t  Encoding started\n")
	tree := huffmanTree(frequency)

	var cds = make(map[uint16]string, size)
	generateCodes(tree, cds)

	compr := createEncodeString(data, cds)

	createEncodedFile("output/compressed.bin", compr, frequency)
	fmt.Println("| File Compressed, File Location: output/compressed.bin |")
}
