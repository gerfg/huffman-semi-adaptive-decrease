package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

func encodeFile(fileName string) {
	start := time.Now()

	var extension = filepath.Ext(fileName)
	f, err := os.Create("log/" + fileName[10:len(fileName)-len(extension)] + ".txt")
	checkError(err)
	f.Close()

	frequency, size, data := getFrequencySlice(fileName)
	backupFrequency := make([]uint16, len(frequency))
	copy(backupFrequency, frequency)

	fmt.Println("\n --\t  Encoding started\n")
	tree := huffmanTree(frequency)

	var cds = make(map[uint16]string, size)
	generateCodes(tree, cds)

	compr := createEncodeString(data, cds, frequency)

	createEncodedFile("encoded/"+fileName[10:], compr, backupFrequency)
	fmt.Println("| File Compressed, File Location: encoded/" + fileName + " |")
	allTime := time.Since(start)
	fmt.Printf("\nTime to Encode: %s\n", allTime)
}

func getFrequencySlice(fileName string) (frequency []uint16, size int, data []byte) {
	data, _, size = ReadFile(fileName)

	var extension = filepath.Ext(fileName)
	log, err := os.OpenFile("log/"+fileName[10:len(fileName)-len(extension)]+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)
	defer log.Close()

	defer log.Close()

	frequency = make([]uint16, 256)
	for i := range frequency {
		frequency[i] = 0
	}

	for _, vl := range data {
		frequency[vl]++
	}
	fmt.Fprintln(log, " -- Encode --\n")
	fmt.Fprintln(log, " >Frequency\n")
	for idx, vl := range frequency {
		if vl > 0 {
			fmt.Fprintf(log, "%d -> %s - %d\n", idx, string(idx), vl)
		}
	}
	return frequency, size, data
}

func createEncodeString(data []byte, codes map[uint16]string, frequency []uint16) (compressed string) {
	for idx, vl := range data {
		if idx != len(data)-1 {
			compressed += codes[uint16(vl)]
			frequency[uint16(vl)]--
			root := huffmanTree(frequency)
			generateCodes(root, codes)
		}
	}
	return compressed
}

func createEncodedFile(fileName string, compress string, frequency []uint16) {
	var bt2 uint8
	var bitsBuffer = 0

	out, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}
	defer out.Close()

	bytesCreated := 0
	lastBits := len(compress) % 8

	var bytesToWrite []byte

	buf := new(bytes.Buffer)
	for _, vl := range frequency {
		err := binary.Write(buf, binary.LittleEndian, uint16(vl))
		if err != nil {
			fmt.Println("binary.Write failed:", err)
		}
	}
	bytesToWrite = buf.Bytes()

	for _, vl := range compress {
		if vl == '0' {
			bt2 = bt2 << 1
		}
		if vl == '1' {
			bt2 = bt2<<1 + 1
		}
		bitsBuffer++
		if bitsBuffer == 8 {
			bytesCreated++
			bytesToWrite = append(bytesToWrite, bt2)
			bitsBuffer = 0
			bt2 = 0
		}
	}
	for i := 0; i < (8 - lastBits); i++ {
		bt2 = bt2 << 1
	}
	bytesCreated++
	bytesToWrite = append(bytesToWrite, bt2)

	err = ioutil.WriteFile(fileName, bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}

}
