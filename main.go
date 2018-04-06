package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		panic("Wrong Input.\n\n example:  ./huffman-semi-adaptative-decrease [-c or -u] [fileName]\n\n-fileName -> local of a file to compress\n-c to Compress the file\n-u to Uncompress the file")
	}
	if _, err := os.Stat(args[1]); err == nil {
		if args[0] == "-C" || args[0] == "-c" {
			encodeFile(args[1])
		}
		if args[0] == "-U" || args[0] == "-u" {
			decodeFile(args[1])
		}
	} else {
		panic("File don't exist in folder.")
	}

}
