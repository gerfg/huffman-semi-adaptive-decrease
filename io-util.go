package main

import (
	"io/ioutil"
	"log"
)

func ReadFile(fileName string) (data []byte, err error, size int) {
	data, err = ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	return data, err, len(data)
}

func getFrequencySlice(fileName string) (frequency []uint16, size int, data []byte) {
	frequency = make([]uint16, 256)
	for i := range frequency {
		frequency[i] = 0
	}

	data, _, size = ReadFile(fileName)
	for _, vl := range data {
		frequency[vl]++
	}
	return frequency, size, data
}
