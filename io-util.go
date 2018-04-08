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
