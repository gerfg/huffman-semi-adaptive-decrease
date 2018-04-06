package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"time"
)

func decodeFile(fileName string) {
	start := time.Now()
	fmt.Println("\n --\t  Decoding started.\n")
	frequency, data := getDecodeData(fileName)
	dcd := dataToString(data)
	root := huffmanTree(frequency)
	decodeStringAndCreateFile(fileName, dcd, root, frequency)
	allTime := time.Since(start)
	fmt.Printf("Time to Decode: %s\n", allTime)
}

func decodeStringAndCreateFile(fileName string, dcd string, root Node, frequency []uint16) {
	var bytesToWrite []byte
	var nodeSearch Node = root
	for len(dcd) > 1 {
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
			frequency[nodeSearch.Letter]--
			if countRemainingLeafs(frequency) > 0 {
				root = huffmanTree(frequency)
				nodeSearch = root
			} else {
				break
			}
		}
	}
	fileName = fileName[:(len(fileName) - 6)]
	ioutil.WriteFile(fileName+".umcmp", bytesToWrite, 0644)
	fmt.Println("| File Uncompressed, File Location: " + fileName + ".uncmp |\n\n")
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
