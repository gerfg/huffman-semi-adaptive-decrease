package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

func decodeFile(fileName string) {
	fmt.Println("\n --\t  Decoding started.\n")
	frequency, data := getDecodeData(fileName)
	dcd := dataToString(data)
	root := huffmanTree(frequency)
	decodeStringAndCreateFile(dcd, root)
}

func decodeStringAndCreateFile(dcd string, root Node) {
	var bytesToWrite []byte
	var nodeSearch Node = root
	for len(dcd) > 2 {
		if nodeSearch.Letter == uint16(257) {
			if dcd[0] == '0' {
				if nodeSearch.Esq != nil {
					nodeSearch = *nodeSearch.Esq
				}
				dcd = dcd[1:]
			} else {
				if nodeSearch.Dir != nil {
					nodeSearch = *nodeSearch.Dir
				}
				dcd = dcd[1:]
			}
		} else {
			bytesToWrite = append(bytesToWrite, byte(nodeSearch.Letter))
			nodeSearch = root
		}
	}
	ioutil.WriteFile("output/uncompressed.bin", bytesToWrite, 0644)
	fmt.Println("| File Uncompressed, File Location: output/uncompressed.bin |\n\n")
}

func getDecodeData(fileName string) (frequency []uint16, data []byte) {
	data, err, _ := ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	frequency = make([]uint16, 256)

	dataFrequency := data[:512]
	data = data[512:]

	bff := bytes.NewReader(dataFrequency)
	binary.Read(bff, binary.LittleEndian, &frequency)
	return frequency, data
}

func dataToString(data []byte) (decompr string) {
	var sliceCompressor []string = make([]string, len(data))
	for idx, vl := range data {
		sliceCompressor[idx] = fmt.Sprintf("%0.8b", vl)
	}
	for _, vl := range sliceCompressor {
		decompr += vl
	}
	return decompr
}
