# Huffman semi Adaptive Decrease

Implementation of the semi adaptive decrement huffman algorithm using Golang.
Unlike traditional huffman, the semi-adaptive decrement huffman decrements its frequency list by each compressed byte to its byte. Modifying your huffman tree each time a user-defined value is reached for reframing.

### Usage

1. Build the execution file
```shell
$ go build -o huffman
```

2. Execute with the instance following the flags to compress or uncompress

```shell
$ ./huffman [targeting] [file]
```
##### file

* **file**: Path to the file

##### targeting

* **-c**: compress the file
* **-u**: uncompress the file

---

#### example:

```shell
$ ./huffman -c instances/t1.bin

--	  Encoding started

| File Compressed, File Location: instances/t1.bin.compr |

$ ./huffman -u instances/t1.bin.compr

 --	  Decoding started.

| File Uncompressed, File Location: instances/t1.bin.uncmp |
```
