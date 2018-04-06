# huffman-semi-adaptive-decrease

go build -o huffman

./huffman [-c or -u] [file]

-c -> compress the file
-u -> uncompress the file
file -> the file path

example:

./huffman -c instances/t1.bin

--	  Encoding started

| File Compressed, File Location: instances/t1.bin.compr |

./huffman -u instances/t1.bin.compr

 --	  Decoding started.

| File Uncompressed, File Location: instances/t1.bin.uncmp |
